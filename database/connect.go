package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	return db
}
