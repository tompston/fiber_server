package user_module

import (
	"context"
	"gofiber_server/settings/database"
	"gofiber_server/sqlc/sqlc"
	"gofiber_server/utils/auth"
	res "gofiber_server/utils/response"
	val "gofiber_server/utils/validate"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// {
// 	"password": "my-username",
// 	"username": "my-password"
//  }
func RegisterUser(c *fiber.Ctx) error {

	payload := new(UserParams)
	if err := val.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return res.ResponseError(c, nil, "Error during hashing the password")
	}

	// assign the new values to the struct from the sqlc package to execute the function
	new_user := sqlc.CreateUserParams{
		Username: payload.Username,
		Password: string(hashed_password),
	}

	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	data, err := sqlc.New(db).CreateUser(context.Background(), new_user)
	// as there is a unique constraint on the username field, if an existing username is
	// submitted, it will cause an error. Refactor this later, maybe edge cases where this is invalid message
	if err != nil {
		return res.ResponseError(c, nil, "Username already taken!")
	}

	return res.ResponseSuccess(c, data, res.CreatedMessage(module_name))
}

func LoginUser(c *fiber.Ctx) error {

	payload := new(UserParams)
	if err := val.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	user, err := sqlc.New(db).LoginUser(context.Background(), payload.Username)
	if (user == sqlc.LoginUserRow{}) {
		return res.ResponseError(c, nil, res.NotFoundOneMessage(module_name))
	}
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	// check if the sent password matches the hashed password that is stored in the db
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return res.ResponseUnauthenticated(c, nil, "Incorrect password!")
	}

	// if the username exists and the passwords match, set JWT Auth cookies with the user details.
	auth.GenerateAccessCookie(user.Username, int64(user.UserID), c)
	auth.GenerateRefreshCookie(user.Username, int64(user.UserID), c)

	return res.ResponseSuccess(c, data, "Login succesful!")
}
