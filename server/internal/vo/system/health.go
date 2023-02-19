/*服务健康检查*/
package system

type Health struct {
	Timestamp   string `json:"timestamp"`
	Environment string `json:"environment"`
	Host        string `json:"host"`
	Status      string `json:"status"`
}
