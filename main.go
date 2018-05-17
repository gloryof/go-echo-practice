package main

import (
	"net/http"

	"github.com/gloryof/go-echo-practice/auth"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.GET("/noauth", func(c echo.Context) error {
		return c.String(http.StatusOK, "Not Authorized!")
	})
	ag := e.Group("/auth")
	ag.Use(middleware.KeyAuthWithConfig(auth.GetConfig()))
	ag.GET("/info", func(c echo.Context) error {

		return c.String(http.StatusOK, "Authorized!")
	})

	e.Start(":8000")
}
