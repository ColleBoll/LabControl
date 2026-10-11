package protocol

import "time"

type AgentHandshake struct {
	AgentID       string    `json:"agent_id"`
	AgentHostname string    `json:"hostname"`
	Timestamp     time.Time `json:"timestamp"`
}

type AgentSession struct {
}
