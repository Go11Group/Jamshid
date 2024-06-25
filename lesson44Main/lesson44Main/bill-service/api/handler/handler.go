package handler

import (
	"bill_service/storage/postgres"
)

type HTTPHandler struct {
	CardRepo        *postgres.CardRepository
	StationRepo     *postgres.StationRepository
	TerminalRepo    *postgres.TerminalRepository
	TransactionRepo *postgres.TransactionRepository
}

func NewHTTPHandler(cr *postgres.CardRepository, sr *postgres.StationRepository, te *postgres.TerminalRepository, tr *postgres.TransactionRepository) *HTTPHandler {
	return &HTTPHandler{
		CardRepo:        cr,
		StationRepo:     sr,
		TerminalRepo:    te,
		TransactionRepo: tr,
	}
}
