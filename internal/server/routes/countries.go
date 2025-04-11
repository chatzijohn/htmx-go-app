package routes

import (
	"my-app/internal/server/handlers"
	"net/http"
)

type Handler struct{}

func NewCountryHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterCountryRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", handlers.GetCountries)
	r.HandleFunc("GET /{id}/", handlers.GetCountry)
	return r
}
