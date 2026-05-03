package http

import (
	"todo-api/internal/auth"

	"github.com/go-fuego/fuego"
)

func (s *Server) LoginHandler(c fuego.ContextWithBody[auth.LoginRequest]) (string, error) {
	req, err := c.Body()
	if err != nil {
		return "", err
	}

	token, err := s.AuthService.Login(c.Context(), req)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "invalid email or password",
		}
	}

	return token, nil
}

func (s *Server) RegisterHandler(c fuego.ContextWithBody[auth.RegisterRequest]) (string, error) {
	req, err := c.Body()
	if err != nil {
		return "", err
	}

	token, err := s.AuthService.Register(c.Context(), req)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "invalid email or password",
		}
	}

	return token, nil
}
