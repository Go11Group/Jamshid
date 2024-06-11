package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	dbname   = "gollang"
	password = "1111"
	user     = "postgres"
)

func ConnectionDB() (*sql.DB, error) {
	dbCon := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", dbCon)
	return db, err
}
