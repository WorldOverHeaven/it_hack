package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const (
	uuidKey = "uuid"
)

type Service interface {
	CreateToken(uuid string) (string, error)
	AuthUser(tokenString string) (string, error)
}

type service struct {
	secret []byte
}

func New(secret string) Service {
	return &service{secret: []byte(secret)}
}

func (s *service) CreateToken(uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		uuidKey: uuid,
	})

	return token.SignedString(s.secret)
}

func (s *service) AuthUser(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return "", errors.Wrap(err, "get token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid, ok := claims[uuidKey].(string)
		if !ok {
			return "", fmt.Errorf("claim parsing %s failed", uuidKey)
		}
		return uuid, nil
	}

	return "", fmt.Errorf("claim missing")
}
