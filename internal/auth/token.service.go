package auth

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = fmt.Errorf("invalid token")
)

type TokenService struct {
	Config TokenConfig
}

func NewTokenService() (*TokenService, error) {
	cfg, err := env.ParseAs[TokenConfig]()
	if err != nil {
		return nil, err
	}

	return &TokenService{Config: cfg}, nil
}

func (t *TokenService) Create(subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Duration(t.Config.ExpiredIn) * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(t.Config.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *TokenService) Parse(raw string) (uuid.UUID, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (any, error) {
		return t.Config.Secret, nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	if !token.Valid {
		return uuid.Nil, ErrInvalidToken
	}

	rawID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
