package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"country-search-api/internal/cache"
	"country-search-api/internal/client"
	"country-search-api/internal/handler"
	"country-search-api/internal/service"
)

func main() {
	cache := cache.NewMemoryCache()
	client := client.NewRestCountriesClient()
	service := service.NewCountryService(cache, client)
	handler := handler.NewCountryHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/countries/search", handler.SearchCountry)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		log.Println("Server running on :8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %s\n", err)
	}

	log.Println("Server exited properly")
}
