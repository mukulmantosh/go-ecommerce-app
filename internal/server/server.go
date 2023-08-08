package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/database"
	"log"
)

type Server interface {
	Start() error
	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error
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
	err := s.echo.Start(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {

	productGroup := s.echo.Group("/products")
	productGroup.GET("", s.GetAllProducts)
	productGroup.GET("/:id", s.GetProductById)
	productGroup.POST("", s.AddProduct)

}
