package config

import "os"

type Config struct {
	DatabaseUrl string
	JwtSecret   []byte
}

func Load() *Config {
	dbUrl, jwtSecret := os.Getenv("DATABASE_URL"), os.Getenv("JWT_SECRET")

	if dbUrl == "" {
		dbUrl = "postgresql://postgres@localhost:5432/movie?sslmode=disable"
	}
	if jwtSecret == "" {
		jwtSecret = "defaultsecret"
	}

	return &Config{DatabaseUrl: dbUrl, JwtSecret: []byte(jwtSecret)}
}

