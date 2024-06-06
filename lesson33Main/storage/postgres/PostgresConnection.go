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
	dbname   = "gollang"
)

func ConnectionDb() (*sql.DB, error) {
	DbConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", DbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	return db, nil

}
