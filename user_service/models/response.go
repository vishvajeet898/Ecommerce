package models

type SignUpUserResponse struct {
	Jwt   string `json:"jwt,omitempty"`
	Error error  `json:"error,omitempty"`
}

type SignInUserResponse struct {
	Jwt   string `json:"jwt,omitempty"`
	Error error  `json:"error,omitempty"`
}

type UpdateUserResponse struct {
	Ok error
}

type GetUserByIDResponse struct {
	User Users `json:"user"`
}

type DeleteUserByIDResponse struct {
	Ok error `json:"ok"`
}
