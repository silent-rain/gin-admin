/*
 * @Author: silent-rain
 * @Date: 2023-01-13 00:20:26
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-13 23:51:47
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dto/system/role.go
 * @Descripttion: 角色
 */
package systemDto

import "gin-admin/internal/dto"

// RoleQueryReq 查询条件
type RoleQueryReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"` // 角色名称
}

// RoleAddReq 添加角色
type RoleAddReq struct {
	Name   string `json:"name" form:"name" binding:"required"` // 角色名称
	Status uint   `json:"status" form:"status"`                // 角色状态,0:停用,1:启用
	Sort   uint   `json:"sort" form:"sort"`                    // 排序
	Note   string `json:"note" form:"note"`                    // 备注
}

// RoleUpdateReq 更新角色
type RoleUpdateReq struct {
	ID     uint   `json:"id" form:"id" binding:"required"`     // 角色ID
	Name   string `json:"name" form:"name" binding:"required"` // 角色名称
	Status uint   `json:"status" form:"status"`                // 角色状态,0:停用,1:启用
	Sort   uint   `json:"sort"  form:"sort"`                   // 排序
	Note   string `json:"note"  form:"note"`                   // 备注
}

// RoleDeleteReq 删除角色
type RoleDeleteReq struct {
	ID uint `json:"id" form:"id" binding:"required"` // 角色ID
}

// 更新角色状态
type RoleStatusReq struct {
	ID     uint `form:"id" binding:"required"`
	Status uint `form:"status"`
}
