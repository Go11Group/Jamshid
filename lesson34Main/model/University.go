package model

import (
	"github.com/google/uuid"
	"time"
)

type University struct {
	Id              uuid.UUID `json:"id"`
	UniversityName  string    `json:"university_name"`
	Rector          string    `json:"rector"`
	UniversityEmail string    `json:"university_email"`
	UniversityPhone string    `json:"university_phone"`
	CreatedAt       time.Time `json:"created_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
