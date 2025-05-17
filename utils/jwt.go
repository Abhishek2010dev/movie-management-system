package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int             `json:"id"`
	Role   models.UserRole `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(secretKey []byte, userId int, role models.UserRole) (string, error) {
	claims := Claims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create JWT token: %w", err)
	}
	return ss, nil
}

func VerifyToken(secret []byte, tokenStr string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid or expired token")
	}

	return &claims, nil
}
