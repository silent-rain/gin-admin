// Package model 菜单, 路由/按钮/操作权限
package model

// MenuSortById 通过对 ID 排序实现了 sort.Interface 接口
type MenuSortById []*Menu

// Len 数据长度
func (a MenuSortById) Len() int { return len(a) }

// Less 数据比较
func (a MenuSortById) Less(i, j int) bool { return a[i].ID < a[j].ID }

// Swap 数据交换位置
func (a MenuSortById) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 菜单类型
type MenuType int

const (
	MenuTypeByMenu MenuType = iota
	MenuTypeByButton
)

// 菜单打开类型
type MenuOpenType int

const (
	MenuOpenTypeByComponent MenuOpenType = iota
	MenuOpenTypeByInnerLink
	MenuOpenTypeByOutLink
)

// 菜单是否可见类型
type MenuHideType int

const (
	MenuHideTypeByShow MenuHideType = iota // 显示
	MenuHideTypeByHide                     // 隐藏
)
