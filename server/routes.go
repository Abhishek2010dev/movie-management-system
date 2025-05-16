package server

import (
	"github.com/Abhishek2010dev/movie-management-system/handler"
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/Abhishek2010dev/movie-management-system/service"
	"github.com/gofiber/fiber/v3"
)

func (s *Server) registerRoutes(app *fiber.App) {
	app.Get("/", RootHandler)

	jwtService := service.NewJwt(s.cfg.JwtSecret)
	passwordService := service.NewPassword()
	userRepository := repository.NewUser(s.db)
	authHandler := handler.NewAuth(jwtService, passwordService, userRepository)
	authHandler.RegisterRoutes(app.Group("/api/auth"))
}
