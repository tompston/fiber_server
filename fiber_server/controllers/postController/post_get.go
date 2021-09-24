package postController

import (
	"context"
	"fiber_server/ent/post"
	"fiber_server/ent/user"
	"fiber_server/settings/database"
	"fiber_server/utils/response"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// posts - get all
func PostsGetAll(c *fiber.Ctx) error {

	client, err := database.GetDbConnEnt()
	if err != nil {
		log.Fatal(err)
	}

	posts, err := client.Post.Query().All(context.Background())

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, posts, "")
}

// posts - get by id parameter in the url
func PostsGetById(c *fiber.Ctx) error {

	// convert the url param to string, if param is not string -> send error
	post_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("Only Number!")
	}

	client, err := database.GetDbConnEnt()
	if err != nil {
		log.Fatal(err)
	}

	posts, err := client.Post.Query().Where(post.ID(post_id)).All(context.Background())

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer client.Close()

	return response.ResponseSuccess(c, posts, "")
}

// posts - get all with the given user id
func PostsGetByAuthorId(c *fiber.Ctx) error {

	// convert the url param to string, if param is not string -> send error
	user_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ResponseError(c, nil, "Only Number!")
	}

	client, err := database.GetDbConnEnt()
	if err != nil {
		log.Fatal(err)
	}

	posts, err := client.User.Query().Where(user.ID(user_id)).Select(user.FieldID).WithPosts().All(context.Background())

	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}

	defer client.Close()

	return response.ResponseSuccess(c, posts, "")
}
