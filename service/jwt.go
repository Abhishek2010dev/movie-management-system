package service

import (
	"fmt"
	"time"

	"github.com/Abhishek2010dev/movie-management-system/dto"
	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	secretKey []byte
}

func NewJwt(secretKey string) *Jwt {
	return &Jwt{[]byte(secretKey)}
}

func (j *Jwt) CreateToken(id int64, role models.UserRole) (string, error) {
	claims := dto.Claims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", fmt.Errorf("Failed to create JWT token: %w", err)
	}
	return ss, nil
}

func (j *Jwt) VerifyToken(tokenStr string) (*dto.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &dto.Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token or claims")
}
