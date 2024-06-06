package model

import "github.com/google/uuid"

type GroupFaculty struct {
	Id         uuid.UUID `json:"id"`
	Groups_id  uuid.UUID `json:"groups_id"`
	Faculty_id uuid.UUID `json:"faculty_id"`
}
