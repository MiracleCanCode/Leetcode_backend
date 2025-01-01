package problems

import (
	"errors"

	"github.com/clone_yandex_taxi/server/auth/internal/models"
	"github.com/clone_yandex_taxi/server/auth/pkg/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IRepository interface {
	GetAll() ([]*models.Problem, error)
	GetById(id uint) (*models.Problem, error)
	Create(data *models.Problem) error
}

type Repository struct {
	db     *db.Db
	logger *zap.Logger
}

func NewRepository(db *db.Db, log *zap.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: log,
	}
}

func (s *Repository) Create(data *models.Problem) error {
	err := s.db.Create(&data).Error
	if err != nil {
		s.logger.Error("Failed to create problem", zap.Error(err))
		return errors.New("failed to create problem")
	}
	return nil
}

func (s *Repository) GetById(id uint) (*models.Problem, error) {
	var problem models.Problem
	err := s.db.First(&problem, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Warn("Problem not found", zap.Uint("id", id))
			return nil, nil
		}
		s.logger.Error("Failed to retrieve problem", zap.Uint("id", id), zap.Error(err))
		return nil, errors.New("failed to retrieve problem")
	}
	return &problem, nil
}

func (s *Repository) GetAll(limit int, offset int) ([]*models.Problem, error) {
	var problems []*models.Problem
	err := s.db.Limit(limit).Offset(offset).Find(&problems).Error
	if err != nil {
		s.logger.Error("Failed to retrieve all problems", zap.Error(err))
		return nil, errors.New("failed to retrieve problems")
	}
	return problems, nil
}
