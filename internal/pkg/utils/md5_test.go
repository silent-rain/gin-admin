/*
 * @Author: silent-rain
 * @Date: 2023-01-08 15:52:19
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 16:16:06
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/md5_test.go
 * @Descripttion: MD5 加密
 */
package utils

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
			if got := Md5(tt.args.v); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
