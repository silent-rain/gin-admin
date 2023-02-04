/* 菜单
 */
package systemDto

import (
	"time"

	"gin-admin/internal/dto"
)

// QueryMenuReq 查询条件
type QueryMenuReq struct {
	dto.Pagination        // 分页
	Title          string `json:"title" form:"title"` // 菜单名称
}

// AddMenuReq 添加菜单
type AddMenuReq struct {
	ParentId     *uint     `json:"parent_id" form:"parent_id"`            // 父菜单ID
	Title        string    `json:"title" form:"title" binding:"required"` // 菜单名称
	Icon         string    `json:"icon" form:"icon"`                      // 菜单图标
	OpenType     string    `json:"open_type" form:"open_type"`            // 打开方式,0:组件,1:内链,2:外联
	Path         string    `json:"path" form:"path"`                      // 路由地址/外链地址
	Component    string    `json:"component" form:"component"`            // 组件路径/内链地址
	Target       string    `json:"target" form:"target"`                  // 链接地址跳转方式, _blank/_self
	Permission   string    `json:"permission" form:"permission"`          // 权限标识
	MenuType     string    `json:"menu_type" form:"menu_type"`            // 菜单类型,0:菜单,1:按钮
	Hide         string    `json:"hide" form:"hide"`                      // 菜单是否隐藏,0:隐藏,1:显示
	Sort         uint      `json:"sort" form:"sort"`                      // 排序
	Note         string    `json:"note" form:"note"`                      // 备注
	Status       uint      `json:"status" form:"status"`                  // 状态,0:停用,1:启用
	CreateUserId uint      `json:"create_user_id" form:"create_user_id"`  // 创建菜单用户ID
	UpdateUserId uint      `json:"update_user_id" form:"update_user_id"`  // 更新菜单用户ID
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`          // 更新时间
}

// UpdateMenuReq 更新菜单
type UpdateMenuReq struct {
	ID           uint      `json:"id" form:"id" binding:"required"`       // 菜单ID
	ParentId     *uint     `json:"parent_id" form:"parent_id"`            // 父菜单ID
	Title        string    `json:"title" form:"title" binding:"required"` // 菜单名称
	Icon         string    `json:"icon" form:"icon"`                      // 菜单图标
	OpenType     string    `json:"open_type" form:"open_type"`            // 打开方式,0:组件,1:内链,2:外联
	Path         string    `json:"path" form:"path"`                      // 路由地址/外链地址
	Component    string    `json:"component" form:"component"`            // 组件路径/内链地址
	Target       string    `json:"target" form:"target"`                  // 链接地址跳转方式, _blank/_self
	Permission   string    `json:"permission" form:"permission"`          // 权限标识
	MenuType     string    `json:"menu_type" form:"menu_type"`            // 菜单类型,0:菜单,1:按钮
	Hide         string    `json:"hide" form:"hide"`                      // 菜单是否隐藏,0:隐藏,1:显示
	Sort         uint      `json:"sort" form:"sort"`                      // 排序
	Note         string    `json:"note" form:"note"`                      // 备注
	Status       uint      `json:"status" form:"status"`                  // 状态,0:停用,1:启用
	CreateUserId uint      `json:"create_user_id" form:"create_user_id"`  // 创建菜单用户ID
	UpdateUserId uint      `json:"update_user_id" form:"update_user_id"`  // 更新菜单用户ID
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`          // 更新时间
}
