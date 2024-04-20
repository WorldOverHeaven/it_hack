package service

import (
	"context"
	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"mephi_hack/cloud/internal/database"
	"mephi_hack/cloud/internal/dto"
	"mephi_hack/cloud/internal/modelscloud"
	"mephi_hack/pkg/auth"
)

type service struct {
	db database.Database
	a  auth.Service
}

func New(db database.Database, a auth.Service) Service {
	return &service{db: db, a: a}
}

type Service interface {
	CreateUser(ctx context.Context, req modelscloud.CreateUserRequest) (string, error)
	AuthUser(ctx context.Context, req modelscloud.AuthUserRequest) (string, error)
	GetPayload(ctx context.Context, req modelscloud.GetPayloadRequest) (modelscloud.GetPayloadResponse, error)
	PutPayload(ctx context.Context, req modelscloud.PutPayloadRequest) (modelscloud.PutPayloadResponse, error)
}

func (s *service) CreateUser(ctx context.Context, req modelscloud.CreateUserRequest) (string, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return "", err
	}

	user := dto.User{
		ID:       id,
		Login:    req.Login,
		Password: req.Password,
		Payload:  "",
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

func (s *service) AuthUser(ctx context.Context, req modelscloud.AuthUserRequest) (string, error) {
	user, err := s.db.GetUserByLoginAndPassword(context.Background(), req.Login, req.Password)
	if err != nil {
		return "", err
	}

	token, err := s.a.CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) GetPayload(ctx context.Context, req modelscloud.GetPayloadRequest) (modelscloud.GetPayloadResponse, error) {
	token := req.Token

	userID, err := s.a.AuthUser(token)
	if err != nil {
		return modelscloud.GetPayloadResponse{}, errors.Wrap(err, "auth user failed")
	}

	payload, err := s.db.GetPayload(context.Background(), userID)
	if err != nil {
		return modelscloud.GetPayloadResponse{}, errors.Wrap(err, "get payload failed")
	}

	return modelscloud.GetPayloadResponse{Payload: payload}, nil
}

func (s *service) PutPayload(ctx context.Context, req modelscloud.PutPayloadRequest) (modelscloud.PutPayloadResponse, error) {
	token := req.Token

	userID, err := s.a.AuthUser(token)
	if err != nil {
		return modelscloud.PutPayloadResponse{}, errors.Wrap(err, "auth user failed")
	}

	err = s.db.PutPayload(ctx, userID, req.Payload)
	if err != nil {
		return modelscloud.PutPayloadResponse{}, errors.Wrap(err, "put payload failed")
	}

	return modelscloud.PutPayloadResponse{}, nil
}
