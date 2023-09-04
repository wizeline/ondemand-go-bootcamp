package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/controller"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/sharedhttp"
)

// ApiHTTP is an HTTP API implementation with http.Server and chi.Mux support.
type ApiHTTP interface {
	Start()
}

type apiHTTP struct {
	cfg    configuration.Config
	server *http.Server
}

// NewApiHTTP returns a new ApiHTTP implementation.
// It uses the configuration instance to set the components.
// It provides the controllers, repository and service dependencies.
// Moreover, implements a http.Server instance with chi.Mux router support.
// Returns an error if any implementation of the dependencies fails.
func NewApiHTTP() (ApiHTTP, func(), error) {
	// Dependencies
	cfg := configuration.GetInstance()
	fruitRepo, err := repository.NewFruitCsv(cfg.Database.CSV)
	if err != nil {
		return nil, nil, err
	}
	fruitSvc := service.NewFruit(fruitRepo)

	// Router
	r := sharedhttp.NewChi(cfg.AppVersion)
	r.Add("HealthCheck", controller.NewHealthCheck())
	r.Add("Home", controller.NewHome())
	r.Add("Fruit", controller.NewFruit(fruitSvc))
	r.RegisterRoutes()

	return apiHTTP{
		cfg:    cfg,
		server: sharedhttp.NewHTTPServer(cfg.HTTP, r.Router()),
	}, nil, nil
}

// Start runs the http API and quits doing a grateful shutdown.
// To stop the server you must send a syscall.SIGINT signal usually through `CTRL+C`.
func (h apiHTTP) Start() {
	go func() {
		err := h.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log().Fatal().Err(err).Msg("http server startup failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	h.shutdownApi()
}

func (h apiHTTP) shutdownApi() {
	ctx, cancel := context.WithTimeout(context.Background(), h.cfg.HTTP.ShutdownTimeout())
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		logger.Log().Error().Err(err).Msg("http server graceful shutdown failed")
	}

	logger.Log().Info().Msg("http api shutdown gracefully")
	os.Exit(0)
}
