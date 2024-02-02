package database

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("CATALOG_POSTGRES_USER"), os.Getenv("CATALOG_POSTGRES_PASSWORD"), os.Getenv("CATALOG_POSTGRES_DB"), os.Getenv("CATALOG_POSTGRES_PORT"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, err
}
