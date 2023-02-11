/*通用数据传输
 */
package DTO

// DeleteReq 根据 ID，删除数据请求
type DeleteReq struct {
	ID uint `json:"id" form:"id" binding:"required"` // 数据 ID
}

// BatchDeleteReq 根据 ID 列表，批量删除数据请求
type BatchDeleteReq struct {
	Ids []uint `json:"ids" form:"ids" binding:"required"` // 数据 ID 列表
}

// UpdateStatusReq 根据 ID，更新数据状态
type UpdateStatusReq struct {
	ID     uint `json:"id" form:"id" binding:"required"` // 数据 ID
	Status uint `json:"status" form:"status"`            // 状态
}
