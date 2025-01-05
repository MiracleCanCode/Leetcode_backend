package postgresql

import (
	"github.com/clone_yandex_taxi/server/auth/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *config.Config, log *zap.Logger) *Db {
	db, err := gorm.Open(postgres.Open(conf.ConnToDbStr), &gorm.Config{})

	if err != nil {
		log.Error("Failed conn to db: " + err.Error())
		return nil
	}

	return &Db{
		db,
	}
}
