package middlewares

import (
	"github.com/labstack/echo"
)

// ServerHeader is a simple middleware that adds a littel something something
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Leggo my Echo")
		return next(c)
	}
}
