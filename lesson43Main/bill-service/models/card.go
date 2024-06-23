package models

import "time"

type Card struct {
	Id        string    `json:"id"`
	Number    string    `json:"number"`
	Userid    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}
