package server

import (
	"log"
	"os"

	"github.com/Abhishek2010dev/movie-management-system/config"
	"github.com/Abhishek2010dev/movie-management-system/database"
	"github.com/Abhishek2010dev/movie-management-system/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/jmoiron/sqlx"
)

func RootHandler(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Server is running")
}

type Server struct {
	cfg *config.Config
	db  *sqlx.DB
}

func New() *Server {
	cfg := config.Load()
	return &Server{
		cfg: cfg,
		db:  database.Connect(cfg.DatabaseUrl),
	}
}

func (s *Server) Setup() *fiber.App {

	if err := os.MkdirAll("./uploads/poster", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		StructValidator: utils.NewStructValidator(),
		ErrorHandler:    utils.ErrorHandler,
		BodyLimit:       5 * 1024 * 1024,
	})

	app.Use(requestid.New())
	app.Use(recoverer.New())
	app.Use(logger.New())

	s.registerRoutes(app)

	return app
}
