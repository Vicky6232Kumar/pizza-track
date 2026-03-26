package main

import (
	"log/slog"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	Port              string
	DBPath            string
	SessionSecrectKey string
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
		Port:              getEnv("PORT", "8080"),
		DBPath:            getEnv("DATABASE_URL", "./data"),
		SessionSecrectKey: getEnv("SESSION_SECRECT_KEY", "pizza-order-secrect-key"),
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

func setupSessionStore(db *gorm.DB, secrectKey []byte) sessions.Store {
	store := gormsessions.NewStore(db, true, secrectKey)
	store.Options(sessions.Options{
		Path:     "",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
		SameSite: 3,
	})

	return store
}

func SetSessionValue(c *gin.Context, key string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(key, value)
	return session.Save()
}

func GetSessionString(c *gin.Context, key string) string {
	session := sessions.Default(c)
	val := session.Get(key)
	if val == nil {
		return ""
	}

	str, _ := val.(string)
	return str
}

func ClearSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}
