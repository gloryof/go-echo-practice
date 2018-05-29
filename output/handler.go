package output

import (
	"github.com/labstack/echo"
	"github.com/tealeg/xlsx"
)

// Sheet シートの出力
func Sheet(c echo.Context) error {

	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	sheet.Cell(0, 0).Value = "1"

	return file.Write(c.Response().Writer)
}
