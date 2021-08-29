package userController

import (
	"context"
	"fiber_server/ent/user"
	"fiber_server/settings/database"
	"fiber_server/utils/response"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// users - get all
func UsersGetAll(c *fiber.Ctx) error {

	client, err := database.GetDbConnEnt()

	if err != nil {
		log.Fatal(err)
	}

	// return all fields for the Users model query example
	// users, err := client.User.Query().All(context.Background())

	users, err := client.User.Query().Select(user.FieldID).All(context.Background())

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, users, "")
}

// users - get by id
func UsersGetById(c *fiber.Ctx) error {

	// convert the url param to string, if param is not string -> send error
	user_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// return c.SendString("Only Numbers!")
		return response.ResponseError(c, nil, "Only Numbers!")
	}

	client, err := database.GetDbConnEnt()

	if err != nil {
		log.Fatal(err)
	}

	users, err := client.User.Query().Where(user.ID(user_id)).All(context.Background())

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, users, "")
}
