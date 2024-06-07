package model

import (
	"github.com/google/uuid"
	"time"
)

type Student struct {
	Id          uuid.UUID `gorm:"type:uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
