package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) AddUserAddress(ctx echo.Context) error {
	userAddress := new(models.UserAddress)
	if err := ctx.Bind(userAddress); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	userAddress, err := s.DB.AddUserAddress(ctx.Request().Context(), userAddress)
	if err != nil {
		var conflictError *common_errors.ConflictError
		var violationError *common_errors.ViolationError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		case errors.As(err, &violationError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	return ctx.JSON(http.StatusCreated, userAddress)
}

func (s *EchoServer) GetUserAddressById(ctx echo.Context) error {
	ID := ctx.Param("id")
	if len(ID) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"error": "ID Missing"})
	}
	userAddress, err := s.DB.GetUserAddressById(ctx.Request().Context(), ID)
	if err != nil {
		var notFoundError *common_errors.NotFoundError
		switch {
		case errors.As(err, &notFoundError):
			return ctx.JSON(http.StatusNotFound, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	return ctx.JSON(http.StatusOK, userAddress)
}

func (s *EchoServer) UpdateUserAddress(ctx echo.Context) error {
	ID := ctx.Param("id")
	if len(ID) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"error": "ID Missing"})
	}
	userAddress := new(models.UserAddress)
	if err := ctx.Bind(userAddress); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	userAddressInfo, _ := s.DB.GetUserAddressById(ctx.Request().Context(), ID)

	if userAddressInfo == nil {
		return ctx.JSON(http.StatusBadRequest, "Data Not Found!")
	}
	userAddress.ID = userAddressInfo.ID
	updateUser, err := s.DB.UpdateUserAddress(ctx.Request().Context(), userAddress)
	if err != nil {
		return err
	}
	if err != nil {
		var notFoundError *common_errors.NotFoundError
		var conflictError *common_errors.ConflictError

		switch {
		case errors.As(err, &notFoundError):
			return ctx.JSON(http.StatusNotFound, map[string]any{"error": err.Error()})
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	if updateUser {
		return ctx.JSON(http.StatusOK, "User Information Updated!")
	}

	return ctx.JSON(http.StatusInternalServerError, "Something went wrong !")
}

func (s *EchoServer) DeleteUserAddress(ctx echo.Context) error {
	ID := ctx.Param("id")
	if len(ID) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"error": "ID Missing"})
	}
	err := s.DB.DeleteUserAddress(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusResetContent)
}
