package routes

import (
	"log"
	"my-app/config"
	"my-app/http/handlers"
	"net/http"
)

func Router() {
	cfg := config.Load()

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /", handlers.CountryList)
	mux.HandleFunc("GET /country/{name}", handlers.CountryDetail)

	log.Printf("Server listening on :%s", cfg.SERVER.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.SERVER.PORT, mux))
}
