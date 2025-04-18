package service

import (
	"context"
	"fmt"
	"log"
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

func (s *CountryService) GetCountry(ctx context.Context, slug string) (*models.Country, error) {
	normalizedName, err := utils.NormalizeCountryName(slug)
	if err != nil {
		return nil, fmt.Errorf("error while trying to normalize %s, %w", slug, err)
	}
	log.Printf("normalizedName: %s", normalizedName)
	country, err := s.repo.FetchCountry(ctx, normalizedName)

	if err != nil {
		return nil, fmt.Errorf("error while fetching %s, %w", slug, err)
	}
	return country, nil
}

func (s *CountryService) SearchCountry(ctx context.Context, name string, limit int) ([]*models.Country, error) {
	countries, err := s.repo.FindCountries(ctx, name, limit)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
