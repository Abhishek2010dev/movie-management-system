package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func Connect(url string) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed to open connection: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping connection: %s", err)
	}

	log.Printf("Successfully connected to postgresql")

	if err := goose.Up(db.DB, "./migrations"); err != nil {
		log.Fatalf("failed to run migration: %s", err)
	}

	return db
}
