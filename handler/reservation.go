package handler

import (
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
)

type Reservation struct {
	repository *repository.Reservation
}

func NewReservation(repository *repository.Reservation) *Reservation {
	return &Reservation{repository}
}

func (r *Reservation) GetAll(c fiber.Ctx) error {
	reservation, err := r.repository.GetAll(c.Context())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(reservation)
}
