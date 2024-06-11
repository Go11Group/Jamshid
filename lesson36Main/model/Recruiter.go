package model

import (
	"github.com/google/uuid"
	"time"
)

type Recruiter struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Birthday    time.Time     `json:"birthday"`
	Gender      string        `json:"gender"`
	CompanyId   uuid.NullUUID `json:"company_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   string        `json:"deleted_at"`
}
