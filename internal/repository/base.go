package repository

import (
	"database/sql"
)

// RepositoryBase is a container for all repositories in the application.
type RepositoryBase struct {
	CountryRepository *CountryRepository
}

// NewRepositoryBase initializes all repositories and returns a container.
func NewRepositoryBase(db *sql.DB) *RepositoryBase {
	return &RepositoryBase{
		CountryRepository: NewCountryRepository(db),
	}
}
