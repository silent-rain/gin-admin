/*服务健康检查*/
package system

import "time"

type Health struct {
	Timestamp   time.Time `json:"timestamp"`
	Environment string    `json:"environment"`
	Host        string    `json:"host"`
	Status      string    `json:"status"`
}
