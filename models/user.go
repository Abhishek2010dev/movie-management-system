package models

type Role string

var (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	Id           string `db:"id"`
	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	role         Role   `db:"role"`
}
