package models

import "time"

type Terminal struct {
	Id        string    `json:"id"`
	StationId string    `json:"station_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}
