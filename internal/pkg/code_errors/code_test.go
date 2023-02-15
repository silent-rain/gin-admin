/*业务状态码
 */
package code_errors

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
		r       StatusCode
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

func TestErrors(t *testing.T) {
	var e1 = errors.New("xxx")
	if errors.Is(e1, e1) {
		assert.Equal(t, e1, e1)
	}
}
