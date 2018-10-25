package middlewares

import (
	"fmt"

	"github.com/labstack/echo"
)

/*
Not particularly interesting
sets a fields in the context in one handler and reads/prints it in another
*/

// SetUser sets a field with a user
func SetUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("userIdFromToken", "12345")
		return next(c)
	}
}

// PrintUser prints the field from the context
func PrintUser(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		fmt.Println("yada yada user:", c.Get("userIdFromToken"))
		return next(c)
	}
}
