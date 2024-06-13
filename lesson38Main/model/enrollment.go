package model

import (
	"github.com/google/uuid"
	"time"
)

type Enrollment struct {
	Id             uuid.UUID `json:"id"`
	UserId         uuid.UUID `json:"user_id"`
	CourseId       uuid.UUID `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      int       `json:"deleted_at"`
}
