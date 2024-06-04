package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "1111"
	port     = 5433
	dbName   = "gollang"
)

func ConnectionDb() (*sql.DB, error) {
	DbConn := fmt.Sprintf("host=%s port=%d password=%s  dbname=%s user=%s sslmode=disable", host, port, password, dbName, user)
	db, err := sql.Open("postgres", DbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
