package main

import (
	"context"
	"fmt"
	"log"

	"my-app/http/routes"
	"my-app/internal/db"
	"my-app/internal/repository"
)

func main() {

	// Initialize database
	db, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize DB:", err)
	}

	ctx := context.Background()

	repoBase := repository.NewRepositoryBase(db)
	countryRepo := repoBase.CountryRepository
	countries, err := countryRepo.FetchCountries(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, country := range countries {
		fmt.Printf("Country: %s, Capital: %s\n", country.Name, country.Capital)
	}

	// Iinitialize router
	routes.Router()
}
