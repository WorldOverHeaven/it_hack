package dto

type User struct {
	ID string

	Login     string
	PublicKey string
}

type Challenge struct {
	ID string

	Payload   string
	UserLogin string
	PublicKey string
}
