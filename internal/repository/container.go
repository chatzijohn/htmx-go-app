package repository

import (
	"database/sql"
)

// RepositoryContainer is a container for all repositories in the application.
type RepositoryContainer struct {
	CountryRepository *CountryRepository
}

// NewRepositoryContainer initializes all repositories and returns a container.
func NewRepositoryContainer(db *sql.DB) *RepositoryContainer {
	return &RepositoryContainer{
		CountryRepository: NewCountryRepository(db),
	}
}
