package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()

		// Custom message
		var message string
		switch err.Tag() {
		case "required":
			message = field + " is required"
		case "email":
			message = "Invalid email format"
		default:
			message = field + " is not valid"
		}

		errors[field] = message
	}

	return errors
}
