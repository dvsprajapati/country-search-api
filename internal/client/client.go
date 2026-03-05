package client

import (
	"context"

	"github.com/dvsprajapati/country-search-api/app/model"
)

type RestCountriesClient interface {
	GetCountry(ctx context.Context, name string) (*model.Country, error)
}
