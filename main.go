package main

import (
	"net/http"
	"strings"

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

	go func() {

		r := echo.New()

		r.Pre(redirectMainPort())

		r.Start(":8001")
	}()

	e.Start(":8000")
}

func redirectMainPort() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			r := c.Request()
			bh := strings.Split(r.Host, ":")[0]

			np := c.Scheme() + "://" + bh + ":8000" + r.RequestURI

			return c.Redirect(http.StatusMovedPermanently, np)

		}
	}
}
