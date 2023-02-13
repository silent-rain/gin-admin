/*API 返回结构
 */
package response

import (
	"errors"
	"reflect"
	"testing"

	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
			if got := New(); !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseAPI_Error(t *testing.T) {

	tests := []struct {
		name    string
		code    statuscode.StatusCode
		wantErr error
	}{
		{
			name:    "ok",
			code:    statuscode.Ok,
			wantErr: nil,
		},
		{
			name:    "UnknownError",
			code:    statuscode.UnknownError,
			wantErr: statuscode.UnknownError.Error(),
		},
		{
			name:    "UploadFileParserError",
			code:    statuscode.UploadFileParserError,
			wantErr: statuscode.UploadFileParserError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New().WithCode(tt.code)
			if err := r.Error(); err != tt.wantErr {
				t.Errorf("ResponseAPI.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseAPI_Error2(t *testing.T) {

	tests := []struct {
		name    string
		code    statuscode.StatusCode
		wantErr error
	}{
		{
			name:    "UnknownError",
			code:    statuscode.UnknownError,
			wantErr: statuscode.UnknownError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New().WithCode(tt.code).WithMsg("附加未知的错误消息")
			if err := r.Error(); err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("ResponseAPI.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	var e1 = errors.New("xxx")
	if errors.Is(e1, e1) {
		assert.Equal(t, e1, e1)
	}
}

func TestError(t *testing.T) {
	if errors.Is(statuscode.UnknownError.Error(), statuscode.UnknownError.Error()) {
		assert.Equal(t, statuscode.UnknownError.Error(), statuscode.UnknownError.Error())
	}
}
