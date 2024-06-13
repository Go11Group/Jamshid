package model

import (
	"time"
)

type Student struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Nation     string    `json:"nation"`
	Field      string    `json:"field"`
	ParentName string    `json:"parent_name"`
	City       string    `json:"city"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
