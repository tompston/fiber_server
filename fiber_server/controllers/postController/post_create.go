package postController

import (
	"context"
	"fiber_server/settings/database"
	"fiber_server/utils/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PostsCreate(c *fiber.Ctx) error {

	type PostInformation struct {
		PostBody  string `json:"post_body"`
		PostTitle string `json:"post_title"`
		UserID    int64  `json:"user_id"`
	}

	payload := new(PostInformation)

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	client, err := database.GetDbConnEnt()

	if err != nil {
		log.Fatal(err)
	}

	new_post, err := client.Post.Create().
		SetPostBody(payload.PostBody).
		SetPostTitle(payload.PostTitle).
		// possible problem if a string submitted
		SetUserID(int(payload.UserID)).
		Save(context.Background())

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, new_post, "")
}
