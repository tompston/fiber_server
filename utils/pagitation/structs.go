package pagitation

// main struct that holds all of the info needed for the controller
type PagitationParams struct {
	CurrPageNum string
	NextPageNum int
	PrevPageNum int
	Limit       int32
	Offset      int32
	PagitationLinks
}

// struct that will hold the generated links for the pages (shown in the api)
type PagitationLinks struct {
	PrevPage string `json:"prev_page"`
	CurrPage string `json:"curr_page"`
	NextPage string `json:"next_page"`
}
