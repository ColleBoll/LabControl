package collectors

import (
	"time"

	"github.com/ColleBoll/LabControl/pkg/protocol"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func SystemMetrics() (protocol.SystemMetrics, error) {
	hostname, err := HOSTName()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}
	platform, err := HOSTPlatform()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}
	cpuType, err := CPUType()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}
	cpuPercent, err := CPUPercent()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}
	memoryUsed, err := RAMUsageMB()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}
	memoryTotal, err := RAMTotalMB()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}
	uptime, err := HOSTUptime()
	if err != nil {
		return protocol.SystemMetrics{}, err
	}

	return protocol.SystemMetrics{
		Hostname:      hostname,
		Platform:      platform,
		CPUType:       cpuType,
		CPUPercent:    cpuPercent,
		MemoryUsed:    float64(memoryUsed),
		MemoryTotal:   float64(memoryTotal),
		UptimeSeconds: uptime,
	}, nil
}

func CPUPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}

	if len(percent) == 0 {
		return 0, nil
	}

	return percent[0], nil
}

func CPUType() (string, error) {
	cpu, err := cpu.Info()
	if err != nil {
		return "not_found", err
	}

	return cpu[0].ModelName, nil
}

func RAMUsageMB() (uint64, error) {
	usage, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}

	// divide to MB usage
	return usage.Used / 1024 / 1024, nil
}

func RAMTotalMB() (uint64, error) {
	usage, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}

	// divide to MB usage
	return usage.Total / 1024 / 1024, nil
}

func HOSTName() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "not_found", err
	}

	return info.Hostname, nil
}

func HOSTPlatform() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "not_found", err
	}

	return info.Platform, nil
}

func HOSTUptime() (uint64, error) {
	info, err := host.Info()
	if err != nil {
		return 0, err
	}

	return info.Uptime, nil
}
