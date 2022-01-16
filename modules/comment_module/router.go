package comment_module

import (
	"fmt"
	"fiber_server/settings"
	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "Comment"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/comment")

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/comment", GetComments) 
	// Get By Id
	api.Get("/comment/:id", GetComment) 
	// Create
	api.Post("/comment", CreateComment) 
	// Update With Id
	api.Put("/comment/:id", UpdateComment)
	// Delete With Id
	api.Delete("/comment/:id", DeleteComment) 
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=6,max=50"`
	ExampleTitle string `json:"example_title" validate:"required,min=6,max=50"`
}
