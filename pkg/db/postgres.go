package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("BOOK_APP_HOST"),
		os.Getenv("BOOK_APP_PORT"),
		os.Getenv("BOOK_APP_USERNAME"),
		os.Getenv("BOOK_APP_NAME"),
		os.Getenv("BOOK_APP_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
