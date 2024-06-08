package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UserProblem uuid.UUID `json:"user_problem"`
}
