package middlewares

import (
	//"github.com/andela/micro-api-gateway/pb/authorization"
	//"github.com/andela/micro-api-gateway/pb/user"
	"github.com/labstack/echo"
)

func SetUserPermissions(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/favicon.ico" {
			return nil
		}
		if c.Path() == "/" {
			return next(c)
		}
		if c.Path() == "/health" {
			return next(c)
		}
		//if c.Request().Header.Get("x-forwarded-proto") == "http" {
		//	to := "https://" + c.Request().Host + c.Request().URL.RequestURI()
		//	return c.Redirect(redirectStatusCode, to)
		//}

		return next(c)
	}
}

func respondWithError(code int, message string, c echo.Context) error {
	return c.JSON(code, echo.Map{"error": message})
}
