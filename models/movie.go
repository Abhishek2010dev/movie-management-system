package models

import (
	"time"
)

type Genre struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"title" json:"name"`
	Description string `db:"description" json:"description"`
}

type Movie struct {
	ID              int       `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Description     string    `db:"description" json:"description"`
	ReleaseDate     time.Time `db:"release_date" json:"release_date"`
	DurationMinutes int       `db:"duration_minutes" json:"duration_minutes"`
	Director        string    `db:"director" json:"director"`
	PosterPath      string    `db:"poster_path" json:"poster_path"`
	Genres          []Genre   `json:"genres"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
