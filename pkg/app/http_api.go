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

// ApiHTTP is an HTTP API interface that runs a http.Server with mux.Router support.
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
// Moreover, implements the http.Server instance with sharedhttp.Mux router.
// Returns an error if any of the dependency implementation gets failed.
func NewApiHTTP() (ApiHTTP, func(), error) {
	// Dependencies
	cfg := configuration.GetInstance()
	fruitRepo, err := repository.NewFruitCsv(cfg.Database.CSV)
	if err != nil {
		return nil, nil, err
	}
	fruitSvc := service.NewFruit(fruitRepo)

	// Routes
	r := sharedhttp.NewMux(cfg.AppVersion)
	r.Add("HealthCheck", controller.NewHealthCheckHTTP())
	r.Add("Home", controller.NewHomeHTTP())
	r.Add("Fruit", controller.NewFruitHTTP(fruitSvc))
	r.RegisterRoutes()

	return &apiHTTP{
		cfg:    cfg,
		server: sharedhttp.NewHTTPServer(cfg, r.Router()),
	}, nil, nil
}

// Start runs the http API and quits doing a grateful shutdown.
// To stop the server you must send a system interrupt signal usually through the `CTRL+C` command
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
