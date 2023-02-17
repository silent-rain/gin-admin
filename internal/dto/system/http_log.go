/*网络请求日志 DTO
 */
package systemDTO

import "gin-admin/internal/dto"

// QueryHttpLogReq 查询条件
type QueryHttpLogReq struct {
	dto.Pagination
	UserId     uint   `json:"user_id" form:"user_id"`
	TraceId    string `json:"trace_id" form:"trace_id"`
	ErrorCode  int    `json:"error_code" form:"error_code"`
	Method     string `json:"method" form:"method"`
	Path       string `json:"path" form:"path"`
	RemoteAddr string `json:"remote_addr" form:"remote_addr"`
	HttpType   string `json:"htpp_type" form:"htpp_type"`
}
