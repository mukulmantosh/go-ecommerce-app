package server

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"net/http"
)

func (s *EchoServer) NewOrder(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	//claims := user.Claims.(*models.CustomJWTClaims)
	claims, err := ParseJWTToken(user.Raw)

	userId := claims.UserID

	_, err = s.DB.OrderPlacement(ctx.Request().Context(), userId)
	if err != nil {
		var conflictError *common_errors.ConflictError
		var cartEmptyError *common_errors.CartEmptyError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		case errors.As(err, &cartEmptyError):
			return ctx.JSON(http.StatusNotFound, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Thank you, Order Placed!"})
}

func (s *EchoServer) ListOrders(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	//claims := user.Claims.(*models.CustomJWTClaims)
	claims, err := ParseJWTToken(user.Raw)
	userId := claims.UserID

	orderList, err := s.DB.OrderList(ctx.Request().Context(), userId)
	if err != nil {
		var conflictError *common_errors.ConflictError
		var cartEmptyError *common_errors.CartEmptyError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		case errors.As(err, &cartEmptyError):
			return ctx.JSON(http.StatusNotFound, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	return ctx.JSON(http.StatusOK, orderList)
}
