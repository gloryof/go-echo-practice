package render

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// View 表示処理
func View(c echo.Context) error {

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
