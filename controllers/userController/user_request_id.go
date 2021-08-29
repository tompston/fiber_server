package userController

import (
	"fiber_server/utils/auth"
	"fiber_server/utils/response"

	"github.com/gofiber/fiber/v2"
)

func GetUserId(c *fiber.Ctx) error {

	jwt_user_id, err := auth.RequestUserId(c)

	// write a redirect to login later
	if err != nil {
		return response.ResponseUnauthenticated(c, nil, "Error occured during token validation!")
	}

	return c.JSON(fiber.Map{
		"jwt_user_id": jwt_user_id,
	})
}

func ProtectedRoute(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, nil, "You are authenthicated!")
}
