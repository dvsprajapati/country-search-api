package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dvsprajapati/country-search-api/internal/service"
)

type CountryHandler struct {
	service *service.CountryService
}

func NewCountryHandler(s *service.CountryService) *CountryHandler {
	return &CountryHandler{service: s}
}

// search country handler to search country name
func (h *CountryHandler) SearchCountry(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name query param required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	country, err := h.service.SearchCountry(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}
