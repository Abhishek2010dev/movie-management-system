package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
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

func (m *Movie) FindAll(ctx context.Context, limit, offset int) ([]models.Movie, error) {
	movieQuery := `
		SELECT id, title, description, release_date, duration_minutes, director, poster_path, created_at
		FROM movie
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	var movies []models.Movie
	err := m.db.SelectContext(ctx, &movies, movieQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movies: %w", err)
	}

	if len(movies) == 0 {
		return movies, nil
	}

	movieIDs := make([]int, len(movies))
	idIndexMap := make(map[int]*models.Movie)
	for i, movie := range movies {
		movieIDs[i] = movie.ID
		idIndexMap[movie.ID] = &movies[i]
	}

	query := `
		SELECT mg.movie_id, g.id, g.name, g.description
		FROM movie_genre mg
		LEFT JOIN genre g ON mg.genre_id = g.id
		WHERE mg.movie_id = ANY($1)
	`

	type genreResult struct {
		MovieID int `db:"movie_id"`
		models.Genre
	}

	var results []genreResult
	err = m.db.SelectContext(ctx, &results, query, pq.Array(movieIDs))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch genres for movies: %w", err)
	}

	for _, res := range results {
		if movie, exists := idIndexMap[res.MovieID]; exists {
			movie.Genres = append(movie.Genres, res.Genre)
		}
	}

	return movies, nil
}

func (m *Movie) DeleteByID(ctx context.Context, id int) (int, error) {
	query := "DELETE FROM movie WHERE id = $1 RETURNING id"
	var databaseId int
	if err := m.db.QueryRowContext(ctx, query, id).Scan(&databaseId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

	}
	return databaseId, nil
}

type UpdateMoviePayload struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ReleaseDate     time.Time `json:"release_date"`
	DurationMinutes int       `json:"duration_minutes"`
	Director        string    `json:"director"`
	GenreIDs        []int     `json:"genre_ids"`
}

func (m *Movie) UpdateByID(ctx context.Context, id int, payload UpdateMoviePayload) (*models.Movie, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	var updatedID int
	updateQuery := `
		UPDATE movie
		SET title = $1,
			description = $2,
			release_date = $3,
			duration_minutes = $4,
			director = $5
		WHERE id = $6
		RETURNING id`

	err = tx.QueryRowContext(ctx, updateQuery,
		payload.Title, payload.Description,
		payload.ReleaseDate, payload.DurationMinutes,
		payload.Director, id).Scan(&updatedID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to update movie: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM movie_genre WHERE movie_id = $1`, updatedID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete old genres: %w", err)
	}

	genreStmt, err := tx.PrepareContext(ctx, `INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare genre insert statement: %w", err)
	}
	defer genreStmt.Close()

	for _, genreID := range payload.GenreIDs {
		_, err := genreStmt.ExecContext(ctx, updatedID, genreID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert movie genre: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return m.FindByID(ctx, updatedID)
}
