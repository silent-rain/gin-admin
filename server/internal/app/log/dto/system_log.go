/*系统日志 DTO
 */
package dto

import (
	DTO "github.com/silent-rain/gin-admin/internal/dto"
)

// QuerySystemLogReq 查询条件
type QuerySystemLogReq struct {
	DTO.Pagination
	UserId    uint   `json:"user_id" form:"user_id"`
	TraceId   string `json:"trace_id" form:"trace_id"`
	Level     string `json:"level" form:"level"`
	ErrorCode uint   `json:"error_code" form:"error_code"`
	ErrorMsg  string `json:"error_msg" form:"error_msg"`
	Msg       string `json:"msg" form:"msg"`
}
