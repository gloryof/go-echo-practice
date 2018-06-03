package output

import (
	"net/http"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
)

// Sheet シートの出力
func Sheet(c echo.Context) error {

	file := excelize.NewFile()

	targetSheet := "Sheet1"
	addData(file, targetSheet)

	file.AddChart(targetSheet, "B15", createChartOption())

	file.SetActiveSheet(file.NewSheet(targetSheet))
	return file.Write(c.Response().Writer)
}

// Pdf PDFを出力する
func Pdf(c echo.Context) error {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {

		return err
	}

	html := `<html>
			  <head>
			  	<meta charset="UTF-8">
			  	<title>テスト</title>
			  </head>
			  <body>
			  	テストだよー
			  </body>
			 </html>`

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	err2 := pdfg.Create()
	if err2 != nil {

		return err2
	}

	return c.Stream(http.StatusOK, "application/pdf", pdfg.Buffer())
}

func createChartOption() string {

	return `
		{
			"type": "line",
			"dimension": {
				"width": 640,
				"height": 640
			},
			"series": [{
				"name": "Sheet1!$C2",
				"categories": "Sheet1!$B$3:$B$13",
				"values": "Sheet1!$C$3:$C$13"
			}],
			"title": {
				"name": "体重グラフ"
			},
			"y_axis": {
				"maximum": 68.0,
				"minimum": 58.0
			}
		}
	`
}

func addData(file *excelize.File, targetSheet string) {

	setTitle(file, targetSheet)
	addRow(file, targetSheet, "3", "2018/04/25", 63.3)
	addRow(file, targetSheet, "4", "2018/04/28", 63.8)
	addRow(file, targetSheet, "5", "2018/05/02", 63.2)
	addRow(file, targetSheet, "6", "2018/05/05", 63.8)
	addRow(file, targetSheet, "7", "2018/05/09", 63.4)
	addRow(file, targetSheet, "8", "2018/05/12", 63.6)
	addRow(file, targetSheet, "9", "2018/05/16", 62.0)
	addRow(file, targetSheet, "10", "2018/05/19", 62.8)
	addRow(file, targetSheet, "11", "2018/05/23", 61.5)
	addRow(file, targetSheet, "12", "2018/05/27", 64.0)
	addRow(file, targetSheet, "13", "2018/05/30", 61.8)

}

func setTitle(s *excelize.File, targetSheet string) {

	s.SetCellValue(targetSheet, "B2", "日付")
	s.SetCellValue(targetSheet, "C2", "体重")
}

func addRow(s *excelize.File, targetSheet string, rowNum string, date string, weight float64) {

	s.SetCellValue(targetSheet, "B"+rowNum, date)
	s.SetCellValue(targetSheet, "C"+rowNum, weight)
}
