/*Http协议接口管理表*/
package apiauth

import "gin-admin/internal/dto"

// QueryApiHttpReq 查询条件
type QueryApiHttpReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"`     // 接口名称
	Method         string `json:"method" form:"method"` // 请求类型:GET,POST,PUT,DELETE
	Uri            string `json:"uri" form:"uri"`       // URI资源
	Status         *uint  `json:"status" form:"status"` // 状态,0:停用,1:启用
}

// AddApiHttpReq 添加角色
type AddApiHttpReq struct {
	Name   string `json:"name" form:"name"`     // 接口名称
	Method string `json:"method" form:"method"` // 请求类型:GET,POST,PUT,DELETE
	Uri    string `json:"uri" form:"uri"`       // URI资源
	Note   string `json:"note" form:"note"`     // 备注
	Status uint   `json:"status" form:"status"` // 状态,0:停用,1:启用
}

// UpdateApiHttpReq 更新角色
type UpdateApiHttpReq struct {
	ID     uint   `json:"id" form:"id" binding:"required"` // 自增ID
	Name   string `json:"name" form:"name"`                // 接口名称
	Method string `json:"method" form:"method"`            // 请求类型:GET,POST,PUT,DELETE
	Uri    string `json:"uri" form:"uri"`                  // URI资源
	Note   string `json:"note" form:"note"`                // 备注
	Status uint   `json:"status" form:"status"`            // 状态,0:停用,1:启用
}
