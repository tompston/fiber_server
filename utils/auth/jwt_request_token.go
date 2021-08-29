package auth

import (
	"fiber_server/settings"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Get the cookie from the request. Need to pass the name of the cookie that holds the jwt.
func RequestToken(c *fiber.Ctx, cookie_name string) (tkn *jwt.Token, err error) {

	var jwtKey = []byte(settings.Config("JWT_KEY"))

	token, err := jwt.Parse(c.Cookies(cookie_name), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return token, err
}
