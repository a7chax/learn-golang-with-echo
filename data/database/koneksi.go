package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabaseNote() (*sql.DB, error) {
	connectionString := "host=localhost user=postgres password=12345 dbname=template1 sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
