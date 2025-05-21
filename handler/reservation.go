package handler

import (
	"github.com/Abhishek2010dev/movie-management-system/middleware"
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/Abhishek2010dev/movie-management-system/utils"
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

func (r *Reservation) CancelReservation(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	deletedId, err := r.repository.Delete(c.Context(), id)
	if err != nil {
		return err
	}
	if deletedId == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Reservation not found")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (r *Reservation) GetUsetRervations(c fiber.Ctx) error {
	claims := fiber.Locals[utils.Claims](c, middleware.AuthPayloadKey)
	reservations, err := r.repository.GetByUserID(c.Context(), claims.UserId)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(reservations)
}
