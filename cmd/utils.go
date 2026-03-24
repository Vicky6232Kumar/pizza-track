package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBPath string
}

func loadConfig() Config {
	envFiles := []string{".env", "cmd/.env"}
	loaded := false
	for _, envFile := range envFiles {
		if err := godotenv.Load(envFile); err == nil {
			loaded = true
			break
		}
	}
	if !loaded {
		slog.Warn(".env file not loaded, falling back to system env/defaults")
	}

	return Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DATABASE_URL", "./data"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	slog.Warn("Environment variable not set, using default", "key", key, "default", defaultValue)

	return defaultValue
}
