package service

import (
	"context"
	"testing"

	"github.com/dvsprajapati/country-search-api/app/model"
	"github.com/dvsprajapati/country-search-api/internal/cache"

	// "country-search-api/internal/mocks"
	"github.com/dvsprajapati/country-search-api/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCountryServiceCacheMiss(t *testing.T) {
	cache := cache.NewMemoryCache()
	mockClient := new(mocks.RestCountriesClient)

	mockCountry := &model.Country{
		Name:       "India",
		Capital:    "New Delhi",
		Currency:   "₹",
		Population: 1000,
	}

	mockClient.On("GetCountry", context.Background(), "India").
		Return(mockCountry, nil)

	service := NewCountryService(cache, mockClient)

	res, err := service.SearchCountry(context.Background(), "India")

	assert.NoError(t, err)
	assert.Equal(t, "India", res.Name)
}

func TestCountryServiceCacheHit(t *testing.T) {
	cache := cache.NewMemoryCache()

	mockCountry := &model.Country{Name: "India"}
	cache.Set("India", mockCountry)

	service := NewCountryService(cache, nil)

	res, err := service.SearchCountry(context.Background(), "India")

	assert.NoError(t, err)
	assert.Equal(t, "India", res.Name)
}
