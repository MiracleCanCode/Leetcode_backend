package main

import (
	"github.com/MiracleCanCode/zaperr"
	"github.com/clone_yandex_taxi/server/auth/cmd/api"
	"github.com/clone_yandex_taxi/server/auth/pkg/logger"
)

func main() {
	logger := logger.Logger()
	loggingErrors := zaperr.NewZaperr(logger)
	api := api.NewApi(logger, loggingErrors)
	api.FillEndpoints()

	if err := api.Run(); err != nil {
		logger.Error("Failed to run server: " + err.Error())
	}
}
