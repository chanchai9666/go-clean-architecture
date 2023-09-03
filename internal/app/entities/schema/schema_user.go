package schema

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRequest struct {
	ID        int    `query:"id" json:"id"`
	Username  string `query:"username" json:"username" validate:"required"`
	Email     string `query:"email" json:"email"`
	CreatedAt string `query:"created_at" json:"created_at"`
	UpdatedAt string `query:"updated_at" json:"updated_at"`
}
