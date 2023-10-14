package server

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) NewOrder(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.CustomJWTClaims)
	userId := claims.UserID

	placement, err := s.DB.OrderPlacement(ctx.Request().Context(), userId)
	if err != nil {
		var conflictError *common_errors.ConflictError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	_ = placement

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Thank you, Order Placed!"})
}
