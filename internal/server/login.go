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
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"net/http"
)

func (s *EchoServer) UserLogin(ctx echo.Context) error {
	login := new(models.Login)
	if err := ctx.Bind(login); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	token, err := s.DB.Login(ctx.Request().Context(), login)
	if err != nil {
		return ctx.JSON(400, map[string]any{"error": err.Error()})
	}
	if len(token) > 0 {
		return ctx.JSON(200, map[string]any{"message": "Thank you! Welcome!", "token": token})
	} else {
		return ctx.JSON(400, map[string]any{"error": "Something went wrong."})
	}

}
