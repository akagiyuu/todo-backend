package auth

import (
	"context"
	"todo-api/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Pool         *pgxpool.Pool
	TokenService *TokenService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthService) Login(ctx context.Context, req LoginRequest) (string, error) {
	queries := database.New(a.Pool)
	account, err := queries.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	return a.TokenService.Create(account.ID.String())
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthService) Register(ctx context.Context, req RegisterRequest) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	queries := database.New(a.Pool)
	id, err := queries.CreateAccount(ctx, database.CreateAccountParams{
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return "", err
	}

	return a.TokenService.Create(id.String())
}
