package handler

import (
	"mime/multipart"

	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
)

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
	File            *multipart.FileHeader `form:"file" validate:"required,file_valid"`
	GenreIDs        []int                 `form:"genre_ids" validate:"required,dive,lte=10"`
}

func (m *Movie) Create(c fiber.Ctx) error {
	payload := new(MoviePayload)
	if err := c.Bind().Form(payload); err != nil {
		return err
	}
	//
	// err := c.SaveFile(payload.File, fmt.Sprintf("%s/%s", server.UploadDir, payload.File.Filename))
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (m *Movie) RegisterRoutes(r fiber.Router) {
	r.Post("/movies", m.Create)
}
