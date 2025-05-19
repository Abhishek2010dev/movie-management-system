package server

import (
	"github.com/Abhishek2010dev/movie-management-system/handler"
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/gofiber/fiber/v3"
)

func (s *Server) registerRoutes(app *fiber.App) {
	app.Get("/", RootHandler)

	app.Get("/poster/:filename<regex([a-zA-Z0-9._-]+\\.(?:jpg|jpeg|png|webp))>", func(c fiber.Ctx) error {
		return c.SendFile("./uploads/poster/" + c.Params("filename"))
	})

	userRepository := repository.NewUser(s.db)
	authHandler := handler.NewAuth(userRepository, s.cfg.JwtSecret)
	authHandler.RegisterRoutes(app.Group("/api/auth"))

	movieRepository := repository.NewMovie(s.db)
	movieHandler := handler.NewMovie(movieRepository)
	movieHandler.RegisterRoutes(app.Group("/api"))
}
