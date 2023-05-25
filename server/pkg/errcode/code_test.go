// Package errcode 业务状态码
package errcode

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	tests := []struct {
		name    string
		err1    error
		err2    error
		wantErr bool
	}{
		{name: "name1", err1: errors.New("Ok"), err2: errors.New("Ok"), wantErr: false},
		{name: "name2", err1: errors.New("Ok"), err2: fmt.Errorf("Ok"), wantErr: false},
		{name: "name3", err1: fmt.Errorf("Ok"), err2: fmt.Errorf("Ok"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.err1 == tt.err2) != tt.wantErr {
				t.Errorf("statuScode.CodeError() error = %v, wantErr %v", tt.err1, tt.err2)
			}
		})
	}
}

func Test_Errors_Is(t *testing.T) {
	tests := []struct {
		name    string
		err1    error
		err2    error
		wantErr bool
	}{
		{name: "name1", err1: errors.New("Ok"), err2: errors.New("Ok"), wantErr: false},
		{name: "name2", err1: errors.New("Ok"), err2: fmt.Errorf("Ok"), wantErr: false},
		{name: "name3", err1: fmt.Errorf("Ok"), err2: fmt.Errorf("Ok"), wantErr: false},
		{name: "name4", err1: Ok.CodeError(), err2: Ok.CodeError(), wantErr: true},
		{name: "name5", err1: fmt.Errorf("%w: %s", UnknownError.CodeError(), "status Ok"), err2: UnknownError.CodeError(), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if errors.Is(tt.err1, tt.err2) != tt.wantErr {
				t.Errorf("statuScode.CodeError() error = %v, error2 = %v, wantErr %v", tt.err1, tt.err2, tt.wantErr)
			}
		})
	}
}

func Test_statuScode_Error(t *testing.T) {
	tests := []struct {
		name    string
		r       ErrorCode
		msg     string
		wantErr bool
	}{
		{name: "ok", r: InternalError, msg: "内部错误", wantErr: true},
		{name: "Unknown", r: UnknownError, msg: "未知错误", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CodeError(); (err.Error() == tt.msg) != tt.wantErr {
				t.Errorf("statuScode.ErroCodeErrorr() error = %v, msg = %v, wantErr %v", err, tt.msg, tt.wantErr)
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

func TestErrorCode(t *testing.T) {
	var err error = UnknownError.WithMsg("xxxxxxx")
	code, ok := err.(*ErrorMsg)
	if !ok {
		t.Errorf("类型判断失败, ok: %#v", ok)
		return
	}
	if code.Code != UnknownError {
		t.Errorf("错误码异常, code: %#v   msg: %#v", code.Code, code.Err)
		return
	}
	if !errors.Is(code.Err, UnknownError.CodeError()) {
		t.Errorf("类型异常, code: %#v   msg: %#v", code.Code, code.Err)
		return
	}
	assert.Equal(t, code.Err.Error(), "未知错误, xxxxxxx")
}
