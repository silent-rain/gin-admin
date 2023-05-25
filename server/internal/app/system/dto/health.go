// Package dto 服务健康检查
package dto

// Health 健康检查
type Health struct {
	Timestamp   string `json:"timestamp"`
	Environment string `json:"environment"`
	Host        string `json:"host"`
	Status      string `json:"status"`
}
