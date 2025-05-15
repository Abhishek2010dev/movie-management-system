package config

import (
	"os"
)

type Config struct {
	DatabaseUrl string
	JwtSecret   string
}

func Load() *Config {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgresql://postgres@localhost:5432/movie?sslmode=disable"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "defaultsecret"
	}

	return &Config{
		DatabaseUrl: dbUrl,
		JwtSecret:   jwtSecret,
	}
}

