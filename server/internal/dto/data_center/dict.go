/*字典维度管理*/
package datacenter

import (
	"github.com/silent-rain/gin-admin/internal/dto"
)

// QueryDictReq 查询条件
type QueryDictReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"` // 字典名称
	Code           string `json:"code" form:"code"` // 字典编码
}

// AddDictReq 添加字典维度信息
type AddDictReq struct {
	Name   string `json:"name" form:"name" binding:"required"` // 字典名称
	Code   string `json:"code" form:"code" binding:"required"` // 字典编码
	Note   string `json:"note" form:"note"`                    // 备注
	Status uint   `json:"status" form:"status"`                // 状态,0:停用,1:启用
}

// UpdateDictReq 更新字典维度信息
type UpdateDictReq struct {
	ID     uint   `json:"id" form:"id" binding:"required"`     // 字典ID
	Name   string `json:"name" form:"name" binding:"required"` // 字典名称
	Code   string `json:"code" form:"code" binding:"required"` // 字典编码
	Note   string `json:"note" form:"note"`                    // 备注
	Status uint   `json:"status" form:"status"`                // 状态,0:停用,1:启用
}
