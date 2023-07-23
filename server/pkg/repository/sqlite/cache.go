// Package sqlite Sqlite3 缓存实现
package sqlite

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/silent-rain/gin-admin/pkg/errcode"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CacheTable 内存缓存数据库结构
type CacheTable struct {
	DBName int           `json:"db_name" gorm:"column:db_name;primaryKey"` // 对应 redis 的数据库名称
	Key    string        `json:"key" gorm:"column:key;primaryKey"`         // redis key
	Value  string        `json:"value" gorm:"column:value"`                // redis value
	Expire time.Duration `json:"expire" gorm:"column:expire"`              // redis key 过期时间
}

// TableName 表名重写
func (CacheTable) TableName() string {
	return "redis_mem_cache"
}

// 数据库连接池
type CachePool struct {
	db     *gorm.DB
	dbName int
}

// Set set some <key,value,ttl> into redis
func (d *CachePool) Set(_ context.Context, key, value string, ttl time.Duration) error {
	bean := CacheTable{
		DBName: d.dbName,
		Key:    key,
		Value:  value,
		Expire: ttl,
	}
	err := d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "db_name"}, {Name: "key"}},
		// 主键冲突时, 更新除主键的所有字段
		// UpdateAll: true,
		// 主键冲突时, 更新指定字段
		DoUpdates: clause.AssignmentColumns([]string{"value", "expire"}),
	}).Create(&bean).Error
	if err != nil {
		return errcode.RedisSetKeyError
	}
	return nil
}

// Get 获取 KEY 的值
func (d *CachePool) Get(_ context.Context, key string) (string, error) {
	bean := CacheTable{}
	err := d.db.Where("db_name = ? and key = ?", d.dbName, key).First(&bean).Error
	if err != nil {
		return "", errcode.RedisGetKeyError
	}
	return bean.Value, nil
}

// TTL 查看 Key 剩余的过期时间，以秒为单位。
func (d *CachePool) TTL(ctx context.Context, key string) (time.Duration, error) {
	bean := CacheTable{}
	err := d.db.Where("db_name = ? and key = ?", d.dbName, key).First(&bean).Error
	if err != nil {
		return -1, errcode.RedisTTLGetKeyError
	}
	ttl := time.Now().Nanosecond() - int(bean.Expire.Nanoseconds())
	if ttl < 0 {
		return 0, nil
	}
	return time.Duration(ttl), nil
}

// Expire 设置 key 的过期时间，以秒为单位
func (d *CachePool) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	bean := CacheTable{
		Expire: ttl,
	}
	err := d.db.Where("db_name = ? and key = ?", d.dbName, key).Updates(&bean).Error
	if err != nil {
		return false, errcode.RedisSetKeyExpireError
	}
	return true, nil
}

// ExpireAt 用于为 key 设置过期时间，不同在于，它的时间参数值采用的是时间戳格式。
func (d *CachePool) ExpireAt(ctx context.Context, key string, ttl time.Time) (bool, error) {
	t := time.Duration(ttl.Second())
	bean := CacheTable{
		Expire: t,
	}
	err := d.db.Where("db_name = ? and key = ?", d.dbName, key).Updates(&bean).Error
	if err != nil {
		return false, errcode.RedisSetKeyExpireAtError
	}
	return true, nil
}

// Exists 用于检查指定的一个 key 或者多个 key 是否存在。
// 若存在则返回 1，否则返回 0
func (d *CachePool) Exists(ctx context.Context, keys ...string) (bool, error) {
	if len(keys) == 0 {
		return true, nil
	}
	beans := make([]CacheTable, 0)
	err := d.db.Where("db_name = ? and key in ?", d.dbName, keys).Find(&beans).Error
	if err != nil {
		return false, errcode.RedisGetKeyExistsError
	}
	return len(beans) == len(keys), nil
}

// Del 若键存在的情况下，该命令用于删除键
func (d *CachePool) Del(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return true, nil
	}
	bean := CacheTable{}
	err := d.db.Where("db_name = ? and key = ?", d.dbName, key).Delete(&bean).Error
	if err != nil {
		return false, errcode.RedisDelKeyError
	}
	return true, nil
}

// Incr 将 key 中储存的数字值增一
func (d *CachePool) Incr(ctx context.Context, key string) (int64, error) {
	v, err := d.Get(ctx, key)
	if err != nil {
		return -1, nil
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return -1, errcode.DataTypeConversionError
	}
	value += 1
	bean := CacheTable{
		Value: strconv.Itoa(value),
	}

	if err := d.db.Where("db_name = ? and key = ?", d.dbName, key).Updates(&bean).Error; err != nil {
		return -1, errcode.RedisSetKeyError
	}
	return int64(value), nil
}

// Decr 将 key 中储存的数字值减一
func (d *CachePool) Decr(ctx context.Context, key string) (int64, error) {
	v, err := d.Get(ctx, key)
	if err != nil {
		return -1, nil
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return -1, errcode.DataTypeConversionError
	}
	value -= 1
	bean := CacheTable{
		Value: strconv.Itoa(value),
	}

	if err := d.db.Where("db_name = ? and key = ?", d.dbName, key).Updates(&bean).Error; err != nil {
		return -1, errcode.RedisSetKeyError
	}
	return int64(value), nil
}

// Close 关闭客户端
func (d *CachePool) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// NewCache 创建 Sqlite3 内存对象
// file:memdb1?mode=memory&cache=shared
func NewCache(dns string, dbName int) (*CachePool, error) {
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("[db connection failed]: %w", err)
	}
	pool := &CachePool{
		db:     db,
		dbName: dbName,
	}
	db.Create(&CacheTable{})
	db.AutoMigrate(&CacheTable{})
	return pool, nil
}
