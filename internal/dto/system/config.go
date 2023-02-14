/*应用配置表*/
package systemDTO

import "gin-admin/internal/dto"

// QueryConfigReq 查询条件
type QueryConfigReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"` // 配置名称
	Key            string `json:"key" form:"key"`   // 配置参数(英文)
}

// AddConfigReq 添加配置
type AddConfigReq struct {
	ID       uint   `json:"id" form:"id"`
	ParentId *uint  `json:"parent_id" form:"parent_id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Key      string `json:"key" form:"key" binding:"required"`
	Value    string `json:"value" form:"value"`
	Sort     uint   `json:"sort" form:"sort"`
	Note     string `json:"note" form:"note"`
	Status   uint   `json:"status" form:"status"`
}

// UpdateConfigReq 更新配置
type UpdateConfigReq struct {
	ID       uint   `json:"id" form:"id" binding:"required"`
	ParentId *uint  `json:"parent_id" form:"parent_id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Key      string `json:"key" form:"key" binding:"required"`
	Value    string `json:"value" form:"value"`
	Sort     uint   `json:"sort" form:"sort"`
	Note     string `json:"note" form:"note"`
	Status   uint   `json:"status" form:"status"`
}
