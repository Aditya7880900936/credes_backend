package main

import (
	"log"

	"github.com/Aditya7880900936/credes-backend/internal/config"
	"github.com/Aditya7880900936/credes-backend/internal/db"
	"github.com/Aditya7880900936/credes-backend/internal/handler"
	"github.com/Aditya7880900936/credes-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDB(cfg.DBUrl)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/auth/register", handler.Register)
	r.POST("/auth/login", handler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		role, _ := c.Get("role")

		c.JSON(200, gin.H{
			"user_id": userID,
			"role":    role,
		})
	})

	auth.GET("/tasks", handler.GetTasks)
	auth.PATCH("/tasks/:id/status", handler.UpdateTaskStatus)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("admin"))

	admin.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin access granted"})
	})

	admin.POST("/tasks", handler.CreateTask)
	admin.DELETE("/tasks/:id", handler.DeleteTask)

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}