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

import echojwt "github.com/labstack/echo-jwt/v4"

func (s *EchoServer) registerRoutes() {
	healthCheckRoute(s)
	productRoutes(s)
	categoryRoutes(s)
	userRoutes(s)
	cartRoutes(s)
	loginRoute(s)
	orderRoutes(s)

}

func healthCheckRoute(s *EchoServer) {
	s.echo.GET("", s.HealthCheck)
}

func loginRoute(s *EchoServer) {
	authGroup := s.echo.Group("/auth")
	authGroup.POST("/login", s.UserLogin)
}

func cartRoutes(s *EchoServer) {
	categoryGroup := s.echo.Group("/cart")
	categoryGroup.Use(echojwt.WithConfig(JWTConfig()))
	categoryGroup.POST("/", s.AddItemToCart)

}

func categoryRoutes(s *EchoServer) {
	categoryGroup := s.echo.Group("/category")
	categoryGroup.POST("", s.AddCategory)
	categoryGroup.GET("/:id", s.GetCategoryById)
	categoryGroup.PUT("/:id", s.UpdateCategory)
	categoryGroup.DELETE("/:id", s.DeleteCategory)

}

func productRoutes(s *EchoServer) {

	productGroup := s.echo.Group("/products")
	productGroup.GET("", s.GetAllProducts)
	productGroup.GET("/:id", s.GetProductById)
	productGroup.POST("", s.AddProduct)
	productGroup.PUT("/:id", s.UpdateProduct)
	productGroup.DELETE("/:id", s.DeleteProduct)
}

func userRoutes(s *EchoServer) {
	userGroup := s.echo.Group("/user")
	userGroup.POST("", s.AddUser)
	userGroup.GET("/:id", s.GetUserById)
	userGroup.PUT("/:id", s.UpdateUser)
	userGroup.DELETE("/:id", s.DeleteUser)

	userGroup.POST("/address", s.AddUserAddress)
	userGroup.GET("/address/:id", s.GetUserAddressById)
	userGroup.PUT("/address/:id", s.UpdateUserAddress)
	userGroup.DELETE("/address/:id", s.DeleteUserAddress)
}

func orderRoutes(s *EchoServer) {
	orderGroup := s.echo.Group("/order")
	orderGroup.Use(echojwt.WithConfig(JWTConfig()))
	orderGroup.POST("/initiate", s.NewOrder)
	orderGroup.GET("/list", s.ListOrders)
}
