package userController

import (
	"context"
	"fiber_server/ent/user"
	"fiber_server/settings/database"
	"fiber_server/utils/auth"
	"fiber_server/utils/response"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// login info that is sent to the server
type Credentials struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

// info that is returned from the db and stored in the jwt
// IF the username and password is correct and exist in the db.
type UserInfo struct {
	UserID   int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
}

func UsersLogin(c *fiber.Ctx) error {

	// Payload holds all of the data that is sent.
	payload := new(Credentials)

	// checks if the passed values in the json match the types defined in the struct.
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	// check if the required fields are not empty.
	if payload.Username == "" {
		return response.ResponseError(c, nil, "No username given!")
	}
	if payload.Password == "" {
		return response.ResponseError(c, nil, "No password given!")
	}

	// proceed with finding user
	client, err := database.GetDbConnEnt()

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	user, err := client.User.Query().Where(user.Username(payload.Username)).Only(context.Background())
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	// check if the sent password matches the hashed password that is stored in the db
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return response.ResponseUnauthenticated(c, nil, "Incorrect password!")
	}

	// if the passwords match, -> generate the access and refresh jwt cookies
	auth.GenerateAccessCookie(user.Username, int64(user.ID), c)
	auth.GenerateRefreshCookies(user.Username, int64(user.ID), c)

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, user, "")
}
