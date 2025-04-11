package handlers

import (
	"fmt"
	"log"
	"my-app/internal/repository"
	"my-app/internal/service"
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request for debugging
	log.Println("Received request to get countries")

	// Use the context from the request directly
	ctx := r.Context()

	countryService := service.NewCountryService(&repository.CountryRepository{})

	// Call the service to get the list of countries
	countries, err := countryService.GetCountries(ctx)
	if err != nil {
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		log.Printf("Error fetching countries: %v", err)
		return
	}

	log.Print(countries)
	// Set the content type as HTML
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// Start the HTML content
	w.Write([]byte("<html><body><h1>Countries List</h1><ul>"))

	// Loop through countries and display them as list items
	for _, country := range countries {
		w.Write([]byte(fmt.Sprintf("<li>%s (Code: %s, Capital: %s, Continent: %s)</li>",
			country.Name, country.Code, country.Capital, country.Continent)))
	}

	// Close the HTML content
	w.Write([]byte("</ul></body></html>"))

	// Additional logging to ensure the handler is being executed
	log.Println("Successfully sent response")
}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
