package main

import (
	"net/http"

	"github.com/gloryof/go-echo-practice/auth"
	"github.com/gloryof/go-echo-practice/config"
	"github.com/gloryof/go-echo-practice/output"
	"github.com/gloryof/go-echo-practice/render"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	t := config.Template{}
	t.SetBasePath("views/*.html")
	e.Renderer = &t

	e.Static("/lib", "static")

	e.GET("/noauth", func(c echo.Context) error {
		return c.String(http.StatusOK, "Not Authorized!")
	})
	ag := e.Group("/auth")
	ag.Use(middleware.KeyAuthWithConfig(auth.GetConfig()))
	ag.GET("/info", render.View)
	ag.GET("/output/excel", output.Sheet)
	ag.GET("/output/pdf", output.Pdf)

	e.Start(":8000")
}
