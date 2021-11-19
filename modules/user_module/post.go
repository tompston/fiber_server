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

// POST example for the struct
// {
//     "username": "my-username",
//     "password": "my-password"
// }
func CreateUser(c *fiber.Ctx) error {

	// define the struct that you want to get from the client
	// pass it through 2 functions that validate if it is correct. first one validates
	// if the submitted struct doesn't pass validation, return the error response.
	payload := new(UserParams)
	if err := c.BodyParser(payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	if err := val.NewValidator().Struct(payload); err != nil {
		return val.CheckForValidationError(c, err)
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

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	data, err := sqlc.New(db).CreateUser(context.Background(), new_user)
	// as there is a unique constraint on the username field, if an existing username is
	// submitted, it will cause an error.
	if err != nil {
		return res.ResponseError(c, nil, "Username already taken!")
	}

	return res.ResponseSuccess(c, data, res.CreatedMessage(module_name))
}

func Login(c *fiber.Ctx) error {

	payload := new(UserParams)
	if err := c.BodyParser(payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	if err := val.NewValidator().Struct(payload); err != nil {
		return val.CheckForValidationError(c, err)
	}

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	user, err := sqlc.New(db).LoginUser(context.Background(), payload.Username)
	if (user == sqlc.LoginUserRow{}) {
		return res.ResponseError(c, nil, res.NotFoundMessage(module_name))
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
