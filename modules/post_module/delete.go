package post_module

import (
	"context"
	"gofiber_server/db/sqlc"
	"gofiber_server/settings/database"
	res "gofiber_server/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeletePost(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	// execute the query
	data := sqlc.New(db).DeletePost(context.Background(), int64(id))

	return res.ResponseSuccess(c, data, res.DeletedMessage(module_name))
}
