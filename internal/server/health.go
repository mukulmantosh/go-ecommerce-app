package server

import (
	"github.com/labstack/echo/v4"
)

func (s *EchoServer) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(200, map[string]any{"message": "OK!"})
}
