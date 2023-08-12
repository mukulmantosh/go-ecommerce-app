package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) AddUser(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	user, err := s.DB.AddUser(ctx.Request().Context(), user)
	if err != nil {
		var conflictError *common_errors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, user)
}

func (s *EchoServer) GetUserById(ctx echo.Context) error {
	ID := ctx.Param("id")
	user, err := s.DB.GetUserById(ctx.Request().Context(), ID)
	if err != nil {
		var notFoundError *common_errors.NotFoundError
		switch {
		case errors.As(err, &notFoundError):
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, user)
}
