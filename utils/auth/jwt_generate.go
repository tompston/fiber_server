package auth

import (
	"fmt"
	"time"

	"gofiber_server/settings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// define the name of the cookies, so that there is a reference + no possible grammar mistakes
var AccessCookieName = "access_cookie"
var RefreshCookieName = "refresh_cookie"

// The time defined in the jwt claim needs to have the .UTC().Unix() so that the commented out function works
// ------------------
// if !token.Valid {
// 	return c.JSON(fiber.Map{
// 		"jwt_exp": "The token has expired!",
// 	})
// }
// ------------------

// pass the username + user_id to store in the payload
func GenerateAccessCookie(Username string, UserId int64, c *fiber.Ctx) error {

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

	access_cookie := GenerateCookie(AccessCookieName, at, AccessTokenExpirationTime, c)

	c.Cookie(&access_cookie) // send cookie to the client

	return err
}

func GenerateRefreshCookie(Username string, UserId int64, c *fiber.Ctx) error {

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

	refresh_cookie := GenerateCookie(RefreshCookieName, rt, RefreshTokenExpirationTime, c)

	// uncomment 1st line + comment out 2nd line to disable
	// _ = refresh_cookie
	c.Cookie(&refresh_cookie)

	if err != nil {
		fmt.Println("Error")
	}
	return err
}

func GenerateCookie(cookie_name string, cookie_value string, exp_time time.Time, c *fiber.Ctx) fiber.Cookie {

	return fiber.Cookie{
		Name:     cookie_name,
		Value:    cookie_value,
		HTTPOnly: true,
		Secure:   true,
		Expires:  exp_time,
	}
}
