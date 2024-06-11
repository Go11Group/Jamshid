package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   string    `json:"deleted_at"`
}
