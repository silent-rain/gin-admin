/* 菜单
 */
package systemDto

import (
	"gin-admin/internal/dto"
)

// QueryMenuReq 查询条件
type QueryMenuReq struct {
	dto.Pagination        // 分页
	Title          string `json:"title" form:"title"` // 菜单名称
}

// AddMenuReq 添加菜单
type AddMenuReq struct {
	ParentId   *uint  `json:"parent_id" form:"parent_id"`            // 父菜单ID
	Title      string `json:"title" form:"title" binding:"required"` // 菜单名称
	Icon       string `json:"icon" form:"icon"`                      // 菜单图标
	MenuType   uint   `json:"menu_type" form:"menu_type"`            // 菜单类型,0:菜单,1:按钮
	OpenType   uint   `json:"open_type" form:"open_type"`            // 打开方式,0:组件,1:内链,2:外链
	Path       string `json:"path" form:"path"`                      // 路由地址
	Component  string `json:"component" form:"component"`            // 组件路径
	Link       string `json:"link" form:"link"`                      // 链接地址: 内链地址/外链地址
	Target     string `json:"target" form:"target"`                  // 链接地址跳转方式, _blank/_self
	Permission string `json:"permission" form:"permission"`          // 权限标识
	Hide       uint   `json:"hide" form:"hide"`                      // 是否隐藏,0:显示,1:隐藏
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
	MenuType   uint   `json:"menu_type" form:"menu_type"`            // 菜单类型,0:菜单,1:按钮
	OpenType   uint   `json:"open_type" form:"open_type"`            // 打开方式,0:组件,1:内链,2:外链
	Path       string `json:"path" form:"path"`                      // 路由地址
	Component  string `json:"component" form:"component"`            // 组件路径
	Link       string `json:"link" form:"link"`                      // 链接地址: 内链地址/外链地址
	Target     string `json:"target" form:"target"`                  // 链接地址跳转方式, _blank/_self
	Permission string `json:"permission" form:"permission"`          // 权限标识
	Hide       uint   `json:"hide" form:"hide"`                      // 是否隐藏,0:显示,1:隐藏
	Sort       uint   `json:"sort" form:"sort"`                      // 排序
	Note       string `json:"note" form:"note"`                      // 备注
	Status     uint   `json:"status" form:"status"`                  // 状态,0:停用,1:启用
}
