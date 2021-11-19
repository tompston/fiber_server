package user_module

import (
	"fmt"
	"gofiber_server/settings"

	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "User"
var BASE = fmt.Sprintf(settings.Config("PAGE_URL") + "/api" + "/user")

func Routes(app *fiber.App, api fiber.Router) {

	// route to get the jwt cookie
	api.Get("/user/cookie", GetUserCookie)

	// Get All
	api.Get("/user", GetAllUsers)

	// Get By Id
	api.Get("/user/:id", GetUser)

	// Register
	api.Post("/user/register", CreateUser)

	// Login
	api.Post("/user/login", Login)

	// Update With Id
	api.Put("/user/:id", UpdateUser)

	// Delete With Id
	api.Delete("/user/:id", DeleteUser)
}

// struct copied from the sqlc generated code with added validate fields
type UserParams struct {
	Username string `json:"username" validate:"required,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
