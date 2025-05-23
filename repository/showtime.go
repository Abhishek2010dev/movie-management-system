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

type ShowtimePayload struct {
	MovieID        int       `json:"movie_id" validate:"required,gt=0"`
	StartTime      time.Time `json:"start_time" validate:"required"`
	EndTime        time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	AvailableSeats int       `json:"available_seats" validate:"required,gte=0"`
	Price          float64   `json:"price" validate:"required,gte=0"`
}

func (s *Showtime) Create(ctx context.Context, payload ShowtimePayload) (*models.Showtime, error) {
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

func (s *Showtime) FindAll(ctx context.Context) ([]models.Showtime, error) {
	query := "SELECT id, movie_id, start_time, end_time, available_seats, price, created_at FROM showtime"
	var showtimes []models.Showtime
	if err := s.db.SelectContext(ctx, &showtimes, query); err != nil {
		return showtimes, fmt.Errorf("failed to fetch all showtime: %w", err)
	}
	return showtimes, nil
}

func (s *Showtime) DeleteById(ctx context.Context, id int) (int, error) {
	query := "DELETE FROM showtime WHERE id = $1 RETURNING id"
	var databaseId int
	if err := s.db.GetContext(ctx, &databaseId, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to delete showtime (ID: %v): %w", id, err)
	}
	return databaseId, nil
}

func (s *Showtime) UpdateById(ctx context.Context, id int, payload ShowtimePayload) (*models.Showtime, error) {
	query := `
UPDATE showtime
SET movie_id = $1,
    start_time = $2,
    end_time = $3,
    available_seats = $4,
    price = $5
WHERE id = $6
RETURNING id, movie_id, start_time, end_time, available_seats, price, created_at
`
	var showtime models.Showtime
	err := s.db.GetContext(ctx, &showtime, query, payload.MovieID, payload.StartTime, payload.EndTime, payload.AvailableSeats, payload.Price, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to update showtime (ID: %v): %w", id, err)
	}
	return &showtime, nil
}
