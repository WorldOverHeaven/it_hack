package service

import (
	"context"
	"example.com/m/v2/backend/internal/auth"
	"example.com/m/v2/backend/internal/database"
	"example.com/m/v2/backend/internal/dto"
	"example.com/m/v2/backend/internal/models"
	"github.com/hashicorp/go-uuid"
)

type service struct {
	db database.Database
	a  auth.Service
}

func New(db database.Database, a auth.Service) Service {
	return &service{db: db, a: a}
}

type Service interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) (string, error)
}

func (s *service) CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return "", err
	}

	user := dto.User{
		ID:    id,
		Login: req.Login,
		Keys: []dto.PairKey{
			{
				PublicKey:  req.OpenKey,
				PrivateKey: req.PrivateKey,
			},
		},
	}

	token, err := s.a.CreateToken(id)
	if err != nil {
		return "", err
	}

	err = s.db.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return token, nil
}
