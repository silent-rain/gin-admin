// Package md5 MD5 加密
package md5

import (
	"testing"
)

func TestMd5(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "na", args: args{v: "xxxx"}, want: "c412593ca5c43077c282dc21257f22e6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptMd5(tt.args.v); got != tt.want {
				t.Errorf("EncryptMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}
