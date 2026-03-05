package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {

	mockResponse := `[{
		"name":{"common":"India"},
		"capital":["New Delhi"],
		"population":1380004385,
		"currencies":{"INR":{"symbol":"₹"}}
	}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewRestCountriesClient()
	client.baseURL = server.URL + "/"

	ctx := context.Background()

	country, err := client.GetCountry(ctx, "India")

	assert.NoError(t, err)
	assert.Equal(t, "India", country.Name)
	assert.Equal(t, "New Delhi", country.Capital)
	assert.Equal(t, "₹", country.Currency)
}
