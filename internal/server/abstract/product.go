package abstract

import "github.com/labstack/echo/v4"

type Product interface {
	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error
	DeleteProduct(ctx echo.Context) error
}
