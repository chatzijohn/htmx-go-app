package service

import (
	"context"
	"fmt"
	"log"
	"my-app/internal/models"
	"my-app/internal/repository"
	"my-app/internal/utils"
	"strconv"
)

type CountryService struct {
	repo *repository.CountryRepository
}

func NewCountryService(repo *repository.CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (s *CountryService) GetCountries(ctx context.Context, encodedCursor string, pageSizeStr string, direction string) ([]*models.Country, string, string, error) {
	var decodedCursor *models.CountryCursor
	if encodedCursor != "" {
		cursor, err := utils.DecodeCursor[models.CountryCursor](encodedCursor)
		if err != nil {
			return nil, "", "", fmt.Errorf("invalid cursor: %w", err)
		}
		decodedCursor = cursor
	}

	pageSize := utils.DefaultPageSize
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil {
			pageSize = ps
		}
	}

	countries, nextCursorStruct, err := s.repo.FetchCountries(ctx, decodedCursor, pageSize, direction)
	if err != nil {
		return nil, "", "", err
	}

	var nextCursor string
	if nextCursorStruct != nil {
		nextCursor, err = utils.EncodeCursor(nextCursorStruct)
		if err != nil {
			return nil, "", "", fmt.Errorf("failed to encode next cursor: %w", err)
		}
	}

	// 👇 Move this block inside the service
	var prevCursor string
	if len(countries) > 0 {
		first := countries[0]
		prevCursorStruct := &models.CountryCursor{ID: first.ID, Name: first.Name}
		prevCursor, _ = utils.EncodeCursor(prevCursorStruct) // ignore error for now
	}

	return countries, nextCursor, prevCursor, nil
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
