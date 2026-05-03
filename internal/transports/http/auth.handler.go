package http

import "github.com/go-fuego/fuego"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) LoginHandler(c fuego.ContextWithBody[LoginRequest]) (string, error) {
	req, err := c.Body()
	if err != nil {
		return "", err
	}

	token, err := s.AuthService.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "invalid email or password",
		}
	}

	return token, nil
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) RegisterHandler(c fuego.ContextWithBody[RegisterRequest]) (string, error) {
	req, err := c.Body()
	if err != nil {
		return "", err
	}

	token, err := s.AuthService.Register(c.Context(), req.Email, req.Password)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "invalid email or password",
		}
	}

	return token, nil
}
