package server

import (
	"github.com/Abhishek2010dev/movie-management-system/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func RootHandler(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Server is running")
}

type Server struct {
	config *config.Config
}

func New() *Server {
	return &Server{
		config: config.Load(),
	}
}

func (s *Server) Setup() *fiber.App {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(recoverer.New())
	app.Use(logger.New())

	app.Get("/", RootHandler)

	return app
}
