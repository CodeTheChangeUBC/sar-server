package db

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

// A Task includes an area to search and additional details.
type Task struct {
	ID          int
	Name        string
	Details     string
	Status      string
	StartCorner Point
	EndCorner   Point
	AssignedTo  []User
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
		var id int
		var name string
		var details string
		var status string
		var laStart string
		var laEnd string
		var loStart string
		var loEnd string
		err = rows.Scan(&id, &name, &details, &status, &laStart, &loStart, &laEnd, &loEnd)
		if err != nil {
			return nil, err
		}
		newTask := Task{
			ID:          id,
			Name:        name,
			Details:     details,
			Status:      status,
			StartCorner: Point{Latitude: laStart, Longitude: loStart},
			EndCorner:   Point{Latitude: laEnd, Longitude: loEnd},
			AssignedTo:  []User{},
		}
		tasks = append(tasks, newTask)
	}

	return tasks, nil
}
