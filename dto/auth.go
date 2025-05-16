package dto

type CreateUserPayload struct {
	Name     string `json:"name" validate:"required,max=100" db:"name"`
	Email    string `json:"email" validate:"required,email,max=254" db:"email"`
	Password string `json:"password" validate:"required,min=8" db:"password_hash"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email,max=254"`
	Password string `json:"password" validate:"required"`
}
