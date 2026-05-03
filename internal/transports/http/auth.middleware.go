package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-fuego/fuego"
)

const (
	authorization string = "Authorization"
	bearer        string = "Bearer "
	AuthKey       string = "id"
)

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorization)
		if header == "" {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Detail: "Missing authorization header",
			})
			return
		}

		raw, isBearer := strings.CutPrefix(header, bearer)
		if !isBearer {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Detail: "Missing authorization token",
			})
			return
		}

		token, err := s.TokenService.Parse(raw)
		if err != nil {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Err:    err,
				Detail: "Invalid authorization token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), AuthKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
