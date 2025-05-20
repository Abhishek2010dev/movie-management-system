package handler

import (
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
)

type Showtime struct {
	repository *repository.Showtime
}

func NewShowtime(repository *repository.Showtime) *Showtime {
	return &Showtime{repository}
}

func (s *Showtime) Create(c fiber.Ctx) error {
	var payload repository.CreateShowtimePayload
	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}
	showtime, err := s.repository.Create(c.Context(), payload)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(showtime)
}
