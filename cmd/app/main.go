package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"my-app/internal/config"
	"my-app/internal/db"
	"my-app/internal/server"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// run is the core application logic that sets up and runs the server
func run(ctx context.Context, w io.Writer, args []string) error {
	// Create a context that cancels on OS interrupt signals (Ctrl+C)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel() // Ensure cancel is called when we're done

	// 1. Load application configuration
	cfg := config.Load()

	// 2. Initialize database connection
	db.ConnectDB(cfg)

	// 3. Create HTTP server with configured routes
	srv := server.New(&cfg.SERVER)

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.SERVER.HOST, cfg.SERVER.PORT), // Combine host:port
		Handler: srv,                                                // Our router/handler
	}

	// Start the server in a goroutine so it doesn't block
	go func() {
		log.Printf("Server starting on %s\n", httpServer.Addr)
		// ListenAndServe blocks until server stops
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Only log errors that aren't normal shutdowns
			fmt.Fprintf(os.Stderr, "Server error: %s\n", err)
		}
	}()

	// Set up graceful shutdown using WaitGroup for synchronization
	var wg sync.WaitGroup
	wg.Add(1) // We're waiting for one goroutine

	go func() {
		defer wg.Done() // Signal we're done when this goroutine completes

		// Wait for cancellation signal (Ctrl+C)
		<-ctx.Done()

		// Create shutdown context with 10 second timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.Println("Server shutting down gracefully...")
		// Attempt graceful shutdown
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "Shutdown error: %s\n", err)
		}
	}()

	// Wait for shutdown to complete
	wg.Wait()
	log.Println("Server stopped")
	return nil
}

// main is the entry point of the application
func main() {
	// Create root context
	ctx := context.Background()

	// Run the application and check for errors
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %s\n", err)
		os.Exit(1) // Exit with error code if something went wrong
	}
}

// TODO: Find a way to group routes together for http.NewServeMux
