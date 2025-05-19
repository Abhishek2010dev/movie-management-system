package models

import "time"

type Showtime struct {
	ID             int       `db:"id" json:"id"`
	MovieID        int       `db:"movie_id" json:"movie_id"`
	StartTime      time.Time `db:"start_time" json:"start_time"`
	EndTime        time.Time `db:"end_time" json:"end_time"`
	AvailableSeats int       `db:"available_seats" json:"available_seats"`
	Price          float64   `db:"price" json:"price"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
