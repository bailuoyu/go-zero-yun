package excel

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/xuri/excelize/v2"
)

const DefaultSheet = "Sheet1"

// ReadExcel 读取 excel 文件
func ReadExcel(filePath string, sheetName string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	sheet := sheetName
	if sheet == "" {
		sheet = DefaultSheet
	}
	rows, err := f.GetRows(sheet)
	return rows, err
}

// WriteExcel 写入 excel 文件
func WriteExcel(outFile string, outData [][]string) error {
	f := excelize.NewFile()
	sheet := DefaultSheet
	sheetIndex, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	for rowKey, row := range outData {
		rowNum := rowKey + 1
		for cellKey, cell := range row {
			colStr, err := getColStr(cellKey + 1)
			if err != nil {
				return err
			}
			f.SetCellValue(sheet, fmt.Sprintf("%s%d", colStr, rowNum), cell)
		}
	}
	f.SetActiveSheet(sheetIndex)
	err = f.SaveAs(outFile)
	return err
}

// 生成excel cell 坐标，eg:A1,A2,BA1,BA2
func getColStr(num int) (string, error) {
	cellStr := ""
	if num <= 26 {
		cellStr = fmt.Sprintf("%c", num+64)
	} else {
		colNum := num / 26
		if colNum > 27 {
			return "", errors.New("字段超过676无法导出")
		}
		cell := int(math.Mod(float64(num), 26))
		if cell == 0 {
			cell = 26
			colNum = colNum - 1
		} else {
			colNum = int(math.Floor(float64(colNum)))
		}
		cellStr = fmt.Sprintf("%c%c", colNum+64, int(cell+64))
	}
	return cellStr, nil
}

// ReadCsv 读取CSV文件
func ReadCsv(filePath string) ([][]string, error) {
	openCsv, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer openCsv.Close()

	readCsv := csv.NewReader(openCsv)
	csvAll, err := readCsv.ReadAll()
	return csvAll, err
}

// WriteCsv 写入CSV文件
func WriteCsv(outFile string, outData [][]string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer file.Close()
	writeCsv := csv.NewWriter(file)
	defer writeCsv.Flush()

	for _, value := range outData {
		err := writeCsv.Write(value)
		if err != nil {
			return err
		}
	}
	return nil
}
