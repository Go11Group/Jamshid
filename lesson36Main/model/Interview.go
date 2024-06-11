package model

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type Interview struct {
	Id            uuid.UUID     `json:"id"`
	UserId        uuid.NullUUID `json:"user_id"`
	VacancyId     uuid.NullUUID `json:"vacancy_id"`
	RecruiterId   uuid.NullUUID `json:"recruiter_id"`
	InterviewDate time.Time     `json:"interview_date"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	DeletedAt     string        `json:"deleted_at"`
}
