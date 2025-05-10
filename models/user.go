package models

type UserRole string

var (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	Id           int64
	Name         string
	Email        string
	Role         UserRole
	PasswordHash string
}
