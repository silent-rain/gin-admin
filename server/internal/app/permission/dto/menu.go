// Package dto 菜单
package dto

import (
	DTO "github.com/silent-rain/gin-admin/internal/dto"
)

// QueryMenuReq 查询条件
type QueryMenuReq struct {
	DTO.Pagination        // 分页
	Title          string `json:"title" form:"title"` // 菜单名称
}

// QueryChildrenMenuReq 查询条件
type QueryChildrenMenuReq struct {
	ParentId uint `json:"parent_id" form:"parent_id"` // 父菜单ID
}

// AddMenuReq 添加菜单
type AddMenuReq struct {
	ParentId   *uint  `json:"parent_id" form:"parent_id"`            // 父菜单ID
	Title      string `json:"title" form:"title" binding:"required"` // 菜单名称
	Icon       string `json:"icon" form:"icon"`                      // 菜单图标
	ElSvgIcon  string `json:"el_svg_icon" form:"el_svg_icon"`        // Element 菜单图标
	MenuType   uint   `json:"menu_type" form:"menu_type"`            // 菜单类型,0:菜单,1:按钮
	OpenType   uint   `json:"open_type" form:"open_type"`            // 打开方式,0:组件,1:内链,2:外链
	Path       string `json:"path" form:"path"`                      // 路由地址/外链地址
	Name       string `json:"name" form:"name"`                      // 路由别名
	Component  string `json:"component" form:"component"`            // 组件路径
	Redirect   string `json:"redirect" form:"redirect"`              // 路由重定向
	Link       string `json:"link" form:"link"`                      // 链接地址: 内链地址
	Target     string `json:"target" form:"target"`                  // 链接地址跳转方式, _blank/_self
	Permission string `json:"permission" form:"permission"`          // 权限标识
	Hidden     uint   `json:"hidden" form:"hidden"`                  // 是否隐藏,0:显示,1:隐藏
	AlwaysShow uint   `json:"always_show" form:"always_show"`        // 始终显示根菜单,0:显示,1:隐藏
	Sort       uint   `json:"sort" form:"sort"`                      // 排序
	Note       string `json:"note" form:"note"`                      // 备注
	Status     uint   `json:"status" form:"status"`                  // 状态,0:停用,1:启用
}

// UpdateMenuReq 更新菜单
type UpdateMenuReq struct {
	ID         uint   `json:"id" form:"id" binding:"required"`       // 菜单ID
	ParentId   *uint  `json:"parent_id" form:"parent_id"`            // 父菜单ID
	Title      string `json:"title" form:"title" binding:"required"` // 菜单名称
	Icon       string `json:"icon" form:"icon"`                      // 菜单图标
	ElSvgIcon  string `json:"el_svg_icon" form:"el_svg_icon"`        // Element 菜单图标
	MenuType   uint   `json:"menu_type" form:"menu_type"`            // 菜单类型,0:菜单,1:按钮
	OpenType   uint   `json:"open_type" form:"open_type"`            // 打开方式,0:组件,1:内链,2:外链
	Path       string `json:"path" form:"path"`                      // 路由地址/外链地址
	Name       string `json:"name" form:"name"`                      // 路由别名
	Component  string `json:"component" form:"component"`            // 组件路径
	Redirect   string `json:"redirect" form:"redirect"`              // 路由重定向
	Link       string `json:"link" form:"link"`                      // 链接地址: 内链地址
	Target     string `json:"target" form:"target"`                  // 链接地址跳转方式, _blank/_self
	Permission string `json:"permission" form:"permission"`          // 权限标识
	Hidden     uint   `json:"hidden" form:"hidden"`                  // 是否隐藏,0:显示,1:隐藏
	AlwaysShow uint   `json:"always_show" form:"always_show"`        // 始终显示根菜单
	Sort       uint   `json:"sort" form:"sort"`                      // 排序
	Note       string `json:"note" form:"note"`                      // 备注
	Status     uint   `json:"status" form:"status"`                  // 状态,0:停用,1:启用
}

// ButtonPermission 按钮权限
type ButtonPermission struct {
	Permission string `json:"permission"` // 按钮权限标识
	Disabled   uint   `json:"disabled"`   // 是否禁用
}
