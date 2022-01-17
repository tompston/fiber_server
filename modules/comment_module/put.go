package comment_module

import (
	res "fiber_server/utils/response"
	"fiber_server/utils/validate"

	"fiber_server/settings/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateComment(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}
	_ = id

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	return res.ResponseSuccess(c, data, res.UpdatedMessage(module_name))
}
