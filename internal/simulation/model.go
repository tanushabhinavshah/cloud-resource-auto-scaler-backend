package simulation

// SimulationConfig defines the targeted simulated metrics
type SimulationConfig struct {
	IsActive           bool    `json:"is_active"`
	CPUUsagePercent    float64 `json:"cpu_usage_percent"`
	MemoryUsagePercent float64 `json:"memory_usage_percent"`
	NetworkUsageKbps   float64 `json:"network_usage_kbps"`
}