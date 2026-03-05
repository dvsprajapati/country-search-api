package service

import (
	"context"
	"country-search-api/app/model"
	"country-search-api/internal/cache"
	"country-search-api/internal/client"
)

type CountryService struct {
	cache  cache.Cache
	client *client.RestCountriesClient
}

func NewCountryService(c cache.Cache, cl *client.RestCountriesClient) *CountryService {
	return &CountryService{
		cache:  c,
		client: cl,
	}
}

func (s *CountryService) SearchCountry(ctx context.Context, name string) (*model.Country, error) {
	val, found := s.cache.Get(name)
	if found {
		return val.(*model.Country), nil
	}

	country, err := s.client.GetCountry(ctx, name)
	if err != nil {
		return nil, err
	}

	s.cache.Set(name, country)
	return country, nil
}
