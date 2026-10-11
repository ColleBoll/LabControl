package protocol

import "time"

type AgentHandshake struct {
	ID        int       `json:"agent_id"`
	Hostname  string    `json:"hostname"`
	Timestamp time.Time `json:"timestamp"`
}

type AgentSession struct {
}
