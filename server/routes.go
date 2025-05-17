package server

import (
	"github.com/Abhishek2010dev/movie-management-system/handler"
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
)

func (s *Server) registerRoutes(app *fiber.App) {
	app.Get("/", RootHandler)

	userRepository := repository.NewUser(s.db)
	authHandler := handler.NewAuth(userRepository, s.cfg.JwtSecret)
	authHandler.RegisterRoutes(app.Group("/api/auth"))
}
