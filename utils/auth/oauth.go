package auth

import (
	"fiber_server/settings"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	gf "github.com/shareed2k/goth_fiber"
)

func OauthConfig(app *fiber.App) {

	base_url := settings.Config("BASE_URL")
	// help for the big brother
	google_client_callback := settings.Config("GOOGLE_CLIENT_CALLBACK")
	google_client_callback_url := fmt.Sprintf("%s%s", base_url, google_client_callback)

	goth.UseProviders(
		google.New(
			settings.Config("GOOGLE_CLIENT_ID"),
			settings.Config("GOOGLE_CLIENT_SECRET"),
			google_client_callback_url),
	)

	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		ctx.JSON(user)
		return nil
	})

	app.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		ctx.Redirect("/")
		return nil
	})

	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
			ctx.JSON(gothUser)
		} else {
			gf.BeginAuthHandler(ctx)
		}
		return nil
	})
}
