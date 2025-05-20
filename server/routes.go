package server

import (
	"github.com/Abhishek2010dev/movie-management-system/handler"
	"github.com/Abhishek2010dev/movie-management-system/middleware"
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

	movieRepository := repository.NewMovie(s.db)
	movieHandler := handler.NewMovie(movieRepository)

	showtimeRepository := repository.NewShowtime(s.db)
	showtimeHandler := handler.NewShowtime(showtimeRepository)

	app.Post("/auth/login", authHandler.LoginHandler)
	app.Post("/auth/register", authHandler.RegisterHandler)

	protectedRoutes := app.Group("/api", middleware.AuthMiddleware(s.cfg.JwtSecret))

	protectedRoutes.Get("/movies", movieHandler.GetAll)
	protectedRoutes.Get("/movies/:id<regex((?:0|[1-9][0-9]{0,18}))>", movieHandler.GetById)

	adminRoutes := protectedRoutes.Group("/", middleware.AdminMiddleware)

	adminRoutes.Post("/movies", movieHandler.Create)
	adminRoutes.Delete("/movies/:id<regex((?:0|[1-9][0-9]{0,18}))>", movieHandler.DeleteById)
	adminRoutes.Put("/movies/:id<regex((?:0|[1-9][0-9]{0,18}))>", movieHandler.UpdateById)

	adminRoutes.Post("/showtimes", showtimeHandler.Create)
}
