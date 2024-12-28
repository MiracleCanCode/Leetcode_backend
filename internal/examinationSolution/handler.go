package examinationsolution

import (
	"net/http"

	"github.com/MiracleCanCode/zaperr"
	jsondecoderandencoder "github.com/clone_yandex_taxi/server/auth/pkg/jsonDecoderAndEncoder"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ExaminationSolutionHandler struct {
	logger        *zap.Logger
	loggingErrors *zaperr.Zaperr
}

func NewExaminationSolutionHandler(router *mux.Router, logger *zap.Logger, loggingErrors *zaperr.Zaperr) {
	handler := &ExaminationSolutionHandler{
		logger:        logger,
		loggingErrors: loggingErrors,
	}

	router.HandleFunc("/api/examination", handler.Examination()).Methods("POST")
}

func (s *ExaminationSolutionHandler) Examination() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload *ExaminationsolutionRequest
		json := jsondecoderandencoder.NewJsonDecoderAndEncoder(r, w)

		s.loggingErrors.LogError(json.Decode(payload), "Error decode response body error")

	}
}
