package db

import (
	"database/sql"
	"log"

	// To register the sqlite3 driver.
	_ "github.com/mattn/go-sqlite3"
)

// Database is the global instance used to access values
var Database *sql.DB

// InitializeDB runs creates the DB and runs "migrations"
func InitializeDB(path string) (*sql.DB, error) {
	db, err := OpenDB(path)

	if err != nil {
		return nil, err
	}

	var schemas = []string{
		tasksSchema,
		usersSchema,
	}

	// TODO check if schemas have been run and if they have skip execution aka.
	// Write a migration manager :/

	for _, schema := range schemas {
		_, err := db.Exec(schema)
		if err != nil {
			log.Fatalf("Error accessing DB. Please delete and try again or contact CTC. %v", err)
		}
	}

	return db, nil
}

type scanner interface {
	Scan(args ...interface{}) error
}

// OpenDB opens the given database
func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
