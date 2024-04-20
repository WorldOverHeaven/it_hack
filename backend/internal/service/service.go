package service

import (
	"context"
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
	CreateUser(ctx context.Context, user models.CreateUserRequest) (string, error)
	GetChallenge(ctx context.Context, req models.GetChallengeRequest) (models.GetChallengeResponse, error)
	SolveChallenge(ctx context.Context, req models.SolveChallengeRequest) (string, error)
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

func (s *service) SolveChallenge(ctx context.Context, req models.SolveChallengeRequest) (string, error) {
	_, err := s.db.GetChallengeByID(context.Background(), req.ChallengeID)
	if err != nil {
		return "", err
	}

	return "", nil
}
