package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	ConnToDbStr string
	Port        string
}

func New(logger *zap.Logger) *Config {
	if err := godotenv.Load(".env.local"); err != nil {
		logger.Error("Failed load .env.local file:" + err.Error())
	}

	connStr := os.Getenv("DB")
	port := os.Getenv("PORT")

	return &Config{
		ConnToDbStr: connStr,
		Port:        port,
	}
}
