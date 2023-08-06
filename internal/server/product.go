package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-ecommerce-app/internal/generic/common_errors"
	"go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	products, err := s.DB.AllProducts(ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		var conflictError *common_errors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}
