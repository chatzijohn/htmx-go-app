package routes

import (
	"context"
	"log"
	"my-app/http/handlers"
	"my-app/internal/config"
	"net/http"
)

type HttpRouter struct {
	Cancel  context.CancelFunc
	Context context.Context
	Config  *config.ServerConfig
}

func NewRouter() *HttpRouter {
	cfg := config.Load()

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /", handlers.CountryList)
	mux.HandleFunc("GET /country/{name}", handlers.CountryDetail)

	log.Printf("Server listening on :%s", cfg.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.PORT, mux))

	return &HttpRouter{}
}
