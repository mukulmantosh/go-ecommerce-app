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

func (s *EchoServer) AddUserAddress(ctx echo.Context) error {
	userAddress := new(models.UserAddress)
	if err := ctx.Bind(userAddress); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	userAddress, err := s.DB.AddUserAddress(ctx.Request().Context(), userAddress)
	if err != nil {
		var conflictError *common_errors.ConflictError
		var violationError *common_errors.ViolationError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		case errors.As(err, &violationError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, userAddress)
}
