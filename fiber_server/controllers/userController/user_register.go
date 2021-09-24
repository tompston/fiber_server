package userController

import (
	"context"
	"fiber_server/settings/database"
	"fiber_server/utils/response"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UsersRegister(c *fiber.Ctx) error {

	type UsersRegisterParams struct {
		Username string `db:"username" json:"username"`
		Password string `db:"password" json:"password"`
	}

	// Payload holds all of the data that is sent.
	payload := new(UsersRegisterParams)

	// checks if the passed values in the json match the types defined in the struct.
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	if payload.Username == "" {
		return c.SendString("No username given!")
	}
	if payload.Password == "" {
		return c.SendString("No password given!")
	}

	// get the password that is sent in the payload and encrypt it.
	passw, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	client, err := database.GetDbConnEnt()

	if err != nil {
		log.Fatal(err)
	}

	user, err := client.User.Create().
		SetUsername(payload.Username).
		SetPassword(string(passw)).
		Save(context.Background())

	// if error occurs, send it back as a string
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, user, "new user is created!")
}

// example of an ent create controller that does not hash the password
// func UsersCreate(c *fiber.Ctx) error {
// 	client, err := database.GetDbConnEnt()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	user, err := client.User.Create().
// 		SetUsername("my_username").
// 		SetPassword("my_password").
// 		Save(context.Background())
// 	if err != nil {
// 		return response.ResponseError(c, nil, err.Error())
// 	}
// 	defer client.Close()
// 	return response.ResponseSuccess(c, user, "User Created!")
// }
