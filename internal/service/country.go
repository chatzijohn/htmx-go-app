package service

import "my-app/internal/repository"

type countryService struct {
	Repo repository.CountryRepository
}
