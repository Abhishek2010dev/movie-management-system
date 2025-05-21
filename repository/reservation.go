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

type Reservation struct {
	db *sqlx.DB
}

func NewReservation(db *sqlx.DB) *Reservation {
	return &Reservation{db: db}
}

type ReservationPayload struct {
	ShowtimeID      int       `validate:"required,gt=0" json:"showtime_id"`
	ReservationTime time.Time `validate:"required" json:"reservation_time"`
}

func (r *Reservation) Create(ctx context.Context, payload ReservationPayload, userID int) (*models.Reservation, error) {
	query := `
		INSERT INTO reservation (user_id, showtime_id, reservation_time)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, showtime_id, reservation_time, created_at
	`

	var res models.Reservation
	if err := r.db.GetContext(ctx, &res, query, userID, payload.ShowtimeID, payload.ReservationTime); err != nil {
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}
	return &res, nil
}

func (r *Reservation) GetByUserID(ctx context.Context, userID int) ([]models.Reservation, error) {
	query := `
		SELECT id, user_id, showtime_id, reservation_time, created_at
		FROM reservation
		WHERE user_id = $1
	`

	var reservations []models.Reservation
	if err := r.db.SelectContext(ctx, &reservations, query, userID); err != nil {
		return nil, fmt.Errorf("failed to get reservations by user id: %w", err)
	}
	return reservations, nil
}

func (r *Reservation) GetAll(ctx context.Context) ([]models.Reservation, error) {
	query := `
		SELECT id, user_id, showtime_id, reservation_time, created_at
		FROM reservation
	`

	var reservations []models.Reservation
	if err := r.db.SelectContext(ctx, &reservations, query); err != nil {
		return nil, fmt.Errorf("failed to get all reservations: %w", err)
	}
	return reservations, nil
}

func (r *Reservation) Delete(ctx context.Context, reservationID int) (int, error) {
	query := `DELETE FROM reservation WHERE id = $1 RETURNING id`

	var deletedID int
	err := r.db.QueryRowContext(ctx, query, reservationID).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to delete reservation: %w", err)
	}

	return deletedID, nil
}
