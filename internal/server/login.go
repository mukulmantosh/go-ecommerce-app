package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) UserLogin(ctx echo.Context) error {
	login := new(models.Login)
	if err := ctx.Bind(login); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	token, err := s.DB.Login(ctx.Request().Context(), login)
	if err != nil {
		return ctx.JSON(400, map[string]any{"error": err.Error()})
	}
	if len(token) > 0 {
		return ctx.JSON(200, map[string]any{"message": "Thank you! Welcome!", "token": token})
	} else {
		return ctx.JSON(400, map[string]any{"error": "Something went wrong."})
	}

}
