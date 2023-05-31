// Package sqlite Sqlite3 数据库
package sqlite

import (
	"testing"

	"github.com/silent-rain/gin-admin/internal/pkg/conf"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg conf.SqliteConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test",
			args: args{
				conf.SqliteConfig{
					FilePath: "../../../../data.dat",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				t.Errorf("New() = %v", got)
			}
		})
	}
}
