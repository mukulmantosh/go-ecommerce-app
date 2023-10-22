package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/database"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server/abstract"
	"log"
	"log/slog"
)

type Server interface {
	Start() error
	abstract.Product
	abstract.User
	abstract.UserAddress
	abstract.Login
	abstract.Category
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DBClient
}

func NewServer(db database.DBClient) Server {

	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	slog.Info("serving at port 8080")
	err := s.echo.Start(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}
	return nil
}
