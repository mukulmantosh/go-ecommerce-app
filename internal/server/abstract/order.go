package abstract

import (
	"github.com/labstack/echo/v4"
)

type Cart interface {
	AddItemToCart(ctx echo.Context) error
	NewOrder(ctx echo.Context) error
	ListOrders(ctx echo.Context) error
}
