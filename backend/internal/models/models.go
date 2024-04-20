package models

type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type CreateUserRequest struct {
	Login      string `json:"login"`
	OpenKey    string `json:"open_key"`
	PrivateKey string `json:"private_key"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}
