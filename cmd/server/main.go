package main

import (
	"log"

	"github.com/Aditya7880900936/credes-backend/internal/config"
	"github.com/Aditya7880900936/credes-backend/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDB(cfg.DBUrl)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
