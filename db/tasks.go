package db

const tasksSchema = `
PRAGMA foreign_keys = ON;

BEGIN;

CREATE TABLE tasks (
	id INTEGER PRIMARY KEY,
	name TEXT,
	details TEXT,
	area_latitude_start TEXT,
	area_longitude_start TEXT,
	area_latitude_end TEXT,
	area_longitude_end TEXT,
	status TEXT
);

CREATE TABLE assignments (
	id INTEGER PRIMARY KEY,
	user_id INTEGER,
	task_id INTEGER,

	FOREIGN KEY (user_id) REFERENCES users(id) DEFERRABLE INITIALLY DEFERRED,
	FOREIGN KEY (task_id) REFERENCES tasks(id) DEFERRABLE INITIALLY DEFERRED
);

COMMIT;
`

// A Task includes an area to search and additional details.
type Task struct {
	StartCorner Point
	EndCorner   Point
	AssignedTo  []User
}
