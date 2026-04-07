package handler

import (
	"net/http"
	"strconv"

	"github.com/Aditya7880900936/credes-backend/internal/config"
	"github.com/Aditya7880900936/credes-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.RegisterUser(req.Email, req.Password, req.FullName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func Login(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cfg := config.LoadConfig()

	token, err := service.LoginUser(req.Email, req.Password, cfg.JWTSecret)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func SoftDeleteUser(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := service.SoftDeleteUser(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "user deactivated"})
}
