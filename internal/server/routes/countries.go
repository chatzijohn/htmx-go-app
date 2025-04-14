package routes

import (
	"my-app/internal/server/handlers"
	"my-app/internal/service"
	"net/http"
)

type CountryRoutes struct {
	handler *handlers.CountryHandler
}

func NewCountryRoutes(service *service.CountryService) *CountryRoutes {
	return &CountryRoutes{
		handler: handlers.NewCountryHandler(service),
	}
}

func (c *CountryRoutes) Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", c.handler.GetCountries)
	mux.HandleFunc("GET /{name}", c.handler.GetCountry)
	mux.HandleFunc("POST /search", c.handler.SearchCountry)
	return mux
}
