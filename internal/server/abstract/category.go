package abstract

import "github.com/labstack/echo/v4"

type Category interface {
	AddCategory(ctx echo.Context) error
	GetCategoryById(ctx echo.Context) error
	UpdateCategory(ctx echo.Context) error
	DeleteCategory(ctx echo.Context) error
}
