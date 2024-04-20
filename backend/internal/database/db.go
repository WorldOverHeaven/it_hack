package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"mephi_hack/backend/internal/dto"
)

type Database interface {
	Ping(ctx context.Context) error
	CreateUser(ctx context.Context, user dto.User) error
	CreateChallenge(ctx context.Context, challenge dto.Challenge) error
	GetChallengeByID(ctx context.Context, id string) (dto.Challenge, error)
	GetUserIDByChallenge(ctx context.Context, challenge dto.Challenge) (string, error)
	GetUserLoginByID(ctx context.Context, id string) (string, error)
}

type database struct {
	client *sql.DB
}

func NewDatabase(config Config) (*database, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)
	client, err := sql.Open("postgres", connInfo)

	if err != nil {
		return nil, err
	}

	return &database{
		client: client,
	}, nil
}

func (db *database) Ping(ctx context.Context) error {
	_, err := db.client.QueryContext(ctx, "SELECT 1 FROM users;")
	if err != nil {
		return err
	}

	return nil
}

func (db *database) CreateUser(ctx context.Context, user dto.User) error {
	_, err := db.client.QueryContext(ctx, createUserQuery, user.ID, user.Login, user.PublicKey, user.PrivateKey)
	if err != nil {
		return err
	}
	fmt.Printf("SAVED USER %+v\n", user)
	return nil
}

func (db *database) CreateChallenge(ctx context.Context, challenge dto.Challenge) error {
	_, err := db.client.QueryContext(ctx, createChallengeQuery, challenge.ID, challenge.Payload, challenge.PublicKey, challenge.UserLogin)
	if err != nil {
		return err
	}
	fmt.Printf("CREATED CHALLENGE %+v\n", challenge)
	return nil
}

func (db *database) GetChallengeByID(ctx context.Context, id string) (dto.Challenge, error) {
	row := db.client.QueryRowContext(ctx, selectChallengeByID, id)

	challenge := dto.Challenge{}

	err := row.Scan(&challenge.ID, &challenge.Payload, &challenge.PublicKey, &challenge.UserLogin)
	if err != nil {
		return dto.Challenge{}, errors.Wrap(err, "error getting challenge")
	}

	fmt.Printf("GOT CHALLENGE %+v\n", challenge)

	return challenge, nil
}

func (db *database) GetUserIDByChallenge(ctx context.Context, challenge dto.Challenge) (string, error) {
	row := db.client.QueryRowContext(ctx, selectUserIDByLoginAndPublicKey, challenge.UserLogin, challenge.PublicKey)

	var userID string

	err := row.Scan(&userID)
	if err != nil {
		return "", errors.Wrap(err, "error getting user")
	}

	fmt.Printf("GOT userID %+v\n", userID)

	return userID, nil
}

func (db *database) GetUserLoginByID(ctx context.Context, id string) (string, error) {
	row := db.client.QueryRowContext(ctx, selectUserLoginByID, id)

	var userID string
	err := row.Scan(&userID)
	if err != nil {
		return "", errors.Wrap(err, "error getting user")
	}

	fmt.Printf("GOT userID %+v\n", userID)
	return userID, nil
}
