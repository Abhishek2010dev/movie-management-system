package handler

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

var ErrMovieNotFound = fiber.NewError(fiber.StatusNotFound, "Movie not found")

type Movie struct {
	repository *repository.Movie
}

func NewMovie(repository *repository.Movie) *Movie {
	return &Movie{repository}
}

type MoviePayload struct {
	Title           string                `form:"title" validate:"required,min=1,max=254"`
	Description     string                `form:"description" validate:"required"`
	ReleaseDate     string                `form:"release_date" validate:"required,datetime=2006-01-02"`
	DurationMinutes int                   `form:"duration_minutes" validate:"required,min=1"`
	Director        string                `form:"director" validate:"required,min=1,max=100"`
	File            *multipart.FileHeader `form:"file" validate:"required"`
	GenreIDs        []int                 `form:"genre_ids" validate:"required,dive,lte=10"`
}

func (m *Movie) Create(c fiber.Ctx) error {
	payload := new(MoviePayload)
	if err := c.Bind().Form(payload); err != nil {
		return err
	}

	ext := filepath.Ext(payload.File.Filename)
	safeFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	err := c.SaveFile(payload.File, fmt.Sprintf("./uploads/poster/%s", safeFileName))
	if err != nil {
		return err
	}

	ReleaseDate, _ := time.Parse("2006-01-02", payload.ReleaseDate)
	createMoviePayload := repository.CreateMoviePayload{
		Title:           payload.Title,
		Description:     payload.Description,
		ReleaseDate:     ReleaseDate,
		DurationMinutes: payload.DurationMinutes,
		Director:        payload.Director,
		PosterPath:      fmt.Sprintf("/poster/%s", safeFileName),
		GenreIDs:        payload.GenreIDs,
	}

	movie, err := m.repository.Create(c.Context(), createMoviePayload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(movie)
}

func (m *Movie) GetById(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	movie, err := m.repository.FindByID(c.Context(), id)
	if err != nil {
		return err
	}
	if movie == nil {
		return ErrMovieNotFound
	}
	return c.Status(fiber.StatusOK).JSON(movie)
}

func (m *Movie) GetAll(c fiber.Ctx) error {
	limit := fiber.Query(c, "limit", 10)
	offset := fiber.Query(c, "offset", 0)
	genreID := fiber.Query(c, "genreID", 0)

	movies, err := m.repository.FindAll(c.Context(), limit, offset, genreID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(movies)
}

func (m *Movie) DeleteById(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	deletedID, err := m.repository.DeleteByID(c.Context(), id)
	if err != nil {
		return err
	}
	if deletedID == 0 {
		return ErrMovieNotFound
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (m *Movie) UpdateById(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	payload := new(repository.UpdateMoviePayload)
	if err := c.Bind().JSON(payload); err != nil {
		return err
	}
	movie, err := m.repository.UpdateByID(c.Context(), id, *payload)
	if err != nil {
		return err
	}
	if movie == nil {
		return ErrMovieNotFound
	}
	return c.Status(fiber.StatusOK).JSON(movie)
}
