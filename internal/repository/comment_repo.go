package repository

import (
	"github.com/Aditya7880900936/credes-backend/internal/db"
	"github.com/Aditya7880900936/credes-backend/internal/models"
)

func CreateComment(cmt *models.Comment) error {
	query := `
	INSERT INTO comments (task_id, author_id, text)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`
	return db.DB.QueryRowx(query,
		cmt.TaskID,
		cmt.AuthorID,
		cmt.Text,
	).Scan(&cmt.ID, &cmt.CreatedAt)
}

func GetCommentsByTask(taskID int64) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT * FROM comments WHERE task_id=$1 ORDER BY created_at DESC`
	err := db.DB.Select(&comments, query, taskID)
	return comments, err
}