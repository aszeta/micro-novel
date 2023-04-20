package account

import (
	"context"
	"errors"

	"github.com/aszeta/micro-novel/account/security"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	ValidateAccount(ctx context.Context, mail, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

type service struct {
	db    *mongo.Client
	redis *redis.Client
	ctx   *context.Context
}

func NewService(ctx *context.Context, db *mongo.Client, redis *redis.Client) *service {
	return &service{
		redis: redis,
		db:    db,
		ctx:   ctx,
	}
}

func (s *service) ValidateAccount(ctx context.Context, email, password string) (string, error) {

	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return "", ErrInvalidUser
	}
	token, err := security.NewToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return "", ErrInvalidToken
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return "", ErrInvalidToken
	}
	return tData["email"].(string), nil
}
