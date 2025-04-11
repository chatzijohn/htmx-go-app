package server

import (
	"my-app/internal/config"
	"my-app/internal/server/routes"
	"net/http"
)

func New(serverConfig *config.ServerConfig) http.Handler {
	mux := http.NewServeMux()

	countryHandler := routes.NewCountryHandler()
	countryRouter := countryHandler.RegisterCountryRoutes()
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)
	mux.Handle("/countries/", http.StripPrefix("/countries", countryRouter))
	var handler http.Handler = mux
	return handler
}
