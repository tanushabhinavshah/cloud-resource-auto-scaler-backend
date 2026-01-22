package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func Collect(prevBytes uint64) (*Metric, uint64, error) {
	cpuPercents, _ := cpu.Percent(0, false)
	var cpuVal float64
	if len(cpuPercents) > 0 {
		cpuVal = cpuPercents[0]
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, 0, err
	}
	// NEW CALCULATION: Matches modern Linux system monitors (like GNOME/KDE)
	// Usage % = (1 - (Available / Total)) * 100
	memoryUsage := (1.0 - (float64(vm.Available) / float64(vm.Total))) * 100

	netIO, err := net.IOCounters(false)
	if err != nil {
		return nil, 0, err
	}
	currentTotalBytes := netIO[0].BytesRecv + netIO[0].BytesSent
	
	var networkKbps float64
	if prevBytes > 0 && currentTotalBytes >= prevBytes {
		networkKbps = float64(currentTotalBytes-prevBytes) / 1024
	}

	return &Metric{
		CPUUsagePercent:    cpuVal,
		MemoryUsagePercent: memoryUsage,
		NetworkUsageKbps:   networkKbps,
		IsOutlier:          false,
		CreatedAt:          time.Now(),
	}, currentTotalBytes, nil
}