package validate

import (
	"gopkg.in/go-playground/validator.v9"
)

// https://docs.gofiber.io/guide/validation

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
