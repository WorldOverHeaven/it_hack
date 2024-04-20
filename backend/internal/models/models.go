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

type GetChallengeRequest struct {
	Login   string `json:"login"`
	OpenKey string `json:"open_key"`
}

type GetChallengeResponse struct {
	ChallengeID string `json:"challenge_id"`
	Challenge   string `json:"challenge"`
}

type SolveChallengeRequest struct {
	ChallengeID     string `json:"challenge_id"`
	SolvedChallenge string `json:"solved_challenge"`
}

type SolveChallengeResponse struct {
	Token string `json:"token"`
}

type RegisterCloudRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterCloudResponse struct {
	Token string `json:"token"`
}

type AuthCloudRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthCloudResponse struct {
	Token string `json:"token"`
}

type GetContainerRequest struct {
	ContainerID string `json:"container_id"`
}

type GetContainerResponse struct {
	ContainerID string `json:"container_id"`
}

type PutContainerRequest struct {
}
