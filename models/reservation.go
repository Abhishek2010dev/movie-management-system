package models

import "time"

type Reservation struct {
	ID              int       `db:"id"`
	UserID          int       `db:"user_id"`
	ShowtimeID      int       `db:"showtime_id"`
	ReservationTime time.Time `db:"reservation_time"`
	CreatedAt       time.Time `db:"created_at"`
}
