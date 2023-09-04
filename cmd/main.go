package main

import (
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
	"github.com/marcos-wz/capstone-go-bootcamp/pkg/app"
)

func main() {
	api, clean, err := app.NewApiHTTP()
	if err != nil {
		logger.Log().Fatal().Err(err).Msg("http rest api startup failed")
	}
	defer clean()

	api.Start()
}
