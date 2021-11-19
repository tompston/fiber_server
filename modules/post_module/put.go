package post_module

import (
	"context"
	"gofiber_server/settings/database"
	"gofiber_server/sqlc/sqlc"
	res "gofiber_server/utils/response"
	"gofiber_server/utils/validate"

	"github.com/gofiber/fiber/v2"
)

// ------------
// How to check if the update  row does not exist
// ------------
// the sqlc raw query for updating post has a RETURNING value, which means that on a
// succesfull update the specified rows will be returned and stored inside the data
// variable. This means that if we update a row that does not exist, an empty row will
// be returned back. We can check if the returned row is empty by comparing it to the struct
// that should be returned.
// * to get the name of the struct, hover over the data variable

func UpdatePostBody(c *fiber.Ctx) error {

	// define the struct that you want to get from the client
	payload := new(UpdatePostBodyParams)
	if err := c.BodyParser(payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	if err := validate.NewValidator().Struct(payload); err != nil {
		return validate.CheckForValidationError(c, err)
	}

	updated_post := sqlc.UpdatePostBodyParams{PostBody: payload.PostBody, PostID: payload.PostID}

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	data, err := sqlc.New(db).UpdatePostBody(context.Background(), updated_post)

	if (data == sqlc.UpdatePostBodyRow{}) {
		return res.ResponseError(c, nil, "Could not update the Post, because it does not exist!")
	}
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	return res.ResponseSuccess(c, data, res.UpdatedMessage(module_name))
}

func UpdatedPostTitle(c *fiber.Ctx) error {

	// define the struct that you want to get from the client
	payload := new(UpdatePostTitleParams)
	if err := c.BodyParser(payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	if err := validate.NewValidator().Struct(payload); err != nil {
		return validate.CheckForValidationError(c, err)
	}

	updated_post := sqlc.UpdatePostTitleParams{PostTitle: payload.PostTitle, PostID: payload.PostID}

	db, err := database.GetDbConnSql()
	if err != nil {
		return res.ResponseError(c, nil, res.FailedDbConnMessage)
	}
	defer db.Close()

	data, err := sqlc.New(db).UpdatePostTitle(context.Background(), updated_post)

	if (data == sqlc.UpdatePostTitleRow{}) {
		return res.ResponseError(c, nil, "Could not update the Post, because it does not exist!")
	}
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}

	return res.ResponseSuccess(c, data, res.UpdatedMessage(module_name))
}
