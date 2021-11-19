package database

import (
	"fmt"
	"gofiber_server/settings"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PagitationParams struct {
	CurrPageNum string
	NextPageNum int
	PrevPageNum int
	Limit       int32
	Offset      int32
	PagitationLinks
}

// struct that will hold the links for the pages
type PagitationLinks struct {
	PrevPage string `json:"prev_page"`
	CurrPage string `json:"curr_page"`
	NextPage string `json:"next_page"`
}

// define the Url snippent that will hold the query
var PageQuery = "/?page="

// ("1") ---> "/?page=1"
func PageQueryString(x string) string {
	return fmt.Sprintf(PageQuery + x)
}

func GetPagitationParams(c *fiber.Ctx, BASE string) (PagitationParams, error) {
	// get the value of PageQuery url param
	page := c.Query("page")

	// if there is an API request that does not have the query, like "http://localhost:5000/api/post/" ,
	// assign it to be a string-int which can be converted into an int, so that there is no error
	if page == "" {
		page = "0"
	}

	page_num, err := strconv.Atoi(page)
	// if the provided query number is negative, we avoid sql error by assigning / correcting it to be 0
	if page_num < 0 {
		page_num = 0
	}
	// if the provided query number is either a string or a float, return error.
	if err != nil {
		return PagitationParams{}, err
	}

	// calculate the offset from the limit.
	// If LIMIT OFFSET is used in the sql query, sqlc will generate a struct with int32 types, so we convert them
	limit, _ := strconv.Atoi(settings.Config("PAGITIATION_LIMIT"))
	limit_32 := int32(limit)
	page_num_32 := int32(page_num)
	offset_32 := limit_32 * page_num_32

	// calculate the next page + prev page int
	prev_page_num := page_num - 1
	next_page_num := page_num + 1

	// generate the links
	pagitation := GeneratePagitationLinks(BASE, prev_page_num, page_num, next_page_num)

	// return the struct that will have all of the needed information for the controller
	pag_params := PagitationParams{
		CurrPageNum:     page,
		NextPageNum:     next_page_num,
		PrevPageNum:     prev_page_num,
		Limit:           limit_32,
		Offset:          offset_32,
		PagitationLinks: pagitation,
	}

	return pag_params, err
}

// return 3 strings that are pagitation links
func GeneratePagitationLinks(BASE string, prev int, curr int, next int) PagitationLinks {
	prev_page_path := BASE + PageQueryString(fmt.Sprint(prev))
	curr_page_path := BASE + PageQueryString(fmt.Sprint(curr))
	next_page_path := BASE + PageQueryString(fmt.Sprint(next))

	if curr == 0 {
		prev_page_path = "null"
	}

	return PagitationLinks{
		CurrPage: curr_page_path,
		NextPage: next_page_path,
		PrevPage: prev_page_path,
	}
}
