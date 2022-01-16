package post_module

import (
	"context"
	"fiber_server/db/sqlc"
	"fiber_server/settings/database"
	res "fiber_server/utils/response"
	"fiber_server/utils/validate"

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
	if err := validate.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	new_post := sqlc.CreatePostParams{
		PostTitle: payload.PostTitle,
		PostBody:  payload.PostBody,
		UserID:    payload.UserID,
	}

	db, err := database.GetDbConn()
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
