package postgres

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	dbname   = "gollang"
	password = "1111"
)

func ConnectionDb() (*sql.DB, error) {
	conDB := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable ", host, port, user, dbname, password)
	db, err := sql.Open("postgres", conDB)
	if err != nil {
		return nil, err
	}
	return db, nil
}
