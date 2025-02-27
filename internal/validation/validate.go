package validation

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationErrors(err error) string {

	validationErrors := err.(validator.ValidationErrors)
	errorMessages := ""
	for _, vErr := range validationErrors {
		errorMessages += vErr.Field() + " is invalid: " + vErr.Tag() + "\n"
	}
	return errorMessages
}
