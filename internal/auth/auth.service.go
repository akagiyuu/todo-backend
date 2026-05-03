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

func (a *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	queries := database.New(a.Pool)
	account, err := queries.GetAccountByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return a.TokenService.Create(account.ID.String())
}

func (a *AuthService) Register(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	queries := database.New(a.Pool)
	id, err := queries.CreateAccount(ctx, database.CreateAccountParams{Email: email, Password: string(hashedPassword)})
	if err != nil {
		return "", err
	}

	return a.TokenService.Create(id.String())
}
