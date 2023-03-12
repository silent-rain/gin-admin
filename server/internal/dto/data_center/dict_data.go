/*字典数据管理*/
package datacenter

import "gin-admin/internal/dto"

// QueryDictDataReq 查询条件
type QueryDictDataReq struct {
	dto.Pagination        // 分页
	DictId         uint  `json:"dict_id" form:"dict_id"` // 字典维度ID
	Name           string `json:"name" form:"name"`       // 字典项名称
	Value          string `json:"value" form:"value"`     // 字典项值
}

// AddDictDataReq 添加字典数据
type AddDictDataReq struct {
	DictId uint   `json:"dict_id" form:"dict_id" binding:"required"` // 字典维度ID
	Name   string `json:"name" form:"name"`                          // 字典项名称
	Value  string `json:"value" form:"value"`                        // 字典项值
	Note   string `json:"note" form:"note"`                          // 备注
	Status uint   `json:"status" form:"status"`                      // 状态,0:停用,1:启用
}

// UpdateDictDataReq 更新字典数据
type UpdateDictDataReq struct {
	ID     uint   `json:"id" form:"id" binding:"required"` // 字典项ID
	DictId uint   `json:"dict_id" form:"dict_id"`          // 字典维度ID
	Name   string `json:"name" form:"name"`                // 字典项名称
	Value  string `json:"value" form:"value"`              // 字典项值
	Note   string `json:"note" form:"note"`                // 备注
	Status uint   `json:"status" form:"status"`            // 状态,0:停用,1:启用
}
