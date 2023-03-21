package domain

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateEntity(entity interface{}) error {
	validate := validator.New()
	err := validate.Struct(entity)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	switch validationError.Tag() {
	case "required":
		return errors.New(validationError.StructField() + " is required")
	case "max":
		return errors.New(validationError.StructField() + " should have a max size of " + validationError.Param())
	case "min":
		return errors.New(validationError.StructField() + " should have a min size of " + validationError.Param())
	case "email|e164":
		return errors.New(validationError.StructField() + " is an invalid e-mail or an invalid phone number")
	}

	return nil
}
