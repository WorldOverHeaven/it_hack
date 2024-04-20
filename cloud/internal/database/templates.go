package database

const (
	createUserQuery              = `INSERT Into users (id, login, password, payload) VALUES ($1, $2, $3, $4);`
	selectUserByLoginAndPassword = `SELECT id, payload FROM users WHERE login = $1 AND password = $2;`
	selectUserByID               = `SELECT login, password, payload FROM users WHERE id = $1;`
	selectPayloadByID            = `SELECT payload FROM users WHERE id = $1;`
	putPayload                   = `UPDATE users SET payload = $1 WHERE id = $2;`
)
