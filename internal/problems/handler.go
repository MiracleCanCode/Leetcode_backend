package problems

import (
	"net/http"
	"strconv"

	"github.com/clone_yandex_taxi/server/auth/pkg/db"
	jsondecoderandencoder "github.com/clone_yandex_taxi/server/auth/pkg/jsonDecoderAndEncoder"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	router  *mux.Router
	service *Service
}

func NewHandler(logger *zap.Logger, router *mux.Router, db *db.Db) {
	handler := &Handler{
		logger:  logger,
		router:  router,
		service: NewService(db, logger),
	}

	handler.router.HandleFunc("/api/problems/create", handler.Create()).Methods("POST")
	handler.router.HandleFunc("/api/problems/getById/{id}", handler.GetById()).Methods("GET")
	handler.router.HandleFunc("/api/problems/getAll", handler.GetAll()).Methods("GET")
}

func (s *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var payload *CreateRequest
		jsonDecodeAndEncode := jsondecoderandencoder.New(r, w)

		if err := jsonDecodeAndEncode.Decode(&payload); err != nil {
			s.logger.Error("Failed decode body, error:" + err.Error())
			http.Error(w, "Failed decode body", http.StatusBadRequest)
			return
		}

		parseToProblemModel := ToProblemModel(payload)
		err := s.service.Create(parseToProblemModel)

		if err != nil {
			s.logger.Error("Failed encode body, error:" + err.Error())
			http.Error(w, "Failed encode body", http.StatusBadRequest)
			return
		}

		w.Write([]byte("Success create problem"))
	}
}

func (s *Handler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		jsonDecodeAndEncode := jsondecoderandencoder.New(r, w)
		problemId := mux.Vars(r)
		parseIdToUint, err := strconv.ParseUint(problemId["id"], 0, 64)
		if err != nil {
			s.logger.Error("Failed parse id, error:" + err.Error())
			http.Error(w, "Failed parse id", http.StatusBadRequest)
			return
		}

		problem, err := s.service.GetById(uint(parseIdToUint))
		if err != nil {
			s.logger.Error("Failed get problem, error:" + err.Error())
			http.Error(w, "Failed get problem", http.StatusBadRequest)
			return
		}

		if err := jsonDecodeAndEncode.Encode(&problem); err != nil {
			s.logger.Error("Failed encode problem, error:" + err.Error())
			http.Error(w, "Failed encode problem", http.StatusBadRequest)
			return
		}
	}
}

func (s *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var payload *GetAllRequest
		jsonDecodeAndEncode := jsondecoderandencoder.New(r, w)

		if err := jsonDecodeAndEncode.Decode(&payload); err != nil {
			s.logger.Error("Failed decode body, error:" + err.Error())
			http.Error(w, "Failed decode body", http.StatusBadRequest)
			return
		}

		problems, err := s.service.GetAll(int(payload.Limit), int(payload.Offset))
		if err != nil {
			s.logger.Error("Failed get problems, error:" + err.Error())
			http.Error(w, "Failed get problems", http.StatusBadRequest)
			return
		}

		if err := jsonDecodeAndEncode.Encode(problems); err != nil {
			s.logger.Error("Failed encode problem, error:" + err.Error())
			http.Error(w, "Failed encode problem", http.StatusBadRequest)
			return
		}

	}
}
