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
	Config  *config.AppConfig
}

func NewRouter() *HttpRouter {
	cfg := config.Load()

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /", handlers.CountryList)
	mux.HandleFunc("GET /country/{name}", handlers.CountryDetail)

	log.Printf("Server listening on :%s", cfg.SERVER.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.SERVER.PORT, mux))

	return &HttpRouter{}
}
