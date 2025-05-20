package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *validator.Validate
}

func NewStructValidator() *StructValidator {
	return &StructValidator{validator.New()}
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
		return fmt.Sprintf("Minimum value/length is %s", fe.Param())
	case "max":
		return fmt.Sprintf("Maximum value/length is %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "gt":
		return fmt.Sprintf("Must be greater than %s", fe.Param())
	case "gtfield":
		return fmt.Sprintf("Must be greater than %s", fe.Param())
	case "datetime":
		return fmt.Sprintf("Invalid date format. Expected format is %s", fe.Param())
	case "dive":
		return "Invalid item in list"
	default:
		return fe.Error()
	}
}
