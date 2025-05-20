package models

import "time"

type Reservation struct {
	ID              int       `db:"id" json:"id"`
	UserID          int       `db:"user_id" json:"user_id"`
	ShowtimeID      int       `db:"showtime_id" json:"showtime_id"`
	ReservationTime time.Time `db:"reservation_time" json:"reservation_time"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
