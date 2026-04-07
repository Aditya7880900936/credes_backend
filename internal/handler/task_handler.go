package handler

import (
	"strconv"

	"github.com/Aditya7880900936/credes-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	AssignedTo  int64  `json:"assigned_to" binding:"required"`
}

func CreateTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task, err := service.CreateTask(req.Title, req.Description, req.AssignedTo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, task)
}

func GetTasks(c *gin.Context) {
	userID := c.GetInt64("user_id")
	role := c.GetString("role")

	tasks, err := service.GetTasks(userID, role)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tasks)
}

func UpdateTaskStatus(c *gin.Context) {
	taskID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	userID := c.GetInt64("user_id")
	role := c.GetString("role")

	var body struct {
		Status string `json:"status"`
	}
	c.ShouldBindJSON(&body)

	err := service.UpdateTaskStatus(taskID, userID, role, body.Status)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "updated"})
}

func DeleteTask(c *gin.Context) {
	taskID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	role := c.GetString("role")

	err := service.DeleteTask(taskID, role)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "deleted"})
}
