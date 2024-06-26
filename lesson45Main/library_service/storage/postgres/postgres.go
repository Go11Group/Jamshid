package main

import (
	"database/sql"
	"fmt"
	"library_service/config"
	"library_service/storage"
	"library_service/storage/postgres"

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

func (s *Storage) User() strorage.UserStorage {
	if s.user == nil {
		s.user = postgres.NewUserRepo(s.Db)
	}
	return s.user
}

func (s *Storage) Borrow() strorage.BorrowStorage {
	if s.borrow == nil {
		s.borrow = postgres.NewBorrowRepo(s.Db)
	}
	return s.borrow
}

func (s *Storage) Book() strorage.BookStorage {
	if s.book == nil {
		s.book = postgres.NewBookRepo(s.Db)
	}
	return s.book
}
