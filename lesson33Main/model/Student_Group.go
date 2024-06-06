package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type Student_Group struct {
	Id         uuid.UUID `json:"id"`
	Student_id uuid.UUID `json:"student_Id"`
	Group_id   uuid.UUID `json:"group_id"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
