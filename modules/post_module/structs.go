package post_module

// structs copied from the sqlc generated code with added validate fields

type PostParams struct {
	PostTitle string `json:"post_title" validate:"required,min=6,max=50"`
	PostBody  string `json:"post_body" validate:"required,min=6,max=50"`
	UserID    int32  `json:"user_id" validate:"required"`
}

type UpdatePostBodyParams struct {
	PostBody string `json:"post_body" validate:"required,min=6,max=50"`
	PostID   int64  `json:"post_id" validate:"required"`
}

type UpdatePostTitleParams struct {
	PostTitle string `json:"post_title" validate:"required,min=6,max=50"`
	PostID    int64  `json:"post_id" validate:"required"`
}
