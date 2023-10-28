package abstract

import "github.com/labstack/echo/v4"

type Health interface {
	HealthCheck(ctx echo.Context) error
}
