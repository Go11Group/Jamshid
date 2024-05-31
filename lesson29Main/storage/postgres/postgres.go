package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"lesson29Main/model"
	_ "lesson29Main/model"
)

func ConnectionDb() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=1111 dbname=gollang port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	writer := model.Writer{}
	err = db.AutoMigrate(&writer)
	if err != nil {

		return nil, err
	}

	if db == nil {
		return db, nil
	}
	return db, nil
}
