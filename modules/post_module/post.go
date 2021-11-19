package post_module

import (
	"context"
	"gofiber_server/settings/database"
	"gofiber_server/sqlc/sqlc"
	res "gofiber_server/utils/response"
	"gofiber_server/utils/validate"

	"github.com/gofiber/fiber/v2"
)

// POST body example
// 	{
//     "post_title": "this is a new post",
//     "post_body": "this is the content of the post",
//     "user_id": 1
// 	}

func CreatePost(c *fiber.Ctx) error {

	// define the struct that you want to get from the client
	payload := new(PostParams)
	if err := c.BodyParser(payload); err != nil {
		return res.ResponseError(c, err.Error(), "")
	}
	errors := validate.ValidateStruct(*payload)
	// if the submitted struct doesn't pass validation, return the error response
	if errors != nil {
		return res.ResponseError(c, errors, "")
	}

	// assign the new values to the struct from the sqlc package to execute the function
	// ----
	// if jwt cookie auth is implemented, you will
	// 	  1. add middleware that will check if you're logged in before accessing this route
	//    2. get the userId from the cookie with this jwt_id, err := auth.RequestUserId(c)
	// ----
	new_post := sqlc.CreatePostParams{
		PostTitle: payload.PostTitle,
		PostBody:  payload.PostBody,
		UserID:    payload.UserID,
	}

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	data, err := sqlc.New(db).CreatePost(context.Background(), new_post)
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	return res.ResponseSuccess(c, data, res.CreatedMessage(module_name))
}
