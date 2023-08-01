package models

type SignUpUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	IsAdmin  bool   `json:"isAdmin"`
}

type SignInUserRequest struct {
	JWT      string `json:"JWT,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	JWT      string `json:"jwt,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

type GetUserByIDRequest struct {
	JWT     string `json:"jwt,omitempty"`
	User_ID string `json:"user_ID,omitempty" validate:"required"`
}

type DeleteUserByIDRequest struct {
	User_ID string `json:"user_ID" validate:"required"`
}
