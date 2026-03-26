package main

import (
	"log/slog"
	"os"
	"pizza-tracking/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Starting....")
	cfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialied database", "error", err)
		os.Exit(1)
	}

	slog.Info("Database connected successfully")

	h := NewHandler(dbModel)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	sessionStore := setupSessionStore(dbModel.DB, []byte(cfg.SessionSecrectKey))

	setUpRoutes(router, h, sessionStore)

	slog.Info("Server starting", "url", "http://localhost:"+cfg.Port)
	router.Run(":" + cfg.Port)
}
