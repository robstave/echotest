package gateway

import (
	"echotest/controllers"
	mw "echotest/middlewares"
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

	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := mw.NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle) // Endpoint to get stats

	e.File("/", "public/index.html")
	e.Use(mw.ServerHeader)

	e.Use(mw.SetUser)
	e.Use(mw.PrintUser)

	// Route => handler
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/users", controllers.GetUsers)
	e.GET("/user/:id", controllers.GetUser)

	// Server
	e.Logger.Fatal(e.Start(portString))

}
