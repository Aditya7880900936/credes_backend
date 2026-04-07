// @title Task Management API
// @version 1.0
// @description Backend API for task management system with RBAC, JWT, and soft delete
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"

	_ "github.com/Aditya7880900936/credes-backend/docs"
	"github.com/Aditya7880900936/credes-backend/internal/config"
	"github.com/Aditya7880900936/credes-backend/internal/db"
	"github.com/Aditya7880900936/credes-backend/internal/handler"
	"github.com/Aditya7880900936/credes-backend/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDB(cfg.DBUrl)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	auth.POST("/tasks/:id/comments", handler.AddComment)
	auth.GET("/tasks/:id/comments", handler.GetComments)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("admin"))

	admin.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin access granted"})
	})

	admin.POST("/tasks", handler.CreateTask)
	admin.DELETE("/tasks/:id", handler.DeleteTask)

	admin.PATCH("/users/:id/soft-delete", handler.SoftDeleteUser)

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
