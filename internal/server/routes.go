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
}
