package main

import (
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadEnv()

	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect ot database")
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":  "TodoAPI is running",
			"status":   "success",
			"database": "connected",
		})
	})

	router.Run(":" + cfg.Port)
}
