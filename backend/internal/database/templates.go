package database

const (
	createUserQuery                 = `INSERT Into users (id, login, public_key, private_key) VALUES ($1, $2, $3, $4);`
	createChallengeQuery            = `INSERT Into challenges (id, payload, public_key, user_login) VALUES ($1, $2, $3, $4);`
	selectChallengeByID             = `SELECT id, payload, public_key, user_login FROM challenges WHERE id = $1;`
	selectUserIDByLoginAndPublicKey = `SELECT id FROM users WHERE login = $1 AND public_key = $2;`
	selectUserLoginByID             = `SELECT login FROM users WHERE id = $1 LIMIT 1;`
)
