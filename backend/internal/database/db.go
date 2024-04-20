package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"mephi_hack/backend/internal/dto"
)

type Database interface {
	Ping(ctx context.Context) error
	CreateUser(ctx context.Context, user dto.User) error
	CreateChallenge(ctx context.Context, challenge dto.Challenge) error
	GetChallengeByID(ctx context.Context, id string) (dto.Challenge, error)
}

type database struct {
	client     *sql.DB
	users      []dto.User
	challenges []dto.Challenge
	accesses   []dto.Access
}

func NewDatabase(config Config) (*database, error) {
	users := make([]dto.User, 0)
	challenges := make([]dto.Challenge, 0)
	accesses := make([]dto.Access, 0)

	return &database{
		client:     nil,
		users:      users,
		challenges: challenges,
		accesses:   accesses,
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

func (db *database) Ping(ctx context.Context) error {
	return nil
}

func (db *database) CreateUser(ctx context.Context, user dto.User) error {
	db.users = append(db.users, user)
	fmt.Printf("SAVED USER %+v\n", user)
	return nil
}

func (db *database) CreateChallenge(ctx context.Context, challenge dto.Challenge) error {
	db.challenges = append(db.challenges, challenge)
	fmt.Printf("CREATED CHALLENGE %+v\n", challenge)
	return nil
}

func (db *database) GetChallengeByID(ctx context.Context, id string) (dto.Challenge, error) {
	items := lo.Filter(db.challenges, func(item dto.Challenge, index int) bool {
		return item.ID == id
	})
	if len(items) != 1 {
		return dto.Challenge{}, errors.Errorf("can't find challenge by id %s, len(items) = %d", id, len(items))
	}

	fmt.Printf("GOT CHALLENGE %+v\n", items[0])

	return items[0], nil
}
