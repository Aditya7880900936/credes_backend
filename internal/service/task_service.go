package service

import (
	"errors"

	"github.com/Aditya7880900936/credes-backend/internal/models"
	"github.com/Aditya7880900936/credes-backend/internal/repository"
)

func CreateTask(title, desc string, assignedTo int64) (*models.Task, error) {
	task := &models.Task{
		Title:       title,
		Description: desc,
		AssignedTo:  assignedTo,
	}

	err := repository.CreateTask(task)
	return task, err
}

func GetTasks(userID int64, role string) ([]models.Task, error) {
	if role == "admin" {
		return repository.GetAllTasks()
	}
	return repository.GetTasksByUser(userID)
}

func UpdateTaskStatus(taskID int64, userID int64, role string, status string) error {
	if role == "admin" {
		return repository.UpdateTaskStatus(taskID, status)
	}

	tasks, _ := repository.GetTasksByUser(userID)
	for _, t := range tasks {
		if t.ID == taskID {
			return repository.UpdateTaskStatus(taskID, status)
		}
	}

	return errors.New("not allowed")
}

func DeleteTask(taskID int64, role string) error {
	if role != "admin" {
		return errors.New("only admin can delete")
	}
	return repository.DeleteTask(taskID)
}