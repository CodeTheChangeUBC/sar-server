package db

import (
	"database/sql"
	"log"

	// To register the sqlite3 driver.
	_ "github.com/mattn/go-sqlite3"
)

// Database is the global instance used to access values
var Database = func() *sql.DB {
	db, err := OpenDB("./db.sqlite")
	if err != nil {
		panic(err)
	}

	return db
}()

func init() {
	var schemas = []string{
		tasksSchema,
		usersSchema,
	}

	for _, schema := range schemas {
		_, err := Database.Exec(schema)
		if err != nil {
			log.Fatalf("Error accessing DB. Please delete and try again or contact CTC. %v", err)
		}
	}
}

// OpenDB opens the given database
func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
