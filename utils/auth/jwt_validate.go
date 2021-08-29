package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func RequestUserId(c *fiber.Ctx) (int64, error) {

	access_token, err := RequestToken(c, "access_cookie")
	if err != nil {
		return 0, err
	}
	claims := access_token.Claims.(jwt.MapClaims)
	jwt_user_id := int64(claims["user_id"].(float64))

	return jwt_user_id, err
}

// if access token is valid, return true
func IsAuthenticated(c *fiber.Ctx) (bool, error) {

	access_token, err := RequestToken(c, "access_cookie")

	if err != nil {
		return false, err
	}

	if !access_token.Valid {
		return false, err
	}

	return true, err
}
