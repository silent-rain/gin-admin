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

import "gin-admin/internal/pkg/utils"

// Menu 菜单
type Menu struct {
	ID           uint             `json:"id" gorm:"column:id;primaryKey"`                           // 菜单ID
	ParentId     *uint            `json:"parent_id" gorm:"column:parent_id"`                        // 父菜单ID
	Title        string           `json:"title" gorm:"column:title"`                                // 菜单名称
	Icon         string           `json:"icon" gorm:"column:icon"`                                  // 菜单图标
	ElSvgIcon    string           `json:"el_svg_icon" gorm:"column:el_svg_icon"`                    // Element 菜单图标
	MenuType     uint             `json:"menu_type" gorm:"column:menu_type"`                        // 菜单类型,0:菜单,1:按钮
	OpenType     uint             `json:"open_type" gorm:"column:open_type"`                        // 打开方式,0:组件,1:内链,2:外链
	Path         string           `json:"path" gorm:"column:path"`                                  // 路由地址/外链地址
	Name         string           `json:"name" gorm:"column:name"`                                  // 路由别名
	Component    string           `json:"component" gorm:"column:component"`                        // 组件路径
	Redirect     string           `json:"redirect" gorm:"column:redirect"`                          // 路由重定向
	Link         string           `json:"link" gorm:"column:link"`                                  // 链接地址: 内链地址
	Target       string           `json:"target" gorm:"column:target"`                              // 链接地址跳转方式, _blank/_self
	Permission   string           `json:"permission" gorm:"column:permission"`                      // 权限标识
	Hidden       uint             `json:"hidden" gorm:"column:hidden;type:tinyint"`                 // 是否隐藏,0:显示,1:隐藏
	AlwaysShow   uint             `json:"always_show" gorm:"column:always_show"`                    // 始终显示根菜单,0:显示,1:隐藏
	Sort         uint             `json:"sort" gorm:"column:sort"`                                  // 排序
	Note         string           `json:"note" gorm:"column:note"`                                  // 备注
	Status       uint             `json:"status" gorm:"column:status"`                              // 状态,0:停用,1:启用
	CreateUserId uint             `json:"-" gorm:"column:create_user_id;"`                          // 创建菜单用户ID
	UpdateUserId uint             `json:"-" gorm:"column:update_user_id;"`                          // 更新菜单用户ID
	CreatedAt    *utils.LocalTime `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"` // 创建时间
	UpdatedAt    *utils.LocalTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"` // 更新时间
	Children     []Menu           `json:"children" gorm:"foreignKey:ParentId"`                      // 子菜单列表
}

// TableName 表名重写
func (Menu) TableName() string {
	return "sys_menu"
}

// MenuSortById 通过对 ID 排序实现了 sort.Interface 接口
type MenuSortById []*Menu

// Len 数据长度
func (a MenuSortById) Len() int { return len(a) }

// Less 数据比较
func (a MenuSortById) Less(i, j int) bool { return a[i].ID < a[j].ID }

// Swap 数据交换位置
func (a MenuSortById) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 菜单类型
type MenuType uint

const (
	MenuTypeByMenu MenuType = iota
	MenuTypeByBUtton
)

// 菜单是否可见类型
type MenuHideType uint

const (
	MenuHideTypeByShow MenuHideType = iota // 显示
	MenuHideTypeByHide                     // 隐藏
)
