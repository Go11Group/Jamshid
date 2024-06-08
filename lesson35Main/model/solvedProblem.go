package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type SolvedProblem struct {
	Id              uuid.UUID `json:"id"`
	Solution        string    `json:"solution"`
	ProblemSolution string    `json:"problem_solution"`
	UserSolution    string    `json:"user_solution"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
