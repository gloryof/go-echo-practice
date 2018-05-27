package main

import (
	"net/http"

	"github.com/gloryof/go-echo-practice/auth"
	"github.com/gloryof/go-echo-practice/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	t := config.Template{}
	t.SetBasePath("views/*.html")
	e.Renderer = &t

	e.GET("/noauth", func(c echo.Context) error {
		return c.String(http.StatusOK, "Not Authorized!")
	})
	ag := e.Group("/auth")
	ag.Use(middleware.KeyAuthWithConfig(auth.GetConfig()))
	ag.GET("/info", func(c echo.Context) error {

		view := ViewInfo{
			Title: "テスト",
			Data: []ViewData{
				{
					ID:    1,
					Value: "test1",
				},
				{
					ID:    2,
					Value: "test2",
				},
				{
					ID:    3,
					Value: "test3",
				},
			},
		}

		return c.Render(http.StatusOK, "index", view)
	})

	e.Start(":8000")
}

// ViewInfo 表示情報
type ViewInfo struct {
	Title string
	Data  []ViewData
}

// ViewData 表示データ
type ViewData struct {
	ID    uint64
	Value string
}
