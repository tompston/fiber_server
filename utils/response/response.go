package response

import (
	"github.com/gofiber/fiber/v2"
)

// idea / code taken from here https://github.com/fachryansyah/fotongo/blob/master/utils/response.go

// ResponseSuccess : returning json structur for success request
func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": message,
		"data":    data,
	})
}

// ResponseNotFound : returning json structur for notfound request
func ResponseNotFound(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(404).JSON(fiber.Map{
		"status":  404,
		"message": message,
	})
}

// ResponseError : returning json structur for error request
func ResponseError(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"status":  500,
		"message": message,
		"data":    data,
	})
}

// ResponseUnauthenticated : returning json structur for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(403).JSON(fiber.Map{
		"status":  403,
		"message": message,
		"data":    data,
	})
}

// ResponseValidationError : returning json structur for validation error request
func ResponseValidationError(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(304).JSON(fiber.Map{
		"status":  304,
		"message": message,
		"data":    data,
	})
}
