package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"mephi_hack/cloud/internal/dto"
)

type Database interface {
	Ping(ctx context.Context) error
	CreateUser(ctx context.Context, user dto.User) error
	GetUserByID(ctx context.Context, id string) (dto.User, error)
	GetPayload(ctx context.Context, userID string) (string, error)
	PutPayload(ctx context.Context, userID string, payload string) error
	GetUserByLoginAndPassword(ctx context.Context, login, password string) (dto.User, error)
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

	return &database{client: client}, nil
}

func (db *database) Ping(ctx context.Context) error {
	_, err := db.client.QueryContext(ctx, "SELECT 1 FROM users;")
	if err != nil {
		return err
	}

	return nil
}

func (db *database) CreateUser(ctx context.Context, user dto.User) error {
	_, err := db.client.QueryContext(ctx, createUserQuery, user.ID, user.Login, user.Password, user.Payload)
	if err != nil {
		return err
	}
	fmt.Printf("SAVED USER %+v\n", user)
	return nil
}

func (db *database) GetUserByLoginAndPassword(ctx context.Context, login, password string) (dto.User, error) {
	row := db.client.QueryRowContext(ctx, selectUserByLoginAndPassword, login, password)

	user := dto.User{
		ID:       "",
		Login:    login,
		Password: password,
		Payload:  "",
	}

	err := row.Scan(&user.ID, &user.Payload)
	if err != nil {
		return dto.User{}, errors.Errorf("user not found with login = %s and password = %s", login, password)
	}

	fmt.Printf("GOT user %+v\n", user)
	return user, nil
}

func (db *database) GetUserByID(ctx context.Context, id string) (dto.User, error) {
	row := db.client.QueryRowContext(ctx, selectUserByID, id)

	user := dto.User{
		ID:       id,
		Login:    "",
		Password: "",
		Payload:  "",
	}

	err := row.Scan(&user.Login, &user.Password, &user.Payload)
	if err != nil {
		return dto.User{}, errors.Errorf("user not found with id = %s ", id)
	}

	fmt.Printf("GOT user %+v\n", user)
	return user, nil
}

func (db *database) GetPayload(ctx context.Context, userID string) (string, error) {
	row := db.client.QueryRowContext(ctx, selectPayloadByID, userID)

	var payload string

	err := row.Scan(&payload)
	if err != nil {
		return "", errors.Errorf("user not found with id = %s ", userID)
	}

	return payload, nil
}

func (db *database) PutPayload(ctx context.Context, userID string, payload string) error {
	_, err := db.client.QueryContext(ctx, putPayload, payload, userID)
	if err != nil {
		return err
	}
	fmt.Printf("updated payload %s where userid = %s", payload, userID)
	return nil
}
