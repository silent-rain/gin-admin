/*
 * @Author: silent-rain
 * @Date: 2023-01-07 17:45:55
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 19:26:25
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/response/response_test.go
 * @Descripttion: API 返回结构
 */
package response

import (
	statuscode "gin-admin/internal/pkg/status_code"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNew(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *ResponseAPI
	}{
		{name: "", args: args{c: nil}, want: &ResponseAPI{
			Code: statuscode.Ok,
			Msg:  statuscode.Ok.Msg(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.c); !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
