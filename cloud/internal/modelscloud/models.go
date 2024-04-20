package modelscloud

type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type AuthUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	Token string `json:"token"`
}

type GetPayloadRequest struct {
	Token string `json:"token"`
}

type GetPayloadResponse struct {
	Payload string `json:"payload"`
}

type PutPayloadRequest struct {
	Token   string `json:"token"`
	Payload string `json:"payload"`
}

type PutPayloadResponse struct{}
