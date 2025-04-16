package routes

import (
	"my-app/internal/service"
	"net/http"
)

type Router struct {
	services *service.ServiceContainer
}

func NewRouter(services *service.ServiceContainer) *Router {
	return &Router{
		services: services,
	}
}

func (r *Router) RegisterRoutes(mux *http.ServeMux) http.Handler {

	// Mount all route groups here
	homeRoutes := NewHomeRoutes()
	countryRoutes := NewCountryRoutes(r.services.CountryService)
	mux.Handle("/home/", http.StripPrefix("/home", homeRoutes.Register(mux)))
	mux.Handle("/countries/", http.StripPrefix("/countries", countryRoutes.Register(mux)))

	return mux
}
