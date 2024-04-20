package dto

type User struct {
	ID string

	Login string
	Keys  []PairKey
}

type PairKey struct {
	PublicKey  string
	PrivateKey string
}

type Challenge struct {
	ID string

	Payload   string
	UserLogin string
	PublicKey string
}

type Access struct {
	userID string
	Number int
}
