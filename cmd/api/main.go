package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akagiyuu/todo-api/internal/auth"
	"github.com/akagiyuu/todo-api/internal/database"
	server "github.com/akagiyuu/todo-api/internal/transports/http"

	"github.com/go-fuego/fuego"
)

func gracefulShutdown(apiServer *fuego.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}

func main() {
	pool, err := database.NewPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	tokenService, err := auth.NewTokenService()
	if err != nil {
		panic(err)
	}

	authService := &auth.AuthService{
		Pool:         pool,
		TokenService: tokenService,
	}

	s, err := server.NewServer(authService)
	if err != nil {
		panic(err)
	}
	server := s.Build()

	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err = server.Run()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
