package middleware

import (
	"fiber_server/utils/auth"
	"fiber_server/utils/response"

	"github.com/gofiber/fiber/v2"
)

// Middleware that checks if the user is authenticated
// If unauthenticated (meaning if the jwt stored in the cookie is not valid), returns unauthendicated response
func IsAuth(c *fiber.Ctx) error {

	is_authenticated, err := auth.IsAuthenticated(c)

	if err != nil || !is_authenticated {
		return response.ResponseUnauthenticated(c, nil, "User Unauthenticated!")
	}

	return c.Next()
}
