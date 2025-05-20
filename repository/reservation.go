package repository

import (
	"context"
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

func (r *Reservation) Delete(ctx context.Context, reservationID int) error {
	query := `DELETE FROM reservation WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, reservationID)
	if err != nil {
		return fmt.Errorf("failed to delete reservation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no reservation found with id %d", reservationID)
	}

	return nil
}

type ReservationWithLeftSeats struct {
	models.Reservation
	LeftSeats int `db:"left_seats" json:"left_seats"`
}

func (r *Reservation) GetAllWithLeftSeatCount(ctx context.Context) ([]ReservationWithLeftSeats, error) {
	query := `
		SELECT
			res.id,
			res.user_id,
			res.showtime_id,
			res.reservation_time,
			res.created_at,
			st.available_seats - COALESCE(reserved_count.count, 0) AS left_seats
		FROM reservation res
		JOIN showtime st ON res.showtime_id = st.id
		LEFT JOIN (
			SELECT showtime_id, COUNT(*) AS count
			FROM reservation
			GROUP BY showtime_id
		) reserved_count ON reserved_count.showtime_id = res.showtime_id
	`

	var results []ReservationWithLeftSeats
	if err := r.db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("failed to get reservations with left seat count: %w", err)
	}

	return results, nil
}

