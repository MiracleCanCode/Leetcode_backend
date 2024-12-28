package api

import (
	"net/http"

	"github.com/MiracleCanCode/zaperr"
	"github.com/clone_yandex_taxi/server/auth/config"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type api struct {
	addr          string
	router        *mux.Router
	logger        *zap.Logger
	loggingErrors *zaperr.Zaperr
}

func NewApi(logger *zap.Logger, loggingErrors *zaperr.Zaperr) *api {
	cfg := config.NewConfig(logger)

	return &api{
		addr:          cfg.Port,
		router:        mux.NewRouter(),
		logger:        logger,
		loggingErrors: loggingErrors,
	}
}

func (s *api) Run() error {
	s.logger.Info("Start server on http://localhost" + s.addr)
	return http.ListenAndServe(s.addr, s.router)
}

func (s *api) FillEndpoints() {

}
