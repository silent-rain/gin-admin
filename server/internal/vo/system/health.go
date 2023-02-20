/*服务健康检查*/
package system

// Health 健康检查
type Health struct {
	Timestamp   string `json:"timestamp"`
	Environment string `json:"environment"`
	Host        string `json:"host"`
	Status      string `json:"status"`
}
