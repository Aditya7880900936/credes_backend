package repository

import (
	"github.com/Aditya7880900936/credes-backend/internal/db"
	"github.com/Aditya7880900936/credes-backend/internal/models"
)

func CreateTask(task *models.Task) error {
	query := `
	INSERT INTO tasks (title, description, assigned_to)
	VALUES ($1, $2, $3)
	RETURNING id, status, created_at, updated_at
	`

	return db.DB.QueryRowx(query,
		task.Title,
		task.Description,
		task.AssignedTo,
	).Scan(&task.ID, &task.Status, &task.CreatedAt, &task.UpdatedAt)
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	query := `SELECT * FROM tasks ORDER BY created_at DESC`

	err := db.DB.Select(&tasks, query)
	return tasks, err
}

func GetTasksByUser(userID int64) ([]models.Task, error) {
	var tasks []models.Task
	query := `SELECT * FROM tasks WHERE assigned_to=$1`

	err := db.DB.Select(&tasks, query, userID)
	return tasks, err
}

func UpdateTaskStatus(taskID int64, status string) error {
	query := `UPDATE tasks SET status=$1, updated_at=NOW() WHERE id=$2`
	_, err := db.DB.Exec(query, status, taskID)
	return err
}

func DeleteTask(taskID int64) error {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := db.DB.Exec(query, taskID)
	return err
}
