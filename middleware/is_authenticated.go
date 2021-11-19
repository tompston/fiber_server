package middleware

import (
	"gofiber_server/utils/auth"
	"gofiber_server/utils/response"

	"github.com/gofiber/fiber/v2"
)

// Middleware that checks if the user is authenticated
// If unauthenticated (meaning if the jwt stored in the cookie is not valid), returns unauthenticated response
func IsAuth(c *fiber.Ctx) error {

	is_authenticated, err := auth.TokenIsValid(c, auth.AccessCookieName)

	// if the access_token is not valid, return unauthenticated response
	if !is_authenticated {
		return response.ResponseUnauthenticated(c, nil, "User Unauthenticated!")
	}

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	return c.Next()
}
