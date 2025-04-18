package handlers

import (
	"log"
	"my-app/web/components/pages"
	"net/http"
)

type HomeHandler struct {
	// service *service.CountryService
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	err := pages.Index().Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}
