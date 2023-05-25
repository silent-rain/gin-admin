/*指标信息*/
package model

import "encoding/json"

// MetricsMessage 指标信息
type MetricsMessage struct {
	ProjectName string  `json:"project_name"` // 项目名，用于区分不同项目告警信息
	Env         string  `json:"env"`          // 运行环境
	TraceID     string  `json:"trace_id"`     // 唯一ID，用于追踪关联
	HOST        string  `json:"host"`         // 请求 HOST
	Path        string  `json:"path"`         // 请求 Path
	Method      string  `json:"method"`       // 请求 Method
	HttpStatus  int     `json:"http_status"`  // HTTP 状态码
	ErrorCode   uint    `json:"error_code"`   // 业务错误码
	CostSeconds float64 `json:"cost_seconds"` // 耗时，单位：秒
}

// Marshal 序列化到JSON
func (m *MetricsMessage) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(m)
	return
}
