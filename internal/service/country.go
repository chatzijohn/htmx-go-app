package service

import (
	"context"
	"my-app/internal/models"
	"my-app/internal/repository"
)

type CountryService struct {
	repo *repository.CountryRepository
}

func NewCountryService(repo *repository.CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (s *CountryService) GetCountries(ctx context.Context) ([]*models.Country, error) {
	countries, err := s.repo.FetchCountries(ctx)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
