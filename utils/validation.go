package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *validator.Validate
}

func NewStructValidator() *StructValidator {
	return &StructValidator{
		validator.New(),
	}
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidationErrors(err error) []ValidationErrorDetail {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		details := make([]ValidationErrorDetail, len(ve))
		for i, fe := range ve {
			details[i] = ValidationErrorDetail{
				Field:   fe.Field(),
				Message: validationErrorMessage(fe),
			}
		}
		return details
	}
	return nil
}

func validationErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Too short"
	case "max":
		return "Too long"
	default:
		return fe.Error()
	}
}
