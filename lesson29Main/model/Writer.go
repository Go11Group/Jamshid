package model

import "gorm.io/gorm"

type Writer struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}
