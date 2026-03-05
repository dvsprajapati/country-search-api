package handler

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"country-search-api/internal/cache"
	"country-search-api/internal/client"
	"country-search-api/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestSearchCountryHandler(t *testing.T) {
	memCache := cache.NewMemoryCache()
	restClient := client.NewRestCountriesClient()
	countryService := service.NewCountryService(memCache, restClient)
	countryHandler := NewCountryHandler(countryService)

	req := httptest.NewRequest(http.MethodGet, "/api/countries/search?name=India", nil)
	respRec := httptest.NewRecorder()

	countryHandler.SearchCountry(respRec, req)

	assert.Equal(t, http.StatusOK, respRec.Code)
}

func TestConcurrentAPIRequests(t *testing.T) {
	cache := cache.NewMemoryCache()
	client := client.NewRestCountriesClient()
	service := service.NewCountryService(cache, client)
	handler := NewCountryHandler(service)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := httptest.NewRequest(http.MethodGet, "/api/countries/search?name=India", nil)
			w := httptest.NewRecorder()
			handler.SearchCountry(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		}()
	}

	wg.Wait()
}
