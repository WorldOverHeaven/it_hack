package dto

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Payload  string `json:"payload"`
}
