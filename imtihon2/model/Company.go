package model

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Workers   int       `json:"workers"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
}
