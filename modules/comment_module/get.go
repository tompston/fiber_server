package comment_module

import (
	"fiber_server/utils/response"

	"github.com/gofiber/fiber/v2"
	"strconv"
	"fiber_server/settings/database"
)

func GetComments(c *fiber.Ctx) error {
	
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	return response.ResponseSuccess(c, data, response.FoundManyMessage(module_name))
}

func GetComment(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ResponseError(c, nil, response.ParamIsNotIntMessage)
	}
	_ = id
	
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	return response.ResponseSuccess(c, data, response.FoundOneMessage(module_name))
}