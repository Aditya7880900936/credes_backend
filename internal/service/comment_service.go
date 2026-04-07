package service

import (
	"errors"

	"github.com/Aditya7880900936/credes-backend/internal/models"
	"github.com/Aditya7880900936/credes-backend/internal/repository"
)

func AddComment(taskID, userID int64, role, text string) (*models.Comment, error) {

	if role != "admin" {
		tasks, _ := repository.GetTasksByUser(userID)

		allowed := false
		for _, t := range tasks {
			if t.ID == taskID {
				allowed = true
				break
			}
		}

		if !allowed {
			return nil, errors.New("not allowed to comment")
		}
	}

	comment := &models.Comment{
		TaskID:   taskID,
		AuthorID: userID,
		Text:     text,
	}

	err := repository.CreateComment(comment)
	return comment, err
}

func GetComments(taskID, userID int64, role string) ([]models.Comment, error) {
	if role == "admin" {
		return repository.GetCommentsByTask(taskID)
	}

	tasks, _ := repository.GetTasksByUser(userID)
	for _, t := range tasks {
		if t.ID == taskID {
			return repository.GetCommentsByTask(taskID)
		}
	}

	return nil, errors.New("not allowed")
}