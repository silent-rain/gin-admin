/*
 * @Author: silent-rain
 * @Date: 2023-01-14 23:06:03
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 23:41:24
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/model/system/menu.go
 * @Descripttion: 菜单, 路由/按钮/操作权限
 */
package systemModel

import "time"

// Menu 菜单
type Menu struct {
	ID           uint      `gorm:"column:id;primaryKey"`                   // 菜单ID
	ParentId     *uint     `gorm:"column:parent_id"`                       // 父菜单ID
	Title        string    `gorm:"column:title"`                           // 菜单名称
	Icon         string    `gorm:"column:icon"`                            // 菜单图标
	OpenType     string    `gorm:"column:open_type"`                       // 打开方式,0:组件,1:内链,2:外联
	Path         string    `gorm:"column:path"`                            // 路由地址/外链地址
	Component    string    `gorm:"column:component"`                       // 组件路径/内链地址
	Target       string    `gorm:"column:target"`                          // 链接地址跳转方式, _blank/_self
	Permission   string    `gorm:"column:permission"`                      // 权限标识
	MenuType     string    `gorm:"column:menu_type"`                       // 菜单类型,0:菜单,1:按钮
	Hide         string    `gorm:"column:hide"`                            // 菜单是否隐藏,0:隐藏,1:显示
	Sort         uint      `gorm:"column:sort"`                            // 排序
	Note         string    `gorm:"column:note"`                            // 备注
	Status       uint      `gorm:"column:status"`                          // 状态,0:停用,1:启用
	CreateUserId uint      `gorm:"column:create_user_id;"`                 // 创建菜单用户ID
	UpdateUserId uint      `gorm:"column:update_user_id;"`                 // 更新菜单用户ID
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
	Children     []Menu    `gorm:"foreignKey:ParentId"`                    // 子菜单列表
}

// TableName 表名重写
func (HttpLog) Menu() string {
	return "sys_menu"
}
