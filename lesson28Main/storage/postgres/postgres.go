package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1111"
	dbname   = "gollang"
)

func ConnectionDB() (*sql.DB, error) {
	con := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable  ", host, port, user, password, dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		fmt.Println("sa1", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("sa1")

		return nil, err

	}
	return db, nil

}
