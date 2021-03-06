package comment_module

import (
	res "fiber_server/utils/response"

	"fiber_server/settings/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteComment(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}
	_ = id

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	return res.ResponseSuccess(c, data, res.DeletedMessage(module_name))
}
