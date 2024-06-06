package model

import "github.com/google/uuid"

type Faculty_University struct {
	Id           uuid.UUID `json:"id"`
	Faculty_id   uuid.UUID `json:"faculty_id"`
	Universities uuid.UUID `json:"universities"`
}
