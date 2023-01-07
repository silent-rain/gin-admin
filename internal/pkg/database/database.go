/*
 * @Author: silent-rain
 * @Date: 2023-01-07 20:54:33
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 21:51:16
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/database/database.go
 * @Descripttion: 数据库
 */
package database

import (
	"time"

	"gin-admin/internal/pkg/conf"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// 数据库初始化
func Init() {
	// 获取 db 实例
	if conf.Instance().DBConfig == nil {
		db = initsqlite()
	} else {
		db = initMysql()
	}

	// 获取连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取数据库连接池失败")
	}

	// 验证与数据库的连接是否仍然有效，必要时建立连接。
	if err := sqlDB.Ping(); err != nil {
		panic("数据库连接失败")
	}

	// 设置空闲连接池中连接的最大数量
	// 如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// 初始化 mysql 实例
func initMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Instance().DBConfig.Dsn()), &gorm.Config{
		// 设置时区
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		panic("数据库初始化失败")
	}
	return db
}

// 初始化 sqlite3 实例
func initsqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(conf.Instance().SqliteConfig.FilePath), &gorm.Config{
		// 设置时区
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		panic("数据库初始化失败")
	}
	return db
}

// 获取 gorm db 对象，其他包需要执行数据库查询的时候，
// 不用担心协程并发使用同样的 db 对象会共用同一个连接，
// db 对象在调用他的方法的时候会从数据库连接池中获取新的连接
func Instance() *gorm.DB {
	return db
}
