package solutionvalidator

type SolutionValidatorRequest struct {
	Lang       string `json:"lang" validator:"required"`
	Code       string `json:"code" validator:"required"`
	SolutionId int    `json:"solution_id" validator:"required"`
}
