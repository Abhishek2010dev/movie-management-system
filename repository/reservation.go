package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Reservation struct {
	db *sqlx.DB
}

func NewReservation(db *sqlx.DB) *Reservation {
	return &Reservation{db}
}

type ReservationPayload struct {
	ShowtimeID      int       `validate:"required,gt=0"`
	ReservationTime time.Time `validate:"required"`
}

func (r *Reservation) Create(c context.Context)
