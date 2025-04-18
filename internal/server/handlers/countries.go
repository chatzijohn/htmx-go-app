package handlers

import (
	"log"
	"my-app/internal/models"
	"my-app/internal/service"
	countriesPage "my-app/web/components/pages/countries"

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

	err = countriesPage.Countries(countries).Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}

func (h *CountryHandler) GetCountry(w http.ResponseWriter, r *http.Request) {

	slug := r.PathValue("slug")
	log.Printf("Received request to get country: %s\n", slug)
	ctx := r.Context()

	country, err := h.service.GetCountry(ctx, slug)
	if err != nil {
		http.Error(w, "Failed to get country", http.StatusInternalServerError)
		log.Printf("Error fetching country: %v", err)
		return
	}

	if country == nil {
		http.NotFound(w, r)
		return
	}

	err = countriesPage.CountryDetails(country).Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}

func (h *CountryHandler) SearchCountry(w http.ResponseWriter, r *http.Request) {

	query := r.FormValue("query")
	ctx := r.Context()
	var countries []*models.Country
	var err error

	if query == "" {
		countries, err = h.service.GetCountries(ctx)
	} else {
		log.Printf("Received request to search countries with query: %s\n", query)
		countries, err = h.service.SearchCountry(ctx, query, 10)
	}

	if err != nil {
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		log.Printf("Error fetching countries: %v", err)
		return
	}

	if countries == nil {
		http.NotFound(w, r)
		return
	}

	err = countriesPage.CountryList(countries).Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}
}
