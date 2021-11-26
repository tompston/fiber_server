package user_module

import (
	"context"
	"gofiber_server/db/sqlc"
	"gofiber_server/settings/database"
	"gofiber_server/utils/auth"
	res "gofiber_server/utils/response"

	val "gofiber_server/utils/validate"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// {
//     "email": "some-user@mail.com",
//     "username": "my-username",
//     "password": "my-password"
// }
func RegisterUser(c *fiber.Ctx) error {

	payload := new(CreateUserParams)
	if err := val.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return res.ResponseError(c, nil, "Error during hashing the password")
	}

	// pass the sent json object through regex to check if the values are correct
	if val.MatchesRegex(payload.Email, val.RegexEmail) == true {
		return res.ResponseError(c, nil, "The provided email is not valid!")
	}
	if val.MatchesRegex(payload.Username, val.RegexUsername) == true {
		return res.ResponseError(c, nil, "The provided username is not valid!")
	}

	// assign the new values to the struct from the sqlc package to execute the function
	new_user := sqlc.CreateUserParams{
		Email:    payload.Email,
		Username: payload.Username,
		Password: string(hashed_password),
	}

	// create a db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	data, err := sqlc.New(db).CreateUser(context.Background(), new_user)

	// as there is a unique constraint on the username field, if an existing username is
	// submitted, it will cause an error. Refactor this later, maybe edge cases where this is invalid message
	if err != nil {

		if err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"` {
			return res.ResponseError(c, nil, "Username already taken!")
		}
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			return res.ResponseError(c, nil, "Email already used!")
		}

		return res.ResponseError(c, nil, "Error occured during registration!")
	}

	return res.ResponseSuccess(c, data, res.CreatedMessage(module_name))
}

// {
//     "username": "my-username",
//     "password": "my-password"
// }
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

	//
	if err != nil {
		return res.ResponseError(c, nil, "Error occured during user login!")
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
