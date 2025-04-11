package service

import (
	"my-app/internal/repository"
)

type ServiceContainer struct {
	CountryService *CountryService
}

func NewServiceContainer(repos *repository.RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		CountryService: NewCountryService(repos.CountryRepository),
	}
}
