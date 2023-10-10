package abstract

import "github.com/labstack/echo/v4"

type Login interface {
	UserLogin(ctx echo.Context) error
}
