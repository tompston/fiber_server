package validate

import (
	"fmt"
	res "gofiber_server/utils/response"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

// credits to this --> https://dev.to/koddr/how-to-make-clear-pretty-error-messages-from-the-go-backend-to-your-frontend-21b2

// NewValidator func for create a new validator for struct fields.
func NewValidator() *validator.Validate {
	return validator.New()
}

// CheckForValidationError func for checking validation errors in struct fields.
func CheckForValidationError(c *fiber.Ctx, errFunc error) error {
	if errFunc != nil {
		return res.ResponseError(c, ValidatorErrors(errFunc), "Object did not pass validation!")
	}
	return nil
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {

	errFields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		structName := strings.Split(err.Namespace(), ".")[0]
		errFields[err.Field()] = fmt.Sprintf(
			"failed '%s' tag check (value '%s' is not valid for %s struct)",
			err.Tag(), err.Value(), structName,
		)
	}

	return errFields
}
