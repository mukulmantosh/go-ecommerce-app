package server

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) AddUser(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	user, err := s.DB.AddUser(ctx.Request().Context(), user)
	if err != nil {
		var conflictError *common_errors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": conflictError.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	return ctx.JSON(http.StatusCreated,
		map[string]interface{}{"message": "User Created Successfully !",
			"data": map[string]interface{}{"user_id": user.ID}})
}

func (s *EchoServer) GetUserById(ctx echo.Context) error {
	ID := ctx.Param("id")
	if len(ID) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"error": "ID Missing"})
	}
	user, err := s.DB.GetUserById(ctx.Request().Context(), ID)
	if err != nil {
		var notFoundError *common_errors.NotFoundError
		switch {
		case errors.As(err, &notFoundError):
			return ctx.JSON(http.StatusNotFound, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	return ctx.JSON(http.StatusOK, user)
}

func (s *EchoServer) UpdateUser(ctx echo.Context) error {
	ID := ctx.Param("id")
	if len(ID) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"error": "ID Missing"})
	}

	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	userInfo, _ := s.DB.GetUserById(ctx.Request().Context(), ID)

	if userInfo == nil {
		return ctx.JSON(http.StatusBadRequest, "Data Not Found!")
	}
	user.ID = userInfo.ID
	updateUser, err := s.DB.UpdateUser(ctx.Request().Context(), user)
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
		return ctx.JSON(http.StatusOK, map[string]any{"message": "User Information Updated!"})
	}

	return ctx.JSON(http.StatusInternalServerError, "Something went wrong !")
}

func (s *EchoServer) DeleteUser(ctx echo.Context) error {
	ID := ctx.Param("id")
	if len(ID) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"error": "ID Missing"})
	}
	err := s.DB.DeleteUser(ctx.Request().Context(), ID)
	fmt.Println(err)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusResetContent)
}
