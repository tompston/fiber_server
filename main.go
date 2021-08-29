package main

import (
	"fiber_server/router"
	"fiber_server/settings"
	"fiber_server/utils/auth"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/csrf"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// create the server
	app := fiber.New()
	// logs the requests
	app.Use(logger.New())
	// changed in production
	// app.Use(csrf.New())

	// use the routes that are defiend by the Url function
	router.Url(app)
	// oauth2 setup with goth_fiber, need to finish
	auth.OauthConfig(app)

	// import the .env int
	port, _ := strconv.Atoi(settings.Config("GOLANG_PORT"))
	// run the server
	app.Listen(fmt.Sprintf("%s%d", ":", port))
}
