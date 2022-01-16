package user_module

import (
	"fiber_server/settings"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "User"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/user")

func Routes(app *fiber.App, api fiber.Router) {

	// route to get the jwt cookie
	api.Get("/user/cookie", GetUserCookie)
	// Get All
	api.Get("/user", GetAllUsers)
	// Get By Id
	api.Get("/user/:id", GetUser)
	// Register
	api.Post("/user/register", RegisterUser)
	// Login
	api.Post("/user/login", LoginUser)
	// Update With Id
	api.Put("/user/:id", UpdateUser)
	// Delete With Id
	api.Delete("/user/:id", DeleteUser)
}
