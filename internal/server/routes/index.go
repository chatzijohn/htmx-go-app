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

func (r *Router) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Mount all route groups here
	countryRoutes := NewCountryRoutes(r.services.CountryService)
	mux.Handle("/countries/", http.StripPrefix("/countries", countryRoutes.Register()))

	return mux
}
