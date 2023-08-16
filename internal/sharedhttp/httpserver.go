package sharedhttp

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
)

func NewHTTPServer(cfg configuration.Config, router *mux.Router) *http.Server {
	return &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         cfg.HTTP.Address(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}
}
