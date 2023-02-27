/*WEB 日志 DTO
 */
package logDTO

import "gin-admin/internal/dto"

// QueryWebLogReq 查询条件
type QueryWebLogReq struct {
	dto.Pagination
	Nickname  string `json:"nickname" form:"nickname"`
	OsType    uint   `json:"os_type" form:"os_type"`
	ErrorType uint   `json:"error_type" form:"error_type"`
	Level     string `json:"level" form:"level"`
	Url       string `json:"url" form:"url"`
	Msg       string `json:"msg" form:"msg"`
}

// AddWebLogReq 添加 WEB 日志
type AddWebLogReq struct {
	OsType     uint   `json:"os_type" form:"os_type"`         // 终端类型: 0: 未知,1: 安卓,2 :ios,3 :web
	ErrorType  uint   `json:"error_type" form:"error_type"`   // 错误类型: 1:接口报错,2:代码报错
	Level      string `json:"level" form:"level"`             // 日志级别
	CallerLine string `json:"caller_line" form:"caller_line"` // 日发生位置
	Url        string `json:"url" form:"url"`                 // 错误页面
	Msg        string `json:"msg" form:"msg"`                 // 日志消息
	Stack      string `json:"stack" form:"stack"`             // 堆栈信息
}
