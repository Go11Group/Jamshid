package model

import (
	"github.com/google/uuid"
	"time"
)

type Vacancy struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Position      string    `json:"position"`
	MinExperience uint      `json:"min_experience"`
	CompanyID     uuid.UUID `json:"company_id"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     string    `json:"deleted_at"`
}
