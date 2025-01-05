package main

import (
	logger "github.com/MiracleCanCode/example_configuration_logger"
	"github.com/clone_yandex_taxi/server/auth/config"
	"github.com/clone_yandex_taxi/server/auth/pkg/db/postgresql"
	"github.com/clone_yandex_taxi/server/auth/pkg/server"
)

func main() {
	log := logger.Logger(logger.DefaultLoggerConfig())
	conf := config.New(log)
	initDb := postgresql.NewDb(conf, log)
	api := server.New(log, initDb)
	api.FillEndpoints()

	if err := api.Run(); err != nil {
		log.Error("Failed to run server: " + err.Error())
	}
}
