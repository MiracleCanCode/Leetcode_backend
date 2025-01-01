package main

import (
	logger "github.com/MiracleCanCode/example_configuration_logger"
	"github.com/clone_yandex_taxi/server/auth/config"
	"github.com/clone_yandex_taxi/server/auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log := logger.Logger(logger.DefaultLoggerConfig())

	db, err := gorm.Open(postgres.Open(config.New(log).ConnToDbStr), &gorm.Config{})
	if err != nil {
		log.Error("Failed open db," + err.Error())
	}


	db.AutoMigrate(&models.Problem{})
}