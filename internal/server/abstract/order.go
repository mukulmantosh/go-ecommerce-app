package abstract

import (
	"github.com/labstack/echo/v4"
)

type Cart interface {
	CreateNewCart(ctx echo.Context) error
}
