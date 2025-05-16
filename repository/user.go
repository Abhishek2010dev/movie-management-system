package repository

import (
	"context"
	"fmt"

	"github.com/Abhishek2010dev/movie-management-system/dto"
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db}
}

func (u *User) Create(ctx context.Context, payload *dto.CreateUserPayload) (int, error) {
	var id int
	query := `
  	      INSERT INTO users (name, email, password_hash)
  	      VALUES ($1, $2, $3)
  	      RETURNING id
	`
	err := u.db.QueryRowxContext(ctx, query, payload.Name, payload.Email, payload.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Failed to create user: %w", err)
	}
	return id, nil
}
