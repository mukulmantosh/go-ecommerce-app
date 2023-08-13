package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) AddCategory(ctx echo.Context) error {
	category := new(models.Category)
	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	category, err := s.DB.AddCategory(ctx.Request().Context(), category)
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
	return ctx.JSON(http.StatusCreated, category)
}

func (s *EchoServer) GetCategoryById(ctx echo.Context) error {
	ID := ctx.Param("id")
	category, err := s.DB.GetCategoryById(ctx.Request().Context(), ID)
	if err != nil {
		var notFoundError *common_errors.NotFoundError
		switch {
		case errors.As(err, &notFoundError):
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, category)
}

func (s *EchoServer) UpdateCategory(ctx echo.Context) error {
	ID := ctx.Param("id")
	category := new(models.Category)
	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	categoryInfo, _ := s.DB.GetCategoryById(ctx.Request().Context(), ID)

	if categoryInfo == nil {
		return ctx.JSON(http.StatusBadRequest, "Data Not Found!")
	}
	category.ID = categoryInfo.ID
	updateUser, err := s.DB.UpdateCategory(ctx.Request().Context(), category)
	if err != nil {
		return err
	}
	if err != nil {
		var notFoundError *common_errors.NotFoundError
		var conflictError *common_errors.ConflictError

		switch {
		case errors.As(err, &notFoundError):
			return ctx.JSON(http.StatusNotFound, err)
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	if updateUser == true {
		return ctx.JSON(http.StatusOK, "Category Information Updated!")
	}

	return ctx.JSON(http.StatusInternalServerError, "Something went wrong !")
}

func (s *EchoServer) DeleteCategory(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteCategory(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
