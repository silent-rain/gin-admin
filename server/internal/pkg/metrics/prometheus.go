package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

const (
	namespace = "metrics"
	subsystem = "gin_admin"
)

// 指标注册
func init() {
	prometheus.MustRegister(metricsRequestsTotal, metricsRequestsCost)
}

// metricsRequestsTotal metrics for request total 计数器（Counter）
var metricsRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "request(ms) total",
	},
	[]string{"method", "path"},
)

// metricsRequestsCost metrics for requests cost 累积直方图（Histogram）
var metricsRequestsCost = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost",
		Help:      "request(ms) cost milliseconds",
	},
	[]string{"method", "path", "success", "http_code", "business_code", "cost_milliseconds", "trace_id"},
)

// RecordMetrics 记录指标
func RecordMetrics(method, path string, httpStatus int, errorCode uint, costSeconds float64, traceId string) {
	metricsRequestsTotal.With(prometheus.Labels{
		"method": method,
		"path":   path,
	}).Inc()

	metricsRequestsCost.With(prometheus.Labels{
		"method":            method,
		"path":              path,
		"http_status":       cast.ToString(httpStatus),
		"error_code":        cast.ToString(errorCode),
		"cost_milliseconds": cast.ToString(costSeconds * 1000),
		"trace_id":          traceId,
	}).Observe(costSeconds)
}
