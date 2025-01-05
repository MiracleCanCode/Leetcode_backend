package solutionvalidator

import (
	"net/http"

	"github.com/clone_yandex_taxi/server/auth/pkg/db/postgresql"
	json "github.com/clone_yandex_taxi/server/auth/pkg/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	service *Service
}

func New(router *mux.Router, logger *zap.Logger, db *postgresql.Db) {
	handler := &Handler{
		logger:  logger,
		service: NewService(logger, db),
	}

	router.HandleFunc("/api/solution/validation", handler.Validate()).Methods("POST")
}

func (s *Handler) Validate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload *RequestPayload
		json := json.New(r, w)

		if err := json.Decode(payload); err != nil {
			s.logger.Error("Error decode response body error")
			http.Error(w, "Error decode response body error", http.StatusBadRequest)
		}

		code := s.service.Compile(payload)

		if err := json.Encode(code); err != nil {
			return
		}
	}
}
