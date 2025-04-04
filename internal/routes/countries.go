package routes

import (
	"net/http"
)

type CountryRoutes struct{}

func (cr *CountryRoutes) Register(mux *http.ServeMux) {
	RegisterCountryRoutes(mux)
}

func RegisterCountryRoutes(mux *http.ServeMux) {
	// mux.HandleFunc("GET /countries", handlers.Make(handlers.GetCountry()))
	// mux.HandleFunc("GET /countries/{id}", handlers.Make(handlers.GetCountry()))
}
