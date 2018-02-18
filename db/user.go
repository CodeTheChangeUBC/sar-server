package db

import "time"

const usersSchema = `
PRAGMA foreign_keys = ON;

BEGIN;

CREATE TABLE users (
	id INTEGER PRIMARY KEY,
	name TEXT,
	joined datetime
);

CREATE TABLE coordinate (
	user INTEGER,
	latitude TEXT,
	longitude TEXT,

	FOREIGN KEY(user) REFERENCES users(id) DEFERRABLE INITIALLY DEFERRED
);

COMMIT;
`

// A User is simply someone with a name.
type User struct {
	ID     int
	Name   string
	Joined time.Time
}

// GetUserByID gets the given user by its unique ID.
func GetUserByID(uid int) (User, error) {
	row, err := Database.QueryRow("SELECT * FROM users WHERE id = ?", uid)
	if err != nil {
		return User{}, err
	}

	user, err := readUser(row)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// A Coordinate is a set of the latitude, longitude and when the volunteer was
// at that position.
type Coordinate struct {
	When  time.Time
	Point Point
}

// Reads a user in from the given scanner
func readUser(sc scanner) (User, error) {
	var id int
	var name string
	var joined time.Time

	err := sc.Scan(&id, &name, &joined)
	if err != nil {
		return User{}, nil
	}

	user := User{
		ID:     id,
		Name:   name,
		Joined: joined,
	}

	return user, nil
}
