package server

func (s *EchoServer) registerRoutes() {
	productRoutes(s)
	userRoutes(s)

}

func productRoutes(s *EchoServer) {
	productGroup := s.echo.Group("/products")
	productGroup.GET("", s.GetAllProducts)
	productGroup.GET("/:id", s.GetProductById)
	productGroup.POST("", s.AddProduct)
	productGroup.PUT("/:id", s.UpdateProduct)
	productGroup.PUT("/:id", s.DeleteProduct)
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
