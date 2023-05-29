// Package sqlite Sqlite3 数据库
package sqlite

import (
	"fmt"

	"github.com/silent-rain/gin-admin/internal/pkg/conf"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBRepo 数据库接口
type DBRepo interface {
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
}

// 数据库
type Pool struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

// GetDbR 获取只读数据库对象
func (d *Pool) GetDbR() *gorm.DB {
	return d.DbR
}

// GetDbW 获取读写数据库对象
func (d *Pool) GetDbW() *gorm.DB {
	return d.DbW
}

// DbRClose 关闭只读数据库对象
func (d *Pool) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// DbWClose 关闭读写数据库对象
func (d *Pool) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// New 创建 Sqlite3 对象
func New(cfg conf.SqliteConfig) (DBRepo, error) {
	db, err := gorm.Open(sqlite.Open(cfg.FilePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("[db connection failed]: %w", err)
	}
	pool := &Pool{
		DbR: db,
		DbW: db,
	}
	return pool, nil
}
