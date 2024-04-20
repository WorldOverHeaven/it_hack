package service

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"mephi_hack/backend/internal/database"
	"mephi_hack/backend/internal/dto"
	"mephi_hack/backend/internal/models"
	"mephi_hack/pkg/auth"

	"github.com/samber/lo"
)

type service struct {
	db database.Database
	a  auth.Service
}

func New(db database.Database, a auth.Service) Service {
	return &service{db: db, a: a}
}

type Service interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) (models.CreateUserResponse, error)
	GetChallenge(ctx context.Context, req models.GetChallengeRequest) (models.GetChallengeResponse, error)
	SolveChallenge(ctx context.Context, req models.SolveChallengeRequest) (models.SolveChallengeResponse, error)
	Verify(ctx context.Context, req models.VerifyRequest) (models.VerifyResponse, error)
}

func (s *service) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.CreateUserResponse, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	user := dto.User{
		ID:         id,
		Login:      req.Login,
		PublicKey:  req.PublicKey,
		PrivateKey: req.PrivateKey,
	}

	token, err := s.a.CreateToken(id)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	err = s.db.CreateUser(ctx, user)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	return models.CreateUserResponse{Token: token}, nil
}

func (s *service) GetChallenge(ctx context.Context, req models.GetChallengeRequest) (models.GetChallengeResponse, error) {
	challengeID, err := uuid.GenerateUUID()
	if err != nil {
		return models.GetChallengeResponse{}, errors.Wrap(err, "error generating challengeID")
	}

	challenge := lo.RandomString(32, []rune("1234567890qwertyuiopasdfghjklzxcvbnm"))

	err = s.db.CreateChallenge(ctx, dto.Challenge{
		ID:        challengeID,
		Payload:   challenge,
		UserLogin: req.Login,
		PublicKey: req.OpenKey,
	})
	if err != nil {
		return models.GetChallengeResponse{}, errors.Wrap(err, "create challenge")
	}

	return models.GetChallengeResponse{
		Challenge:   challenge,
		ChallengeID: challengeID,
	}, nil
}

func (s *service) SolveChallenge(ctx context.Context, req models.SolveChallengeRequest) (models.SolveChallengeResponse, error) {
	challenge, err := s.db.GetChallengeByID(context.Background(), req.ChallengeID)
	if err != nil {
		return models.SolveChallengeResponse{}, errors.Wrap(err, "get challenge")
	}

	// TODO проверить challenge

	userID, err := s.db.GetUserIDByChallenge(ctx, challenge)
	if err != nil {
		return models.SolveChallengeResponse{}, errors.Wrap(err, "get user id by challenge")
	}

	token, err := s.a.CreateToken(userID)
	if err != nil {
		return models.SolveChallengeResponse{}, err
	}

	return models.SolveChallengeResponse{Token: token}, nil
}

func (s *service) Verify(ctx context.Context, req models.VerifyRequest) (models.VerifyResponse, error) {
	userID, err := s.a.AuthUser(req.Token)
	if err != nil {
		return models.VerifyResponse{}, errors.Wrap(err, "auth token failed")
	}

	login, err := s.db.GetUserLoginByID(ctx, userID)
	if err != nil {
		return models.VerifyResponse{}, errors.Wrap(err, "get user login failed")
	}

	return models.VerifyResponse{Message: fmt.Sprintf("Ваш логин: %s", login)}, nil
}
