package task

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type TaskHandler struct {
	logger *zap.Logger
	router *mux.Router
}

func NewTaskHandler(logger *zap.Logger, router *mux.Router) {
	handler := &TaskHandler{
		logger: logger,
		router: router,
	}

	handler.router.HandleFunc("/api/createTask", handler.CreateTask()).Methods("POST")
	handler.router.HandleFunc("/api/getTaskById", handler.GetTaskById()).Methods("POST")
}

func (s *TaskHandler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s *TaskHandler) GetTaskById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
