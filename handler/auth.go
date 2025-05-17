package handler

import (
	"errors"

	"github.com/Abhishek2010dev/movie-management-system/models"
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

type TokenResponse struct {
	Token string `json:"token"`
}

func (a *Auth) RegisterHandler(c fiber.Ctx) error {
	payload := new(RegisterPayload)
	if err := c.Bind().JSON(payload); err != nil {
		return err
	}

	passwordHash, err := a.passwordService.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	createUserPayload := repository.CreateUserPayload{
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: passwordHash,
	}

	id, err := a.repository.Create(c.Context(), createUserPayload)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEmail) {
			return fiber.NewError(fiber.StatusConflict, "User already exists")
		}
		return err
	}

	token, err := a.jwtService.CreateToken(id, models.RoleUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(TokenResponse{token})
}

func (a *Auth) RegisterRoutes(r fiber.Router) {
	r.Post("/register", a.RegisterHandler)
}
