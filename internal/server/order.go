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
