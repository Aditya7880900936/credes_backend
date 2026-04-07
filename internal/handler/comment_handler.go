package handler

import (
	"strconv"

	"github.com/Aditya7880900936/credes-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// AddComment godoc
// @Summary Add comment
// @Tags Comments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request body map[string]string true "Comment"
// @Success 201 {object} map[string]interface{}
// @Router /tasks/{id}/comments [post]
func AddComment(c *gin.Context) {
	taskID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	userID := c.GetInt64("user_id")
	role := c.GetString("role")

	var body struct {
		Text string `json:"text" binding:"required"`
	}
	c.ShouldBindJSON(&body)

	comment, err := service.AddComment(taskID, userID, role, body.Text)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, comment)
}

// GetComments godoc
// @Summary Get comments for a task
// @Tags Comments
// @Security BearerAuth
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {array} map[string]interface{}
// @Router /tasks/{id}/comments [get]

func GetComments(c *gin.Context) {
	taskID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	userID := c.GetInt64("user_id")
	role := c.GetString("role")

	comments, err := service.GetComments(taskID, userID, role)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}
