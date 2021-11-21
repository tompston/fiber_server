package post_module

import (
	"fmt"
	"gofiber_server/settings"

	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "Post"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/post")

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
