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

func (r *CountryRepository) FetchCountries(ctx context.Context, limit int, cursor string, isNext bool) ([]*models.Country, xid.ID, xid.ID, error) {
	// Validate limit
	if limit <= 0 {
		limit = 10 // Default limit
	} else if limit > 100 {
		limit = 100 // Maximum limit to prevent excessive results
	}

	var query string
	var args []interface{}

	// If cursor is provided, we need to fetch records after (next) or before (previous) the cursor
	if cursor != "" {
		if isNext {
			// Fetch records after the cursor (next)
			query = `
				SELECT id, name, code, capital, continent
				FROM countries
				WHERE name > $1
				ORDER BY name ASC
				LIMIT $2`
			args = append(args, cursor, limit)
		} else {
			// Fetch records before the cursor (previous)
			query = `
				SELECT id, name, code, capital, continent
				FROM countries
				WHERE name < $1
				ORDER BY name DESC
				LIMIT $2`
			args = append(args, cursor, limit)
		}
	} else {
		// If no cursor, fetch the first page (ordering by name)
		query = `
			SELECT id, name, code, capital, continent
			FROM countries
			ORDER BY name ASC
			LIMIT $1`
		args = append(args, limit)
	}

	// Execute query
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, xid.ID{}, xid.ID{}, fmt.Errorf("failed to query countries: %w", err)
	}
	defer rows.Close()

	var countries []*models.Country
	var lastID xid.ID
	var firstID xid.ID

	for rows.Next() {
		var country models.Country
		if err := rows.Scan(&country.ID, &country.Name, &country.Code, &country.Capital, &country.Continent); err != nil {
			return nil, xid.ID{}, xid.ID{}, fmt.Errorf("failed to scan country row: %w", err)
		}
		countries = append(countries, &country)

		// Keep track of the first and last IDs for the cursor
		if firstID == (xid.ID{}) {
			firstID = country.ID // first record of the page
		}
		lastID = country.ID // last record of the page
	}

	if err := rows.Err(); err != nil {
		return nil, xid.ID{}, xid.ID{}, fmt.Errorf("rows iteration error: %w", err)
	}

	// If there are fewer than `limit` results, then there is no next page
	var nextCursor xid.ID
	if len(countries) == limit {
		nextCursor = lastID
	}

	// For previous, we use the first ID of the current page
	var prevCursor xid.ID
	if len(countries) > 0 {
		prevCursor = firstID
	}

	return countries, nextCursor, prevCursor, nil
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
