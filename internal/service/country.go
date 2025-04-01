package service

import "my-app/internal/repository"

type countryService struct {
	repo repository.CountryRepository
}
