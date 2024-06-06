package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type Faculty struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	GroupCount   int       `json:"group_count"`
	Decant       string    `json:"decant"`
	FacultyEmail string    `json:"faculty_email"`
	FacultyPhone string    `json:"faculty_phone"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
