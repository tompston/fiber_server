package post_module

import (
	"context"
	"gofiber_server/settings/database"
	"gofiber_server/sqlc/sqlc"
	res "gofiber_server/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	pag_param, err := database.GetPagitationParams(c, BASE)
	// if the provided page query param is a string / float, return an error
	if err != nil {
		return res.ResponseError(c, nil, res.PageQueryIsNotIntMessage)
	}
	// assign the limit offset values to a struct that will be passed to the query generated by sqlc
	limit_offset := sqlc.GetPostsParams{pag_param.Limit, pag_param.Offset}

	data, err := sqlc.New(db).GetPosts(context.Background(), limit_offset)
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

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

func GetPost(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	data, err := sqlc.New(db).GetPost(context.Background(), int64(id))
	if err != nil {
		return res.ResponseError(c, nil, res.NotFoundMessage(module_name))
	}

	return res.ResponseSuccess(c, data, res.FoundOneMessage(module_name))
}

func GetPostsFromUser(c *fiber.Ctx) error {

	username := c.Params("username")

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	data, err := sqlc.New(db).GetPostsFromUser(context.Background(), username)

	// if the returned array is empty
	if len(data) == 0 {
		return res.ResponseError(c, nil, "No posts found for the user!")
	}
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	return res.ResponseSuccess(c, data, "Posts for the User Found!")
}
