// Package excel excel 读写
package excel

import (
	"testing"
)

func Test_excel_GetRawData(t *testing.T) {
	type fields struct {
		filepath string
		sheet    string
		headers  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		{
			name: "demo",
			fields: fields{
				filepath: "./demo.xlsx",
				sheet:    "Sheet1",
				headers:  []string{"ID", "订单号", "价格", "用户名"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New(tt.fields.filepath).
				WithSheet(tt.fields.sheet).
				WithHeaders(tt.fields.headers).
				Read()
			if e.Error() != nil {
				t.Errorf("excel.Error() = %v", e.Error())
			}
			if got := e.GetRawData(); len(got) == 0 {
				t.Errorf("excel.GetRawData() = %v, want %v", got, tt.want)
			}
		})
	}
}
