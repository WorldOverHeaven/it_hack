package service

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"mephi_hack/backend/internal/database"
	"mephi_hack/backend/internal/dto"
	"mephi_hack/backend/internal/models"
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
		ID:        id,
		Login:     req.Login,
		PublicKey: req.PublicKey,
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

	challengeString := lo.RandomString(32, []rune("1234567890qwertyuiopasdfghjklzxcvbnm"))

	challenge := base64.StdEncoding.EncodeToString([]byte(challengeString))

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

	err = s.verify(challenge.Payload, req.SolvedChallenge, challenge.PublicKey)
	if err != nil {
		return models.SolveChallengeResponse{}, errors.Wrap(err, "verify failed")
	}

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

func (*service) verify(challengeString, solvedChallengeString, publicKeyString string) error {
	publicKeyString = "-----BEGIN PUBLIC KEY-----\n" + publicKeyString + "\n-----END PUBLIC KEY-----"
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyString))
	if publicKeyBlock == nil {
		return errors.Errorf("Failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return errors.Wrap(err, "Failed to parse public key")
	}

	rsaPubKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return errors.Errorf("Could not convert to rsa.PublicKey")
	}

	fmt.Println("Successfully parsed public key")

	challenge, err := base64.StdEncoding.DecodeString(challengeString)
	if err != nil {
		return errors.Wrap(err, "Failed to parse challenge")
	}

	solvedChallenge, err := base64.StdEncoding.DecodeString(solvedChallengeString)
	if err != nil {
		return errors.Wrap(err, "Failed to parse solved challenge")
	}

	hashed := sha256.Sum256(challenge)

	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], solvedChallenge)
	if err != nil {
		return errors.Wrap(err, "Failed to verify")
	}

	fmt.Println("Successfully verified")

	return nil
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
