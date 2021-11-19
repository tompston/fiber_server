package post_module

import (
	"fmt"
	"gofiber_server/settings"

	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "Post"
var BASE = fmt.Sprintf(settings.Config("PAGE_URL") + "/api" + "/post")

func Routes(app *fiber.App, api fiber.Router) {

	// Get All
	api.Get("/post", GetAllPosts)

	// Get By Id
	api.Get("/post/:id", GetPost)

	// Create
	api.Post("/post", CreatePost)

	// Update Post Body
	api.Put("/post/body", UpdatePostBody)

	// Update Post Title
	api.Put("/post/title", UpdatedPostTitle)

	// Delete With Id
	api.Delete("/post/:id", DeletePost)

	// Get all posts that belong to :username
	api.Get("/user/:username/posts", GetPostsFromUser)

}

// struct copied from the sqlc generated code with added validate fields
type PostParams struct {
	PostTitle string `json:"post_title" validate:"required,min=6,max=50"`
	PostBody  string `json:"post_body" validate:"required,min=6,max=50"`
	UserID    int32  `json:"user_id" validate:"required"`
}

type UpdatePostBodyParams struct {
	PostBody string `json:"post_body" validate:"required,min=6,max=50"`
	PostID   int64  `json:"post_id" validate:"required"`
}

type UpdatePostTitleParams struct {
	PostTitle string `json:"post_title" validate:"required,min=6,max=50"`
	PostID    int64  `json:"post_id" validate:"required"`
}
