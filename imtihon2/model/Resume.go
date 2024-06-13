package model

import (
	"github.com/google/uuid"
	"time"
)

type Resume struct {
	Id          uuid.NullUUID `json:"id"`
	Position    string        `json:"position"`
	Experience  int           `json:"experience"`
	Description string        `json:"description"`
	UserID      uuid.NullUUID `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   string        `json:"deleted_at"`
}
