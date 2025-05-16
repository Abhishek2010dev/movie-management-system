package dto

import (
	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int64           `json:"id"`
	Role   models.UserRole `json:"role"`
	jwt.RegisteredClaims
}
