package models

type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type CreateUserRequest struct {
	Login     string `json:"login"`
	PublicKey string `json:"public_key"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type GetChallengeRequest struct {
	Login     string `json:"login"`
	PublicKey string `json:"public_key"`
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

type VerifyRequest struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	Message string `json:"message"`
}
