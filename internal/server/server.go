package server

import (
	"my-app/internal/config"
	"my-app/internal/server/routes"
	"my-app/internal/service"
	"net/http"
)

func New(serverConfig *config.ServerConfig, services *service.ServiceContainer) http.Handler {
	mux := http.NewServeMux()

	countryRoutes := routes.NewCountryRoutes(services.CountryService)
	mux.Handle("/countries/", http.StripPrefix("/countries", countryRoutes.Register()))

	return mux
}
