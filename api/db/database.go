package db

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, err
}
