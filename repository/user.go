package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Abhishek2010dev/movie-management-system/dto"
	"github.com/Abhishek2010dev/movie-management-system/models"
)

type User interface {
	Create(payload *dto.CreateUserPayload) (int64, error)
	CheckEmail(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return &userRepo{db}
}

func (u *userRepo) Create(payload *dto.CreateUserPayload) (int64, error) {
	query := `
		INSERT INTO users(name, email, password_hash) VALUES ($1, $2, $3)
	  	RETURNING id;
	`
	var id int64
	err := u.db.QueryRow(query, payload.Name, strings.ToLower(payload.Email), payload.Password).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("Failed to create user: %w", err)
		}
	}
	return id, nil
}

func (u *userRepo) CheckEmail(email string) (bool, error) {
	query := "SELECT 1 FROM users WHERE email = $1"
	var exists bool
	if err := u.db.QueryRow(query, email).Scan(&exists); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("Failed to check email existence: %w", err)
	}
	return exists, nil
}

func (u *userRepo) FindByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, email, role, password_hash FROM users WHERE email = $1"
	var user models.User
	rows := u.db.QueryRow(query, email)
	err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to find user by email %s: %w", email, err)
	}
	return &user, nil
}
