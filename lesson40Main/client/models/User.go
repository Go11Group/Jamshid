package models

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Name      string
	Age       int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
