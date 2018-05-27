package config

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Template テンプレートエンジンの設定
type Template struct {
	templates *template.Template
}

// SetBasePath ベースパスを設定する
func (t *Template) SetBasePath(path string) {
	t.templates = template.Must(template.ParseGlob(path))
}

// Render レンダリング処理（echo.Rendererの実装）
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
