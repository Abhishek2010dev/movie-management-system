package handler

import (
	"database/sql"
	"net/http"

	"github.com/Abhishek2010dev/movie-management-system/config"
	"github.com/Abhishek2010dev/movie-management-system/repository"
	"github.com/Abhishek2010dev/movie-management-system/service"
	"github.com/go-chi/chi/v5"
)

type Auth struct {
	repository      repository.User
	jwtService      service.Jwt
	passwordService service.Password
}

func NewAuth(db *sql.DB, cfg config.Auth) *Auth {
	return &Auth{
		repository:      repository.NewUser(db),
		passwordService: service.NewPassword(),
		jwtService:      service.NewJwt(cfg.JwtSecret),
	}
}

func (a *Auth) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register User"))
}

func (a *Auth) RegisterRoutes(r chi.Router) {
	r.Get("/register", a.RegisterUser)
}
