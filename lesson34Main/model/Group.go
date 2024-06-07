package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type Group struct {
	Id            uuid.UUID `json:"id"`
	GroupName     string    `json:"group_name"`
	Student_count int       `json:"student_number"`
	CreatedAt     time.Time `json:"created_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
