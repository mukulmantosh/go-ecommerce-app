package abstract

import "github.com/labstack/echo/v4"

type User interface {
	AddUser(ctx echo.Context) error
	GetUserById(ctx echo.Context) error
}

type UserAddress interface {
	AddUserAddress(ctx echo.Context) error
}
