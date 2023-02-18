/*MySQL 数据库插件*/
package mysql

import (
	"gorm.io/gorm"
)

const (
	callBackAfterName  = "created_at"
	callBackBeforeName = "created_at"
)

// 本地时间插件
type LocalTimePlugin struct{}

func (p *LocalTimePlugin) Name() string {
	return "localTimePlugin"
}

func (p *LocalTimePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	// _ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	// _ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	// _ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)

	// 结束后
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	// _ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	// _ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	// _ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	return
}

func before(db *gorm.DB) {
}

func after(db *gorm.DB) {
}
