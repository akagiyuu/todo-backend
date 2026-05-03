package server

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/gorilla/schema"

	"github.com/akagiyuu/todo-api/internal/auth"
	"github.com/akagiyuu/todo-api/internal/todo"
)

type Server struct {
	Config Config

	Decoder *schema.Decoder

	AuthService  *auth.AuthService
	TokenService *auth.TokenService
	TodoService  *todo.TodoService
}

func NewServer(
	authService *auth.AuthService,
	todoService *todo.TodoService,
) (*Server, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}

	return &Server{
		Config:       cfg,
		Decoder:      schema.NewDecoder(),
		AuthService:  authService,
		TokenService: authService.TokenService,
		TodoService:  todoService,
	}, nil
}

func (s *Server) Build() *fuego.Server {
	f := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", s.Config.Port)),
		fuego.WithGlobalMiddlewares(s.CorsMiddleware),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				UIHandler:            s.OpenAPIHandler,
				DisableDefaultServer: true,
				DisableMessages:      true,
				Info: &openapi3.Info{
					Title:       "General Service",
					Description: "General Service",
				},
			}),
		),
		fuego.WithSecurity(openapi3.SecuritySchemes{
			"bearerAuth": &openapi3.SecuritySchemeRef{
				Value: openapi3.NewSecurityScheme().
					WithType("http").
					WithScheme("bearer").
					WithBearerFormat("JWT").
					WithDescription("Enter your JWT token in the format: Bearer <token>"),
			},
		}),
	)
	s.RegisterRoutes(f)

	return f
}
