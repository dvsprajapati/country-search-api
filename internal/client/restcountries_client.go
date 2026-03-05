package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dvsprajapati/country-search-api/app/model"
)

type restCountriesClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewRestCountriesClient() RestCountriesClient {
	return &restCountriesClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://restcountries.com/v3.1/name/",
	}
}

func (c *restCountriesClient) GetCountry(ctx context.Context, name string) (*model.Country, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+name, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp []struct {
		Name struct {
			Common string `json:"common"`
		} `json:"name"`
		Capital    []string `json:"capital"`
		Population int64    `json:"population"`
		Currencies map[string]struct {
			Symbol string `json:"symbol"`
		} `json:"currencies"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	if len(apiResp) == 0 {
		return nil, fmt.Errorf("no data found for country: %s", name)
	}
	data := apiResp[0]

	currency := ""
	for _, v := range data.Currencies {
		currency = v.Symbol
		break
	}

	capital := ""
	if len(data.Capital) > 0 {
		capital = data.Capital[0]
	}

	country := &model.Country{
		Name:       data.Name.Common,
		Capital:    capital,
		Currency:   currency,
		Population: data.Population,
	}

	return country, nil
}
