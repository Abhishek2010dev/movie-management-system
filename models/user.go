package models

type UserRole string

var (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	Id           string   `db:"id"`
	Name         string   `db:"name"`
	Email        string   `db:"email"`
	PasswordHash string   `db:"password_hash"`
	Role         UserRole `db:"role"`
}
