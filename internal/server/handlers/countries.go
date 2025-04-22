package handlers

import (
	"log"
	"my-app/internal/service"
	"my-app/internal/utils"
	countriesPage "my-app/web/components/pages/countries"
	"strings"

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

	pageSizeStr := r.URL.Query().Get("pageSize")
	encodedCursor := r.URL.Query().Get("cursor")

	countries, nextCursor, err := h.service.GetCountries(ctx, encodedCursor, pageSizeStr)
	if err != nil {
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		log.Printf("Error fetching countries: %v", err)
		return
	}

	log.Printf("Fetched %d countries, nextCursor: %q", len(countries), nextCursor)

	isHTMX := r.Header.Get("HX-Request") != ""

	if isHTMX && encodedCursor != "" {
		err = countriesPage.CountryList(countries, nextCursor).Render(ctx, w)
	} else {
		// Full page request: render entire layout
		err = countriesPage.Countries(countries, nextCursor).Render(ctx, w)
	}

	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Templ render error: %v", err)
	}

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
	ctx := r.Context()
	query := strings.TrimSpace(r.FormValue("query"))

	// Optional: make limit configurable via query param

	countries, err := h.service.SearchCountry(ctx, query, utils.DefaultPageSize)
	if err != nil {
		http.Error(w, "Failed to search countries", http.StatusInternalServerError)
		log.Printf("Search error (query=%q): %v", query, err)
		return
	}

	// Render partial country list for HTMX response
	err = countriesPage.CountryList(countries, "").Render(ctx, w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Render error (query=%q): %v", query, err)
		return
	}
}
