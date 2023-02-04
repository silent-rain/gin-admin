/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:20:26
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 13:37:04
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/role.go
 * @Descripttion: 角色
 */
package systemDto

import "gin-admin/internal/dto"

// QueryRoleReq 查询条件
type QueryRoleReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"` // 角色名称
}

// AddRoleReq 添加角色
type AddRoleReq struct {
	Name   string `json:"name" form:"name" binding:"required"` // 角色名称
	Status uint   `json:"status" form:"status"`                // 角色状态,0:停用,1:启用
	Sort   uint   `json:"sort" form:"sort"`                    // 排序
	Note   string `json:"note" form:"note"`                    // 备注
}

// UpdateRoleReq 更新角色
type UpdateRoleReq struct {
	ID     uint   `json:"id" form:"id" binding:"required"`     // 角色ID
	Name   string `json:"name" form:"name" binding:"required"` // 角色名称
	Status uint   `json:"status" form:"status"`                // 角色状态,0:停用,1:启用
	Sort   uint   `json:"sort"  form:"sort"`                   // 排序
	Note   string `json:"note"  form:"note"`                   // 备注
}
