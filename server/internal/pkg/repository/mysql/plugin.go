// Package mysql Mysql 数据库插件
package mysql

import (
	"time"

	timeutil "github.com/silent-rain/gin-admin/pkg/utils/time"

	"gorm.io/gorm"
)

// 本地时间插件
type LocalTimePlugin struct{}

func (p *LocalTimePlugin) Name() string {
	return "localTimePlugin"
}

func (p *LocalTimePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	// _ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Create().Before("gorm:before_create").Register("gorm:created_at", beforeByCreate)
	_ = db.Callback().Update().Before("gorm:before_update").Register("gorm:updated_at", beforeByUpdate)
	// _ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)

	// 结束后
	// _ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	// _ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	// _ = db.Callback().Delete().After("gorm:after_delete").Register("gorm:created_at", beforeByCreate)
	// _ = db.Callback().Update().After("gorm:after_update").Register("gorm:updated_at", beforeByUpdate)
	return
}

// 创建注册, 初始化时间
func beforeByCreate(db *gorm.DB) {
	t := time.Now().Format(timeutil.CSTMilliLayout)
	if field := db.Statement.Schema.LookUpField("CreatedAt"); field != nil {
		db.Statement.SetColumn("created_at", t, true)
	}
	if field := db.Statement.Schema.LookUpField("UpdatedAt"); field != nil {
		db.Statement.SetColumn("updated_at", t, true)
	}
}

// 更新注册, 更新时间
func beforeByUpdate(db *gorm.DB) {
	t := time.Now().Format(timeutil.CSTMilliLayout)
	if field := db.Statement.Schema.LookUpField("UpdatedAt"); field != nil {
		db.Statement.SetColumn("updated_at", t)
	}
}
