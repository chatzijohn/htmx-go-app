package service

import (
	"context"
	"fmt"
	"my-app/internal/models"
	"my-app/internal/repository"
	"my-app/internal/utils"
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

func (s *CountryService) GetCountry(ctx context.Context, name string) (*models.Country, error) {
	normalizedName, err := utils.NormalizeCountryName(name)
	if err != nil {
		return nil, fmt.Errorf("error while trying to normalize %s, %w", name, err)
	}
	country, err := s.repo.FetchCountry(ctx, normalizedName)

	if err != nil {
		return nil, fmt.Errorf("error while fetching %s, %w", name, err)
	}
	return country, nil
}
