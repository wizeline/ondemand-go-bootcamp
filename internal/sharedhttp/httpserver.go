package sharedhttp

import (
	"net/http"
	"time"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"

	"github.com/go-chi/chi/v5"
)

// NewHTTPServer returns a new implementation of a http.Server with chi.Mux support.
func NewHTTPServer(cfg configuration.HTTP, router *chi.Mux) *http.Server {
	return &http.Server{
		Handler:      router,
		Addr:         cfg.Address(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}
}
