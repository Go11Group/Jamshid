package models

type Filter struct {
	Amount          int    `json:"amount"`
	Number          string `json:"number"`
	UserId          string `json:"user_id"`
	Name            string `json:"name"`
	StationId       string `json:"station_id"`
	CardId          string `json:"card_id"`
	TerminalId      string `json:"terminal_id"`
	TransactionType string `json:"transaction_type"`
	Limit, Offset   int
}
