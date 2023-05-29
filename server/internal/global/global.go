// Package global 全局对象
package global

import (
	"fmt"

	"github.com/silent-rain/gin-admin/internal/pkg/conf"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/mysql"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/redis"
	"github.com/silent-rain/gin-admin/internal/pkg/repository/sqlite"
)

var (
	instance = &globalImpl{}
)

// 全局对象
type globalImpl struct {
	config *conf.Config
	redis  *Redis
	mysql  mysql.DBRepo
	sqlite sqlite.DBRepo
}

// Redis 对象
type Redis struct {
	Default       redis.DBRepo
	UserLogin     redis.DBRepo
	ApiTokenLogin redis.DBRepo
}

// NewGlobal 创建全局对象
func NewGlobal() *globalImpl {
	return &globalImpl{}
}

// 初始化配置
func (g *globalImpl) initConfig() *globalImpl {
	config := conf.New(conf.ConfigFile)
	g.config = config
	return g
}

// 获取全局配置
func (g *globalImpl) Config() *conf.Config {
	return g.config
}

// 初始化 Redis 对象
func (g *globalImpl) initRedis() *globalImpl {
	cfg := g.Config().Redis

	// 用户登录表
	defaultDB, err := redis.New(cfg, redis.Default)
	if err != nil {
		panic(fmt.Sprintf("初始化 Redis 数据库失败! err: %v", err))
	}
	// 用户登录表
	userLoginDB, err := redis.New(cfg, redis.UserLogin)
	if err != nil {
		panic(fmt.Sprintf("初始化 Redis 数据库失败! err: %v", err))
	}
	// 登录信息表
	apiTokenLoginDB, err := redis.New(cfg, redis.ApiTokenLogin)
	if err != nil {
		panic(fmt.Sprintf("初始化 Redis 数据库失败! err: %v", err))
	}

	g.redis = &Redis{
		Default:       defaultDB,
		UserLogin:     userLoginDB,
		ApiTokenLogin: apiTokenLoginDB,
	}
	return g
}

// 获取 Redis 全局对象
func (g *globalImpl) Redis(dbName redis.DBName) redis.DBRepo {

	// redis 缓存
	switch dbName {
	case redis.UserLogin:
		return g.redis.UserLogin
	case redis.ApiTokenLogin:
		return g.redis.ApiTokenLogin
	default:
		return g.redis.Default
	}
}

// 初始化 Sqlite3 全局对象
func (g *globalImpl) initSqlite() *globalImpl {
	cfg := g.Config().Sqlite
	db, err := sqlite.New(*cfg)
	if err != nil {
		panic(fmt.Sprintf("初始化 Mysql 数据库失败! err: %v", err))
	}
	g.sqlite = db
	return g
}

// 初始化 Mysql 全局对象
func (g *globalImpl) initMysql() *globalImpl {
	cfg := g.Config().MySQL
	db, err := mysql.New(cfg.Read, cfg.Write, cfg.Options)
	if err != nil {
		panic(fmt.Sprintf("初始化 Mysql 数据库失败! err: %v", err))
	}
	g.mysql = db
	return g
}

// 获取 Mysql 全局对象
func (g *globalImpl) Mysql() mysql.DBRepo {
	if g.Config().Sqlite != nil {
		return g.sqlite
	}
	return g.mysql
}

// 全局对象初始化
func Init() {
	instance = NewGlobal().
		initConfig().
		initRedis().
		// initSqlite().
		initMysql()
}

// Instance 获取数据库实例
func Instance() *globalImpl {
	return instance
}
