// Package excel excel 读写
package excel

import (
	"encoding/json"

	"github.com/xuri/excelize/v2"
)

// excel 对象
type excel struct {
	filepath string
	sheet    string
	headers  []string
	data     [][]string
	err      error
}

// 创建 excel 对象
func New(filepath string) *excel {
	return &excel{
		filepath: filepath,
		sheet:    "Sheet1",
		err:      nil,
	}
}

// WithSheet 设置 sheet
func (e *excel) WithSheet(sheet string) *excel {
	e.sheet = sheet
	return e
}

// WithHeaders 设置 Header 标题
func (e *excel) WithHeaders(headers []string) *excel {
	e.headers = headers
	return e
}

// Error 获取错误信息
func (e *excel) Error() error {
	return e.err
}

// Read 读取 Excel 文件
func (e *excel) Read() *excel {
	f, err := excelize.OpenFile(e.filepath)
	if err != nil {
		e.err = err
		return e
	}

	defer f.Close()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(e.sheet)
	if err != nil {
		e.err = err
		return e
	}

	e.data = rows
	return e
}

// GetHeader 获取 Headers 列表
// 如没有指定 Headers 第一行将作为 Headers
func (e *excel) GetHeader() []string {
	headers := make([]string, 0)
	if len(e.headers) == 0 {
		headers = e.data[1]
	}
	return headers
}

// GetRawData 获取原始数据
func (e *excel) GetRawData() [][]string {
	if len(e.data) == 0 {
		return [][]string{}
	}
	return e.data
}

// GetMapData 获取 map 数据
// 如没有指定 Headers 第一行将作为 Headers
func (e *excel) GetMapData() []map[string]string {
	if len(e.data) == 0 {
		return nil
	}
	dataList := make([]map[string]string, 0)
	n := 1
	headers := make([]string, 0)
	if len(e.headers) == 0 {
		headers = e.data[1]
		n = 1
	}
	for _, row := range e.data[n:] {
		m := make(map[string]string, 0)
		for i, col := range row {
			m[headers[i]] = col
		}
		dataList = append(dataList, m)
	}
	return dataList
}

// Unmarshal 解析为指定类型数据
func (e *excel) Unmarshal(value interface{}) error {
	dataList := e.GetMapData()
	buf, err := json.Marshal(dataList)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf, value); err != nil {
		return err
	}
	return nil
}
