package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) CreateNewCart(ctx echo.Context) error {
	cart := new(models.Cart)
	if err := ctx.Bind(cart); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	cart, err := s.DB.NewCartForUser(ctx.Request().Context(), cart)
	if err != nil {
		var conflictError *common_errors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated,
		map[string]interface{}{"message": "New Cart Created!",
			"data": map[string]interface{}{"cart_id": cart.ID}})
}

func (s *EchoServer) AddItemToCart(ctx echo.Context) error {
	cartId := "dab3efe8-714e-48dc-be60-e2bdb084a499"
	productId := "f3abff1d-1f3a-42bb-a912-b6459ac09f98"

	_, err := s.DB.AddItemToCart(ctx.Request().Context(), cartId, productId)
	if err != nil {
		var conflictError *common_errors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated,
		map[string]interface{}{"message": "Item Added to Cart!"})
}
