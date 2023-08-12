package abstract

import "github.com/labstack/echo/v4"

type User interface {
	AddUser(ctx echo.Context) error
	AddUserAddress(ctx echo.Context) error
}
