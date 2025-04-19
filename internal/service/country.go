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

func (s *CountryService) GetCountries(ctx context.Context, encodedCursor string, pageSizeStr string) ([]*models.Country, string, string, error) {
	// Decode cursor (if present)
	var decodedCursor *models.CountryCursor
	if encodedCursor != "" {
		cursor, err := utils.DecodeCursor[models.CountryCursor](encodedCursor)
		if err != nil {
			return nil, "", "", fmt.Errorf("invalid cursor: %w", err)
		}
		decodedCursor = cursor
	}

	// Parse or default page size
	pageSize := utils.DefaultPageSize
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil {
			pageSize = ps
		}
	}

	// Fetch countries and encode nextCursor
	countries, nextCursorStruct, err := s.repo.FetchCountries(ctx, decodedCursor, pageSize)
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

	// 💡 This is the cursor we used to fetch the current page
	prevCursor := encodedCursor

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
