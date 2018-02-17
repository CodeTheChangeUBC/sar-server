package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDB opens the given database
func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return nil, err
	}

	return db, nil
}
