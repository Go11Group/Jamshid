package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Birthday    string    `json:"birthday"`
	Password    string    `json:"password"`
	Courses     *Course
	Enrollments Enrollment
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int       `json:"deleted_at"`
}
