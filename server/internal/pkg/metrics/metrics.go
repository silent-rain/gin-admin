/*指标*/
package metrics

import (
	systemModel "gin-admin/internal/model/system"
)

// RecordHandler 指标处理
func RecordHandler(msg systemModel.MetricsMessage) {
	RecordMetrics(
		msg.Method,
		msg.Path,
		msg.HttpStatus,
		msg.ErrorCode,
		msg.CostSeconds,
		msg.TraceID,
	)
}
