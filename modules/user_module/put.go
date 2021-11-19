package user_module

import (
	res "gofiber_server/utils/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	return res.ResponseSuccess(c, data, "User updated succesfully!")
}
