package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	Port              string
	DatabaseURL       *string
	JwtSecret         string
	JwtExpirationTime int
}

func NewEnv() Env {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "DEV"
	}

	envFilePath := envFilePath(env)
	if envFilePath != nil {
		godotenv.Load(*envFilePath)
	} else {
		godotenv.Load()
	}

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
		expirationTime, err := strconv.Atoi(jwtExpirationTime)
		if err != nil {
			log.Fatal(err)
		}
		defaultConfig.JwtExpirationTime = expirationTime
	}

	return defaultConfig
}

func envFilePath(env string) *string {
	var path string
	switch env {
	case "LOCAL":
		path = ".env.local"
		return &path
	case "DEV":
		path = ".env.development.local"
		return &path
	case "PROD":
		path = ".env.production.local"
		return &path
	default:
		return nil
	}
}
