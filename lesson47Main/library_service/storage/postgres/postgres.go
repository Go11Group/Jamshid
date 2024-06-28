package postgres

import (
	"database/sql"
	"fmt"
	"library_service/config"
	"library_service/storage"

	"strconv"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db     *sql.DB
	user   strorage.UserStorage
	borrow strorage.BorrowStorage
	book   strorage.BookStorage
}

func ConnectDB() (*Storage, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		strconv.Itoa(cfg.PostgresPort),
		cfg.PostgresDatabase,
	)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil
}
