package handler

import (
	"errors"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/Abhishek2010dev/movie-management-system/utils"
	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	repository *repository.User
	secretKey  []byte
}

func NewAuth(repository *repository.User, secretKey []byte) *Auth {
	return &Auth{repository, secretKey}
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

	passwordHash, err := utils.HashPassword(payload.Password)
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

	token, err := utils.CreateToken(a.secretKey, id, models.RoleUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(TokenResponse{token})
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email,max=254"`
	Password string `json:"password" validate:"required,min=8"`
}

func (a *Auth) LoginHandler(c fiber.Ctx) error {
	payload := new(LoginPayload)
	if err := c.Bind().JSON(payload); err != nil {
		return err
	}

	user, err := a.repository.FindByEmail(c.Context(), payload.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if !utils.VerifyPassword(payload.Password, user.PasswordHash) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	token, err := utils.CreateToken(a.secretKey, user.Id, user.Role)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(TokenResponse{token})
}

func (a *Auth) RegisterRoutes(r fiber.Router) {
	r.Post("/register", a.RegisterHandler)
	r.Post("/login", a.LoginHandler)
}
