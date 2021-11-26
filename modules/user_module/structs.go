package user_module

// structs copied from the sqlc generated code with added validate fields
type CreateUserParams struct {
	Email    string `json:"email" validate:"required,min=6,max=50"`
	Username string `json:"username" validate:"required,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserParams struct {
	Username string `json:"username" validate:"required,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
