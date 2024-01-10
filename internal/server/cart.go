/*
	Copyright 2022 Google LLC

#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/
package server

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) AddItemToCart(ctx echo.Context) error {
	product := new(models.ProductParams)
	user := ctx.Get("user").(*jwt.Token)
	//claims := user.Claims.(*models.CustomJWTClaims)
	claims, err := ParseJWTToken(user.Raw)
	userId := claims.UserID

	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}

	if len(product.ProductID) == 0 {
		return ctx.JSON(http.StatusBadRequest,
			map[string]any{"message": "Missing Product ID!"})
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
				return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
			default:
				return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
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
		map[string]any{"message": "Item Added to Cart!"})

}
