package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/mux"
	"time"
)

type Problem struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
