/*
 * @Author: silent-rain
 * @Date: 2023-01-07 16:56:09
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 20:39:26
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/status_code/status_code_test.go
 * @Descripttion:
 */
/**业务状态码
 */
package statuscode

import (
	"errors"
	"fmt"
	"testing"
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
				t.Errorf("statuScode.Error() error = %v, wantErr %v", tt.err1, tt.err2)
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
		{name: "name3", err1: Ok.Error(), err2: Ok.Error(), wantErr: true},
		{name: "name4", err1: fmt.Errorf("%w: %s", Ok.Error(), "status Ok"), err2: Ok.Error(), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if errors.Is(tt.err1, tt.err2) != tt.wantErr {
				t.Errorf("statuScode.Error() error = %v, error2 = %v, wantErr %v", tt.err1, tt.err2, tt.wantErr)
			}
		})
	}
}

func Test_statuScode_Error(t *testing.T) {
	tests := []struct {
		name    string
		r       StatuScode
		msg     string
		wantErr bool
	}{
		{name: "ok", r: Ok, msg: "Ok", wantErr: true},
		{name: "Unknown", r: UnknownError, msg: "Unknown", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Error(); (err.Error() == tt.msg) != tt.wantErr {
				t.Errorf("statuScode.Error() error = %v, msg = %v, wantErr %v", err, tt.msg, tt.wantErr)
			}
		})
	}
}
