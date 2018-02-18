package db

import "errors"

const tasksSchema = `
PRAGMA foreign_keys = ON;

BEGIN;

CREATE TABLE tasks (
	id INTEGER PRIMARY KEY,
	name TEXT,
	details TEXT,
	status TEXT
	area_latitude_start TEXT,
	area_longitude_start TEXT,
	area_latitude_end TEXT,
	area_longitude_end TEXT,
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

var (
	// ErrLTZeroUID is returned when the UID for that user is less than zero.
	ErrLTZeroUID = errors.New("UID was less than zero")
)

// A Task includes an area to search and additional details.
type Task struct {
	ID          int
	Name        string
	Details     string
	Status      string
	StartCorner Point
	EndCorner   Point
}

// GetUserTasks gets the given user's assigned tasks.
func GetUserTasks(user User) ([]Task, error) {
	return GetUserIDTasks(user.ID)
}

// GetUserIDTasks gets the given user's assigned tasks.
func GetUserIDTasks(uid int) ([]Task, error) {
	if uid < 0 {
		return nil, ErrLTZeroUID
	}

	taskIDs, err := Database.Query("SELECT task_id FROM assignments WHERE user_id = ?", uid)
	if err != nil {
		return nil, err
	}

	defer taskIDs.Close()

	tasks := make([]Task, 0, 8)
	for taskIDs.Next() {
		var id int
		err = taskIDs.Scan(&id)
		if err != nil {
			return nil, err
		}

		row := Database.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		task, err := readTask(row)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetAllTasks does exactly what its name implies
func GetAllTasks() ([]Task, error) {
	rows, err := Database.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		newTask, err := readTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, newTask)
	}

	return tasks, nil
}

type scanner interface {
	Scan(args ...interface{}) error
}

func readTask(sc scanner) (Task, error) {
	var id int
	var name string
	var details string
	var status string
	var laStart string
	var laEnd string
	var loStart string
	var loEnd string
	err := sc.Scan(&id, &name, &details, &status, &laStart, &loStart, &laEnd, &loEnd)

	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		ID:          id,
		Name:        name,
		Details:     details,
		Status:      status,
		StartCorner: Point{Latitude: laStart, Longitude: loStart},
		EndCorner:   Point{Latitude: laEnd, Longitude: loEnd},
	}

	return newTask, nil
}
