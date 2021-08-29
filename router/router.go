package router

import (
	"fiber_server/controllers"
	"fiber_server/controllers/postController"
	"fiber_server/controllers/userController"
	"fiber_server/middleware"
	"fiber_server/settings/database"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func Url(app *fiber.App) {
	// home_view
	app.Get("/", Home)
	app.Get("/debug", Debug)
	app.Get("/migrate", controllers.EntMigrate)

	// add a prefix to the routes that start with the name of the variable.
	// Can extend this and chain as many prefixes as you need
	// example can be seen here https://docs.gofiber.io/guide/grouping
	api := app.Group("/api")

	// users
	api.Get("/users", userController.UsersGetAll)
	api.Post("/users/register", userController.UsersRegister)
	api.Post("/users/login", userController.UsersLogin)
	api.Get("/users/request-id", userController.GetUserId) // login -> copy returned access_cookie -> set as a header and make a request to this route
	api.Get("/users/protected", middleware.IsAuth, userController.ProtectedRoute)
	api.Get("/users/:id", userController.UsersGetById) // if the url has a parameter, leave it as the last route

	// posts
	api.Get("/posts", postController.PostsGetAll)
	api.Post("/posts/create", postController.PostsCreate)
	api.Get("/users/:id/posts", postController.PostsGetByAuthorId)
	api.Get("/posts/:id", postController.PostsGetById)

}

// a random route that can be used to test things
func Debug(c *fiber.Ctx) error {

	db, err := database.GetDbConnSqlx()
	if err != nil {
		panic(err)
	}
	db.Exec("")
	db.Close()

	return c.SendString("hello there")
}
