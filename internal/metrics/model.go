package metrics

import "time"

type Metric struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CPUUsagePercent     float64   `gorm:"column:cpu_usage_percent;not null" json:"cpu_usage_percent"`
	MemoryUsagePercent  float64   `gorm:"column:memory_usage_percent;not null" json:"memory_usage_percent"`
	NetworkUsageKbps    float64   `gorm:"column:network_usage_kbps;not null" json:"network_usage_kbps"`
	IsOutlier           bool      `gorm:"column:is_outlier;default:false" json:"is_outlier"`
	CreatedAt           time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (Metric) TableName() string {
	return "system_metrics"
}