package user_module

import (
	"context"
	"gofiber_server/db/sqlc"
	"gofiber_server/settings/database"
	res "gofiber_server/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}

	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()

	data := sqlc.New(db).DeleteUser(context.Background(), int64(id))

	return res.ResponseSuccess(c, data, res.DeletedMessage(module_name))
}
