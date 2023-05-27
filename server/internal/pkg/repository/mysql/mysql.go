// Package mysql Mysql 数据库
package mysql

import (
	"fmt"
	"time"

	"github.com/silent-rain/gin-admin/internal/pkg/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dbInstance DBRepo
)

// DBRepo 数据库接口
type DBRepo interface {
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
}

// New 新建数据库对象
func New(dbCfgR, dbCfgW conf.MySQLAuthConfig, options conf.MySQLOptionsConfig) (DBRepo, error) {
	dbr, err := dbConnect(dbCfgR, options)
	if err != nil {
		return nil, err
	}
	dbw, err := dbConnect(dbCfgW, options)
	if err != nil {
		return nil, err
	}

	return &dbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}

// 数据库
type dbRepo struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

// GetDbR 获取只读数据库对象
func (d *dbRepo) GetDbR() *gorm.DB {
	return d.DbR
}

// GetDbW 获取读写数据库对象
func (d *dbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

// DbRClose 关闭只读数据库对象
func (d *dbRepo) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// DbWClose 关闭读写数据库对象
func (d *dbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// 连接数据库
func dbConnect(dbCfg conf.MySQLAuthConfig, options conf.MySQLOptionsConfig) (*gorm.DB, error) {
	// 数据库地址
	dsn := SourceDsn(dbCfg)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 命名策略表，列命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// 更改创建时间使用的函数
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		// 日志配置
		Logger: logger.Default.LogMode(options.GetLogLevel()),
	})
	if err != nil {
		return nil, fmt.Errorf("[db connection failed] Database name: %s, %w", dbCfg.DbName, err)
	}
	// 设置表字符类型
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 验证与数据库的连接是否仍然有效，必要时建立连接。
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("[db connection failed] Database name: %s, %w", dbCfg.DbName, err)
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(options.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(options.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * options.ConnMaxLifeTime)

	// 使用插件
	db.Use(&LocalTimePlugin{})

	return db, nil
}

// Dsn 拼接 mysql 数据库地址
func SourceDsn(cfg conf.MySQLAuthConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		false,
		"Asia%2FShanghai", // or Local
	)
}

// Init 初始化数据库
func Init() error {
	cfg := conf.Instance().MySQL
	db, err := New(cfg.Read, cfg.Write, cfg.Options)
	if err != nil {
		panic(fmt.Sprintf("初始化 Mysql 数据库失败! err: %v", err))
	}
	dbInstance = db
	return err
}

// Instance 获取数据库实例
func Instance() DBRepo {
	return dbInstance
}
