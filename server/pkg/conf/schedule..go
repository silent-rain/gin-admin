// Package conf 任务调度配置
package conf

// ScheduleConfig 任务调度配置
type ScheduleConfig struct {
	Ticker map[string]bool `toml:"ticker"` // 即时器
	Timer  map[string]bool `toml:"timer"`  // 定时器
}

// IsEnableTicker 是否启用即时器
func (t *ScheduleConfig) IsEnableTicker(taskName string) bool {
	falg, ok := t.Ticker[taskName]
	if !ok {
		return false
	}
	return falg
}

// IsEnableTicker 是否启用定时器
func (t *ScheduleConfig) IsEnableTimer(taskName string) bool {
	falg, ok := t.Timer[taskName]
	if !ok {
		return false
	}
	return falg
}
