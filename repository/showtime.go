package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/jmoiron/sqlx"
)

type Showtime struct {
	db *sqlx.DB
}

func NewShowtime(db *sqlx.DB) *Showtime {
	return &Showtime{db}
}

type CreateShowtimePayload struct {
	MovieID        int       `json:"movie_id" validate:"required,gt=0"`
	StartTime      time.Time `json:"start_time" validate:"required"`
	EndTime        time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	AvailableSeats int       `json:"available_seats" validate:"required,gte=0"`
	Price          float64   `json:"price" validate:"required,gte=0"`
}

func (s *Showtime) Create(ctx context.Context, payload CreateShowtimePayload) (*models.Showtime, error) {
	query := `
	INSERT INTO showtime (movie_id, start_time, end_time, available_seats, price)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, movie_id, start_time, end_time, available_seats, price, created_at
	`
	var showtime models.Showtime
	err := s.db.GetContext(ctx, &showtime, query, payload.MovieID, payload.StartTime, payload.EndTime, payload.AvailableSeats, payload.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to create showtime: %w", err)
	}
	return &showtime, nil
}

func (s *Showtime) FindById(ctx context.Context, id int) (*models.Showtime, error) {
	query := `
		SELECT id, movie_id, start_time, end_time, available_seats, price, created_at
		FROM showtime WHERE id = $1
	`
	var showtime models.Showtime
	if err := s.db.GetContext(ctx, &showtime, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to select showtime (ID: %v): %w", id, err)
	}
	return &showtime, nil
}
