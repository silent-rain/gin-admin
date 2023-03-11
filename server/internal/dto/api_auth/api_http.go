/*Http协议接口管理表*/
package apiauth

import (
	"encoding/json"
	"gin-admin/internal/dto"
)

// QueryApiHttpReq 查询条件
type QueryApiHttpReq struct {
	dto.Pagination        // 分页
	Name           string `json:"name" form:"name"`     // 接口名称
	Method         string `json:"method" form:"method"` // 请求类型:GET,POST,PUT,DELETE
	Uri            string `json:"uri" form:"uri"`       // URI资源
	Status         *uint  `json:"status" form:"status"` // 状态,0:停用,1:启用
}

// AddApiHttpReq 添加接口
type AddApiHttpReq struct {
	ParentId *uint  `json:"parent_id" form:"parent_id"` // 父菜单ID
	Name     string `json:"name" form:"name"`           // 接口名称
	Method   string `json:"method" form:"method"`       // 请求类型:GET,POST,PUT,DELETE
	Uri      string `json:"uri" form:"uri"`             // URI资源
	Note     string `json:"note" form:"note"`           // 备注
	Status   uint   `json:"status" form:"status"`       // 状态,0:停用,1:启用
}

// UpdateApiHttpReq 更新接口
type UpdateApiHttpReq struct {
	ParentId *uint  `json:"parent_id" form:"parent_id"`      // 父菜单ID
	ID       uint   `json:"id" form:"id" binding:"required"` // 自增ID
	Name     string `json:"name" form:"name"`                // 接口名称
	Method   string `json:"method" form:"method"`            // 请求类型:GET,POST,PUT,DELETE
	Uri      string `json:"uri" form:"uri"`                  // URI资源
	Note     string `json:"note" form:"note"`                // 备注
	Status   uint   `json:"status" form:"status"`            // 状态,0:停用,1:启用
}

// ApiHttpUserCache API Token 请求的存储用户结构
type ApiHttpUserCache struct {
	UserId   uint   `json:"user_id"`  // 用户ID
	Nickname string `json:"nickname"` // 用户昵称
}

// String 转为字符串
func (c *ApiHttpUserCache) String() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Unmarshal 字符串解码为结构体
func (c *ApiHttpUserCache) Unmarshal(value string) error {
	if err := json.Unmarshal([]byte(value), &c); err != nil {
		return err
	}
	return nil
}
