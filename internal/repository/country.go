package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-app/internal/models"

	"github.com/rs/xid"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (r *CountryRepository) FetchCountries(ctx context.Context, cursor *models.CountryCursor, pageSize int) ([]*models.Country, *models.CountryCursor, error) {
	query := `
		SELECT id, name, code, capital, continent
		FROM countries
		WHERE ((id > $1 AND name = $2) OR name > $2)
		ORDER BY name, id
		LIMIT $3;`

	cursorID := xid.ID{}
	cursorName := ""
	if cursor != nil {
		cursorID = cursor.ID
		cursorName = cursor.Name
	}

	// Fetch pageSize + 1 to detect if there's a next page
	rows, err := r.db.QueryContext(ctx, query, cursorID, cursorName, pageSize+1)
	if err != nil {
		return nil, nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var countries []*models.Country

	for rows.Next() {
		c := &models.Country{}
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.Capital, &c.Continent); err != nil {
			return nil, nil, fmt.Errorf("scan error: %v", err)
		}
		countries = append(countries, c)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %v", err)
	}

	// Determine if there's a next page
	var nextCursor *models.CountryCursor
	if len(countries) > pageSize {
		last := countries[pageSize] // this is the extra one
		nextCursor = &models.CountryCursor{
			ID:   last.ID,
			Name: last.Name,
		}
		countries = countries[:pageSize] // trim to actual pageSize
	}

	return countries, nextCursor, nil
}

func (r *CountryRepository) FetchCountry(ctx context.Context, name string) (*models.Country, error) {
	query := "SELECT id, name, code, capital, continent FROM countries WHERE name LIKE $1 COLLATE NOCASE"
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
