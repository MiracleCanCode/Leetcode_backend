package problems

import "github.com/clone_yandex_taxi/server/auth/internal/models"

func ToProblemModel(data *CreateRequest) *models.Problem {
	return &models.Problem{
		Name:        data.Name,
		Description: data.Description,
		Input:       data.Input,
		Output:      data.Output,
	}
}

type CreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Input       string `json:"input" validate:"required"`
	Output      string `json:"output" validate:"required"`
}

type GetByIdRequest struct {
	Id uint `json:"id" validate:"required"`
}

type GetAllRequest struct {
	Limit  uint `json:"limit" validate:"required"`
	Offset uint `json:"offset" validate:"required"`
}
