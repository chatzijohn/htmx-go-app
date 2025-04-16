package routes

import (
	"my-app/internal/server/handlers"
	"net/http"
)

type HomeRoutes struct {
	handler *handlers.HomeHandler
}

func NewHomeRoutes() *HomeRoutes {
	return &HomeRoutes{
		handler: handlers.NewHomeHandler(),
	}
}

func (h *HomeRoutes) Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.handler.GetHome) // GET
	return mux
}
