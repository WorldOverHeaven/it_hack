package database

import (
	"context"
	"github.com/pkg/errors"
	"mephi_hack/cloud/internal/dto"
)

type Database interface {
	Ping(ctx context.Context) error
	CreateUser(ctx context.Context, user dto.User) error
	GetUserByID(id string) (dto.User, error)
	GetPayload(ctx context.Context, userID string) (string, error)
	PutPayload(ctx context.Context, userID string, payload string) error
	GetUserByLoginAndPassword(ctx context.Context, login, password string) (dto.User, error)
}

type database struct {
	users map[string]dto.User
}

func NewDatabase(config Config) (*database, error) {
	// users map[uuid]dto.User
	users := make(map[string]dto.User)

	return &database{
		users: users,
	}, nil
	//connInfo := fmt.Sprintf(
	//	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	config.Host, config.Port, config.User, config.Password, config.Database,
	//)
	//client, err := sql.Open("postgres", connInfo)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &database{client: client}, nil
}

func (d *database) Ping(ctx context.Context) error {
	return nil
}

func (d *database) CreateUser(ctx context.Context, user dto.User) error {
	if _, ok := d.users[user.ID]; ok {
		return errors.New("user already exists")
	}
	d.users[user.ID] = user
	return nil
}

func (d *database) GetUserByLoginAndPassword(ctx context.Context, login, password string) (dto.User, error) {
	for _, user := range d.users {
		if user.Login == login && user.Password == password {
			return user, nil
		}
	}
	return dto.User{}, errors.Errorf("user not found with login = %s and password = %s", login, password)
}

func (d *database) GetUserByID(id string) (dto.User, error) {
	user, ok := d.users[id]
	if !ok {
		return dto.User{}, errors.Errorf("user with id = %s not found", id)
	}
	return user, nil
}

func (d *database) GetPayload(ctx context.Context, userID string) (string, error) {
	user, ok := d.users[userID]
	if !ok {
		return "", errors.Errorf("not find user with userid = %s", userID)
	}
	return user.Payload, nil
}

func (d *database) PutPayload(ctx context.Context, userID string, payload string) error {
	user, ok := d.users[userID]
	if !ok {
		return errors.Errorf("not find user with userid = %s", userID)
	}
	d.users[userID] = dto.User{
		ID:       user.ID,
		Login:    user.Login,
		Password: user.Password,
		Payload:  payload,
	}
	return nil
}
