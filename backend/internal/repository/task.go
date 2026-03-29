package repository

import (
	"github.com/sanyathecreator/task-manager/internal/db"
	"github.com/sanyathecreator/task-manager/internal/model"
)

func Save(t model.Task) error {
	query := `
	INSERT INTO tasks(title, completed, created_at)
	VALUES (?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Completed, t.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func GetTasks() ([]model.Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
