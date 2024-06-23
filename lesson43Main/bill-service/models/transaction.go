package models

import "time"

type Transaction struct {
	Id              string    `json:"id"`
	CardId          string    `json:"card_id"`
	Amount          int       `json:"amount"`
	TerminalId      string    `json:"terminal_id"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       int       `json:"deleted_at"`
}
