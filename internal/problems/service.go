package problems

import (
	"github.com/clone_yandex_taxi/server/auth/internal/models"
	"github.com/clone_yandex_taxi/server/auth/pkg/db"
	"go.uber.org/zap"
)

type Service struct {
	repository *Repository
	db         *db.Db
	logger     *zap.Logger
}

func NewService(db *db.Db, log *zap.Logger) *Service {
	return &Service{
		db:         db,
		logger:     log,
		repository: NewRepository(db, log),
	}
}

func (s *Service) Create(data *models.Problem) error {
	return s.repository.Create(data)
}

func (s *Service) GetById(id uint) (*models.Problem, error) {
	return s.repository.GetById(id)
}

func (s *Service) GetAll(limit int, offset int) ([]*models.Problem, error) {
	return s.repository.GetAll(limit, offset)
}
