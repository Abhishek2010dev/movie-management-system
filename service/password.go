package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}
	return string(hashPassword), nil
}

func (p *Password) VerifyPassword(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
