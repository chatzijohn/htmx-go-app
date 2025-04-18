package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"my-app/internal/config"
	"my-app/internal/db"
	"my-app/internal/repository"
	"my-app/internal/server"
	"my-app/internal/service"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg := config.Load()

	// Connect DB
	sqlDB, err := db.ConnectDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	defer sqlDB.Close()

	// Initialize repository and services
	repos := repository.NewRepositoryContainer(sqlDB)
	services := service.NewServiceContainer(repos)

	// Initialize server with services
	srv := server.New(cfg, services)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.SERVER.HOST, cfg.SERVER.PORT),
		Handler: srv,
	}

	go func() {
		log.Printf("Server starting on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Server error: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.Println("Server shutting down gracefully...")
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "Shutdown error: %s\n", err)
		}
	}()

	wg.Wait()
	log.Println("Server stopped")
	return nil
}

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %s\n", err)
		os.Exit(1)
	}
}
