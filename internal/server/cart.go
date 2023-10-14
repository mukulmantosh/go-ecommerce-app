package server

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
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
	product := new(models.ProductParams)
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.CustomJWTClaims)
	userId := claims.UserID

	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	if len(product.ProductID) == 0 {
		return ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{"message": "Missing Product ID!"})
	}

	cartData, cartExist, err := s.DB.GetCartInfoByUserID(ctx.Request().Context(), userId)
	if err != nil {
		return err
	}

	if cartExist == 0 {
		// Create a new cart
		cart := new(models.Cart)
		cart.UserID = userId // initialize
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

		// Now add product to the cart.
		_, err = s.DB.AddItemToCart(ctx.Request().Context(), cart.ID, product.ProductID)
		if err != nil {
			return err
		}

	} else {
		_, err := s.DB.AddItemToCart(ctx.Request().Context(), cartData.ID, product.ProductID)
		if err != nil {
			return err
		}
	}

	return ctx.JSON(http.StatusCreated,
		map[string]interface{}{"message": "Item Added to Cart!"})

}