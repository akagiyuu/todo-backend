package server

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (s *Server) RegisterRoutes(f *fuego.Server) {
	authOptions := []fuego.RouteOption{
		option.Middleware(s.AuthMiddleware),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	}

	fuego.Get(f, "/", s.PingHandler)

	auth := fuego.Group(f, "/auth")
	fuego.Post(auth, "/login", s.LoginHandler)
	fuego.Post(auth, "/register", s.RegisterHandler)

	todo := fuego.Group(f, "/todo", authOptions...)
	fuego.Post(todo, "/", s.CreateTodo)
	fuego.Get(todo, "/", s.FilterTodo,
		option.Query("query", "Title or content of todo", param.Nullable()),
		option.Query("priority", "Priority of todo", param.Nullable()),
		option.QueryBool("isDone", "Status of todo", param.Nullable()),
	)
	fuego.Get(todo, "/{id}", s.GetTodo)
	fuego.Patch(todo, "/{id}", s.UpdateTodo)
	fuego.Delete(todo, "/{id}", s.DeleteTodo)
}

func (s *Server) OpenAPIHandler(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}

func (s *Server) PingHandler(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}
