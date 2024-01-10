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
	"github.com/mukulmantosh/go-ecommerce-app/internal/database"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server/abstract"
	"log"
	"log/slog"
)

type Server interface {
	Start() error
	abstract.Product
	abstract.User
	abstract.UserAddress
	abstract.Login
	abstract.Category
	abstract.Cart
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DBClient
}

func NewServer(db database.DBClient) Server {

	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	slog.Info("serving at port 8080")
	err := s.echo.Start(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}
	return nil
}
