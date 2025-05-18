package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *validator.Validate
}

func NewStructValidator() *StructValidator {
	validate := validator.New()
	validate.RegisterValidation("file_valid", FileValidator)
	return &StructValidator{validate}
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
	case "datetime":
		return fmt.Sprintf("Invalid date format. Expected format is %s", fe.Param())
	case "dive":
		return "Invalid item in list"
	case "file_valid":
		return "Invalid file: must be JPEG/PNG/WEBP and ≤ 5MB"
	default:
		return fe.Error()
	}
}

const MaxFileSize = 5 << 20

var allowedMIMETypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

func FileValidator(fl validator.FieldLevel) bool {
	fileHeader, ok := fl.Field().Interface().(*multipart.FileHeader)
	if !ok || fileHeader == nil {
		return false
	}

	if fileHeader.Size > MaxFileSize {
		return false
	}

	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return false
	}

	mimeType := http.DetectContentType(buffer)
	if !allowedMIMETypes[mimeType] {
		return false
	}

	return true
}
