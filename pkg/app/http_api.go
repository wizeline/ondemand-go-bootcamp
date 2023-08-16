package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/controller"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/customtype"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/sharedhttp"
)

type ApiHTTP interface {
	Start()
}

type apiHTTP struct {
	cfg    configuration.Config
	server *http.Server
}

func NewApiHTTP() (ApiHTTP, func(), error) {
	cfg := configuration.GetInstance()

	// TODO: implement the validating of data file path
	//filePath := fmt.Sprintf("../../%v/%v", dataDir, fileName)
	//if err := validateDataFilePath(filePath); err != nil {
	//	return nil, Err(err.Error())
	//}

	dbDrv := customtype.NewDriverDB(cfg.Database.Driver())
	if dbDrv == customtype.UndefinedDriverDB {
		return nil, nil, fmt.Errorf("database driver is not supported: %v", cfg.Database.Driver())
	}
	logger.Log().Debug().Str("type", dbDrv.String()).Msg("database driver")

	// Dependencies
	fruitRepo := repository.NewFruitCsv(cfg.Database.CSV)
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
