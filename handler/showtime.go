package handler

import (
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
)

var ErrNoShowtimeFound = fiber.NewError(fiber.StatusNotFound, "Showtime not found")

type Showtime struct {
	repository *repository.Showtime
}

func NewShowtime(repository *repository.Showtime) *Showtime {
	return &Showtime{repository}
}

func (s *Showtime) Create(c fiber.Ctx) error {
	var payload repository.ShowtimePayload
	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}
	showtime, err := s.repository.Create(c.Context(), payload)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(showtime)
}

func (s *Showtime) GetById(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	showtime, err := s.repository.FindById(c.Context(), id)
	if err != nil {
		return err
	}
	if showtime == nil {
		return ErrNoShowtimeFound
	}
	return c.Status(fiber.StatusOK).JSON(showtime)
}

func (s *Showtime) GetAll(c fiber.Ctx) error {
	showtimes, err := s.repository.FindAll(c.Context())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(showtimes)
}

func (s *Showtime) DeleteById(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	deletedId, err := s.repository.DeleteById(c.Context(), id)
	if err != nil {
		return err
	}
	if deletedId == 0 {
		return ErrNoShowtimeFound
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (s *Showtime) UpdateById(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	var payload repository.ShowtimePayload
	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}
	showtime, err := s.repository.UpdateById(c.Context(), id, payload)
	if err != nil {
		return err
	}
	if showtime == nil {
		return ErrMovieNotFound
	}
	return c.Status(fiber.StatusOK).JSON(showtime)
}
