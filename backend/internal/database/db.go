package database

import (
	"context"
	"database/sql"
	"example.com/m/v2/backend/internal/dto"
	"fmt"
)

type Database interface {
	Ping(ctx context.Context) error
	CreateUser(ctx context.Context, user dto.User) error
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
	return nil
}

func (db *database) CreateAccess(ctx context.Context, access dto.Access) error {
	db.accesses = append(db.accesses, access)
	return nil
}
