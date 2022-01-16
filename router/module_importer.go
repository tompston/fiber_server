// this file will be updated on each run. Do not edit

package router

import (
	"fiber_server/modules/comment_module"
	"fiber_server/modules/post_module"
	"fiber_server/modules/user_module"

	"github.com/gofiber/fiber/v2"
)

func ProjectModules(app *fiber.App) {

	// add a prefix for the routes
	api := app.Group("/api")

	// pass down the app + the api prefix
	user_module.Routes(app, api)
	post_module.Routes(app, api)
	comment_module.Routes(app, api)

}
