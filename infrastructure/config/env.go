package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port              string
	DatabaseURL       *string
	JwtSecret         string
	JwtExpirationTime string
}

func NewEnv() Env {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "DEV"
	}

	godotenv.Load(envFilePath(env))
	defaultConfig := Env{
		Port:        ":8080",
		DatabaseURL: nil,
	}

	port := os.Getenv("PORT")
	if port != "" {
		defaultConfig.Port = ":" + port
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		defaultConfig.DatabaseURL = &databaseURL
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret != "" {
		defaultConfig.JwtSecret = jwtSecret
	}

	jwtExpirationTime := os.Getenv("JWT_EXPIRATION_TIME")
	if jwtExpirationTime != "" {
		defaultConfig.JwtExpirationTime = jwtExpirationTime
	}

	return defaultConfig
}

func envFilePath(env string) string {
	switch env {
	case "LOCAL":
		return ".env.local"
	case "DEV":
		return ".env.development.local"
	case "PROD":
		return ".env.production.local"
	default:
		return ".env"
	}
}
