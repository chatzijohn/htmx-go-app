package server

import (
	"my-app/internal/config"
	"my-app/internal/server/routes"
	"my-app/internal/service"
	"net/http"
)

func New(serverConfig *config.ServerConfig, services *service.ServiceContainer) http.Handler {
	mux := http.NewServeMux()

	// Serve static files (CSS, JS, etc.)
	staticFs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", staticFs))

	// Your other route(s)
	countryRoutes := routes.NewCountryRoutes(services.CountryService)
	mux.Handle("/countries/", http.StripPrefix("/countries", countryRoutes.Register()))

	return mux
}
