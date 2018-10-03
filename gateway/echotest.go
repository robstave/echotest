package gateway

import (
	"echotest/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunEcho is the actual gateway
func RunEcho(port int64) {

	portString := ":" + strconv.FormatInt(port, 10)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/users", controllers.GetUsers)
	e.GET("/user/:id", controllers.GetUser)

	// Server
	e.Logger.Fatal(e.Start(portString))

}
