package server

import (
	"context"
	"log"
	"my-app/internal/config"
	"my-app/internal/db"
	"my-app/internal/repository"
	"net/http"
)

type HttpServer struct {
	Cancel  context.CancelFunc
	Context context.Context
	Config  *config.AppConfig
	Repos   *repository.RepositoryContainer
	Mux     *http.ServeMux
}

func NewServer() *HttpServer {
	cfg := config.Load()
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database %v\n", err)
	}

	// Initialize the repository container with the database connection.
	repoContainer := repository.NewRepositoryContainer(conn)

	// log.Printf("Server listening on :%s", cfg.SERVER.PORT)
	// log.Fatal(http.ListenAndServe(":"+cfg.SERVER.PORT, mux))

	return &HttpServer{Mux: http.NewServeMux(), Cancel: cancel, Context: ctx, Repos: repoContainer, Config: cfg}
}
