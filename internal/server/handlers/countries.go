package handlers

import (
	"log"
	"my-app/internal/service"
	"my-app/web/components"

	"net/http"
)

type CountryHandler struct {
	service *service.CountryService
}

func NewCountryHandler(svc *service.CountryService) *CountryHandler {
	return &CountryHandler{service: svc}
}

func (h *CountryHandler) GetCountries(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get countries")
	ctx := r.Context()

	countries, err := h.service.GetCountries(ctx)
	if err != nil {
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		log.Printf("Error fetching countries: %v", err)
		return
	}

	// Render Templ component
	err = components.Countries(countries).Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}

func (h *CountryHandler) GetCountry(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request to get country: %s \n", r.PathValue("name"))
	ctx := r.Context()

	country, err := h.service.GetCountry(ctx, r.PathValue("name"))
	if err != nil {
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		log.Printf("Error fetching countries: %v", err)
		return
	}

	// If not found, return 404
	if country == nil {
		http.NotFound(w, r)
		return
	}

	// Render Templ component
	err = components.CountryDetails(country).Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}

func (h *CountryHandler) SearchCountry(w http.ResponseWriter, r *http.Request) {

	query := r.FormValue("query")
	log.Printf("Received request to get country: %s \n", query)
	ctx := r.Context()

	countries, err := h.service.SearchCountry(ctx, query, 10)
	if err != nil {
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		log.Printf("Error fetching countries: %v", err)
		return
	}

	// If not found, return 404
	if countries == nil {
		http.NotFound(w, r)
		return
	}

	// Render Templ component
	err = components.CountryList(countries).Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}
