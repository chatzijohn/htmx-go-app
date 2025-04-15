package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-app/internal/models"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (r *CountryRepository) FetchCountries(ctx context.Context) ([]*models.Country, error) {
	query := "SELECT id, name, code, capital, continent FROM countries"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	countries := []*models.Country{}

	for rows.Next() {
		country := &models.Country{}

		// Scan the result into the country struct (in the same order as the query)
		if err := rows.Scan(&country.ID, &country.Name, &country.Code, &country.Capital, &country.Continent); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}

		// Append the country to the slice (using pointer)
		countries = append(countries, country)
	}

	// Check for errors after iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the slice of countries
	return countries, nil

}

func (r *CountryRepository) FetchCountry(ctx context.Context, name string) (*models.Country, error) {
	query := "SELECT id, name, code, capital, continent FROM countries WHERE name = $1"
	// Create an empty Country struct to hold the fetched data
	country := &models.Country{}

	// Execute the query using QueryRow
	err := r.db.QueryRowContext(ctx, query, name).Scan(&country.ID, &country.Name, &country.Code, &country.Capital, &country.Continent)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil if no country found
			return nil, nil
		}
		// Return error if there is a problem with executing the query
		return nil, fmt.Errorf("could not fetch country: %v", err)
	}

	// Return the slice of countries
	return country, nil
}

func (r *CountryRepository) FindCountries(ctx context.Context, name string, limit int) ([]*models.Country, error) {
	// Validate input parameters
	if limit <= 0 {
		limit = 10 // Default limit
	} else if limit > 100 {
		limit = 100 // Maximum limit to prevent excessive results
	}

	query := `
		SELECT id, name, code, capital, continent 
		FROM countries 
		WHERE name LIKE $1 || '%' COLLATE NOCASE
		LIMIT $2`

	rows, err := r.db.QueryContext(ctx, query, name, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query countries: %w", err)
	}
	defer rows.Close()

	countries := []*models.Country{}

	for rows.Next() {
		var country models.Country
		err := rows.Scan(
			&country.ID,
			&country.Name,
			&country.Code,
			&country.Capital,
			&country.Continent,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan country row: %w", err)
		}
		countries = append(countries, &country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return countries, nil
}
