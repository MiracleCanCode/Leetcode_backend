package solutionvalidator

import (
	"os"
	"os/exec"

	"github.com/clone_yandex_taxi/server/auth/internal/problems"
	"github.com/clone_yandex_taxi/server/auth/pkg/db/postgresql"
	"go.uber.org/zap"
)

type Service struct {
	logger         *zap.Logger
	problemService *problems.Service
}

func NewService(logger *zap.Logger, db *postgresql.Db) *Service {
	return &Service{
		logger:         logger,
		problemService: problems.NewService(db, logger),
	}
}

func (s *Service) Compile(data *RequestPayload) (stdr string) {
	nameFile := "compile." + data.Lang
	_, _ = s.getInputAndOutput(uint(data.ProblemId))

	createFile, err := os.Create(nameFile)
	if err != nil {
		s.logger.Error("Failed to create file, error: " + err.Error())
		return
	}
	defer createFile.Close()

	_, err = createFile.Write([]byte(data.Code))
	if err != nil {
		s.logger.Error("Failed to write code, error: " + err.Error())
		return
	}

	cmd := exec.Command("go", "build", nameFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		s.logger.Error("Failed to build file, error: " + err.Error())
		return
	}

	executableName := nameFile[:len(nameFile)-len(data.Lang)-1]
	runCmd := exec.Command("./" + executableName)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	// runCmd.Stderr = bytes.NewBufferString(input)

	if err := runCmd.Run(); err != nil {
		s.logger.Error("Failed to run the executable, error: " + err.Error())
		return
	}

	return nameFile
}

func (s *Service) getInputAndOutput(problemId uint) (input string, output string) {
	getProblem, err := s.problemService.GetById(problemId)
	if err != nil {
		s.logger.Error("Failed to get problem by id, error:" + err.Error())
		return "", ""
	}

	return getProblem.Input, getProblem.Output
}
