package models

import "time"

// AgentTelemetry represents a snapshot of metrics reported by an OpsNexus agent.
type AgentTelemetry struct {
	AgentID   string         `json:"agent_id"`
	Timestamp time.Time      `json:"timestamp"`
	Metrics   map[string]any `json:"metrics"`
}
