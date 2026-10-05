package protocol

import "time"

type AgentReport struct {
	AgentID   string         `json:"agent_id"`
	Timestamp time.Time      `json:"timestamp"`
	System    SystemMetrics  `json:"system"`
	Storage   StorageMetrics `json:"storage"`
	Docker    DockerMetrics  `json:"docker"`
}

type SystemMetrics struct {
	Hostname      string  `json:"hostname"`
	Platform      string  `json:"platform"`
	CPUType       string  `json:"cpu_type"`
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryUsed    float64 `json:"memory_used"`
	MemoryTotal   float64 `json:"memory_total"`
	UptimeSeconds uint64  `json:"uptime_seconds"`
}

type StorageMetrics struct {
	UsedPercent float64 `json:"used_percent"`
}

type DockerMetrics struct {
	Running int `json:"running"`
	Stopped int `json:"stopped"`
}
