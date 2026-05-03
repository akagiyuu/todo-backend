package http

import (
	"fmt"
	"todo-backend/internal/auth"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
)

type Server struct {
	Config *Config

	AuthService *auth.AuthService
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
