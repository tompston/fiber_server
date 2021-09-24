package auth

import (
	"fmt"

	"fiber_server/settings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// pass the username + user_id to store in the payload
func GenerateAccessCookie(Username string, UserId int64, c *fiber.Ctx) error {

	// The time defined in the jwt claim needs to have the .UTC().Unix() so that the commented out function works
	// same applies to refresh cookies

	// if !token.Valid {
	// 	return c.JSON(fiber.Map{
	// 		"jwt_exp": "The token has expired!",
	// 	})
	// }

	AccessTokenExpirationTime := settings.AccessTokenExpirationTime()

	var jwtKey = []byte(settings.Config("JWT_KEY"))

	accessToken := jwt.New(jwt.SigningMethodHS256)
	atClaims := accessToken.Claims.(jwt.MapClaims)
	// define the info that is stored inside the access jwt payload
	atClaims["user_id"] = UserId
	atClaims["username"] = Username
	atClaims["exp"] = AccessTokenExpirationTime.UTC().Unix()
	at, err := accessToken.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error creting accessToken!")
	}

	access_cookie := fiber.Cookie{ // initialize a new cookie
		Name:     "access_cookie", // set the name of the key that will be passed to the value
		Value:    at,              // set the generated jwt as the value of the cookie
		HTTPOnly: true,            // not accesible with javascript from the browser
		Secure:   true,
		Expires:  AccessTokenExpirationTime,
	}

	c.Cookie(&access_cookie) // send cookie to the client

	return err
}

func GenerateRefreshCookies(Username string, UserId int64, c *fiber.Ctx) error {

	RefreshTokenExpirationTime := settings.RefreshTokenExpirationTime()

	var jwtKey = []byte(settings.Config("JWT_KEY"))

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	// define the fields that are stored inside the refresh jwt payload
	rtClaims["user_id"] = UserId
	rtClaims["username"] = Username
	rtClaims["exp"] = RefreshTokenExpirationTime.UTC().Unix()
	rt, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error creting refreshToken!")
	}

	refresh_cookie := fiber.Cookie{
		Name:     "refresh_cookie",
		Value:    rt,
		HTTPOnly: true,
		Secure:   true,
		Expires:  RefreshTokenExpirationTime,
	}

	// uncomment 1st line + comment out 2nd line to disable
	// _ = refresh_cookie
	c.Cookie(&refresh_cookie)

	if err != nil {
		fmt.Println("Error")
	}
	return err
}
