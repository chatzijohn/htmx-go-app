package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-app/internal/model"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (r *CountryRepository) FetchCountries(ctx context.Context) ([]*model.Country, error) {
	query := "SELECT id, name, code, capital, continent FROM countries"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var countries []*model.Country

	for rows.Next() {
		var country model.Country
		// Scan the result into the country struct (in the same order as the query)
		if err := rows.Scan(&country.ID, &country.Name, &country.Code, &country.Capital, &country.Continent); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}

		// Append the country to the slice (using pointer)
		countries = append(countries, &country)
	}

	// Check for errors after iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the slice of countries
	return countries, nil

}
