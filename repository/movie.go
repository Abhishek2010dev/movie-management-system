package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/jmoiron/sqlx"
)

type Movie struct {
	db *sqlx.DB
}

func NewMovie(db *sqlx.DB) *Movie {
	return &Movie{db}
}

type CreateMoviePayload struct {
	Title           string
	Description     string
	ReleaseDate     time.Time
	DurationMinutes int
	Director        string
	PosterPath      string
	GenreIDs        []int
}

func (m *Movie) Create(ctx context.Context, payload CreateMoviePayload) (*models.Movie, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        INSERT INTO movie (title, description, release_date, duration_minutes, director, poster_path)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	var movieID int
	err = tx.QueryRowContext(ctx, query, payload.Title, payload.Description, payload.ReleaseDate, payload.DurationMinutes, payload.Director, payload.PosterPath).Scan(&movieID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert movie: %w", err)
	}

	genreQuery := `INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)`
	stmt, err := tx.PrepareContext(ctx, genreQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare genre insert statement: %w", err)
	}
	defer stmt.Close()

	for _, genreID := range payload.GenreIDs {
		_, err := stmt.ExecContext(ctx, movieID, genreID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert movie genre: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return m.FindByID(ctx, movieID)
}

func (m *Movie) FindByID(ctx context.Context, id int) (*models.Movie, error) {
	var movie models.Movie
	selectMovieQuery := `
	SELECT  id, title, description, release_date, duration_minutes, director, poster_path, created_at 
	FROM movie WHERE  id = $1;
	`
	err := m.db.GetContext(ctx, &movie, selectMovieQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to select movie (ID %d): %w", id, err)
	}

	selectGenreQuery := `
		SELECT g.id, g.name, g.description
		FROM movie_genre mg
	        LEFT JOIN genre g ON mg.genre_id = g.id
		WHERE mg.movie_id = $1
	`
	var genres []models.Genre
	err = m.db.SelectContext(ctx, &genres, selectGenreQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to select genre: %w", err)
	}

	movie.Genres = genres
	return &movie, nil
}
