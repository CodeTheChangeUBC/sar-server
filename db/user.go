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
	Name   string
	Joined time.Time
}

// A Coordinate is a set of the latitude, longitude and when the volunteer was
// at that position.
type Coordinate struct {
	When  time.Time
	Point Point
}
