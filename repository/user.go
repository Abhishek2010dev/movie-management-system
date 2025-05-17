package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db}
}

type CreateUserPayload struct {
	Name         string
	Email        string
	PasswordHash string
}

var ErrDuplicateEmail = errors.New("duplicate email")

func (u *User) Create(ctx context.Context, payload CreateUserPayload) (int, error) {
	var id int
	query := `INSERT INTO users (name, email, password_hash)
	          VALUES ($1, $2, $3) RETURNING id`

	err := u.db.QueryRowContext(ctx, query, payload.Name, payload.Email, payload.PasswordHash).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return 0, ErrDuplicateEmail
		}
		return 0, fmt.Errorf("create user: %w", err)
	}
	return id, nil
}

func (u *User) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password_hash, role FROM users WHERE email = $1"
	if err := u.db.GetContext(ctx, &user, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch user data: %w", err)
	}
	return &user, nil
}
