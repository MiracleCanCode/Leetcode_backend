package solutionvalidator

import (
	"net/http"

	"github.com/MiracleCanCode/zaperr"
	jsondecoderandencoder "github.com/clone_yandex_taxi/server/auth/pkg/jsonDecoderAndEncoder"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	logger        *zap.Logger
	loggingErrors *zaperr.Zaperr
}

func New(router *mux.Router, logger *zap.Logger, loggingErrors *zaperr.Zaperr) {
	handler := &Handler{
		logger:        logger,
		loggingErrors: loggingErrors,
	}

	router.HandleFunc("/api/solution/validation", handler.Validate()).Methods("POST")
}

func (s *Handler) Validate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload *SolutionValidatorRequest
		json := jsondecoderandencoder.New(r, w)

		s.loggingErrors.LogError(json.Decode(payload), "Error decode response body error")

	}
}
