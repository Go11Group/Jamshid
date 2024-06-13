package model

import (
	"github.com/google/uuid"
	"time"
)

type Lesson struct {
	Id        uuid.UUID `json:"id"`
	CourseId  uuid.UUID `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}
