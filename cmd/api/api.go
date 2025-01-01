package api

import (
	"net/http"

	"github.com/clone_yandex_taxi/server/auth/config"
	"github.com/clone_yandex_taxi/server/auth/internal/problems"
	"github.com/clone_yandex_taxi/server/auth/pkg/db"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type api struct {
	addr   string
	router *mux.Router
	logger *zap.Logger
	db     *db.Db
}

func New(logger *zap.Logger, db *db.Db) *api {
	cfg := config.New(logger)

	return &api{
		addr:   cfg.Port,
		router: mux.NewRouter(),
		logger: logger,
		db:     db,
	}
}

func (s *api) Run() error {
	s.logger.Info("Start server on http://localhost" + s.addr)
	return http.ListenAndServe(s.addr, s.router)
}

func (s *api) FillEndpoints() {
	problems.NewHandler(s.logger, s.router, s.db)
}
