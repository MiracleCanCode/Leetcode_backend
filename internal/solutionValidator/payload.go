package solutionvalidator

type RequestPayload struct {
	Lang      string `json:"lang" validator:"required"`
	Code      string `json:"code" validator:"required"`
	ProblemId int    `json:"problem_id" validator:"required"`
}
