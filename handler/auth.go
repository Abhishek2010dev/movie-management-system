package handler

import (
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/Abhishek2010dev/movie-management-system/service"
	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	jwtService      *service.Jwt
	passwordService *service.Password
	repository      *repository.User
}

func NewAuth(jwtService *service.Jwt, passwordService *service.Password, repository *repository.User) *Auth {
	return &Auth{
		jwtService,
		passwordService,
		repository,
	}
}

type RegisterPayload struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=254"`
	Password string `json:"password" validate:"required,min=8"`
}

func (a *Auth) RegisterHandler(c fiber.Ctx) error {
	payload := new(RegisterPayload)
	if err := c.Bind().Body(payload); err != nil {
		return err
	}
	return nil
}

func (a *Auth) RegisterRoutes(r fiber.Router) {
	r.Post("/register", a.RegisterHandler)
}
