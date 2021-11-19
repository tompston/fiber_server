package user_module

import (
	"context"
	"gofiber_server/settings/database"
	"gofiber_server/sqlc/sqlc"
	"gofiber_server/utils/auth"
	res "gofiber_server/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	pag_param, err := database.GetPagitationParams(c, BASE)
	if err != nil {
		return res.ResponseError(c, nil, res.PageQueryIsNotIntMessage)
	}
	// assign the limit offset values to a struct that will be passed to the query generated by sqlc
	limit_offset := sqlc.GetUsersParams{Limit: pag_param.Limit, Offset: pag_param.Offset}

	data, err := sqlc.New(db).GetUsers(context.Background(), limit_offset)
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	// conver this into a function?
	var message string
	message = res.FoundManyMessage(module_name)
	if len(data) == 0 {
		message = res.NotFoundMessage(module_name)
	}
	// if the returned array has less values than the limit specified in the .env file, then
	// that means that there are no more values that can be returned from the query. Thus, no next pages
	if len(data) < int(pag_param.Limit) {
		pag_param.PagitationLinks.NextPage = "null"
	}

	return res.ResponseSuccessWithPagitation(c, data, pag_param.PagitationLinks, message)
}

func GetUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	data, err := sqlc.New(db).GetUser(context.Background(), int64(id))
	if err != nil {
		return res.ResponseError(c, nil, res.NotFoundMessage(module_name))
	}

	return res.ResponseSuccess(c, data, res.FoundOneMessage(module_name))
}

// GetUserCookie
func GetUserCookie(c *fiber.Ctx) error {

	data, err := auth.RequestToken(c, "access_cookie")
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	return res.ResponseSuccess(c, data.Claims, "User JWT Cookie Information")
}
