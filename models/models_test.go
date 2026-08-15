package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAgentJSONContract(t *testing.T) {
	createdAt := time.Date(2026, time.January, 2, 3, 4, 5, 0, time.UTC)
	lastSeen := time.Date(2026, time.January, 2, 3, 4, 6, 0, time.UTC)
	agent := Agent{
		ID: "agent-1", Name: "edge-1", Hostname: "host-1", OS: "linux",
		Arch: "amd64", Version: "1.2.3", Status: "online",
		LastSeen: lastSeen, CreatedAt: createdAt,
	}

	assertJSON(t, agent, `{"id":"agent-1","name":"edge-1","hostname":"host-1","os":"linux","arch":"amd64","version":"1.2.3","status":"online","last_seen":"2026-01-02T03:04:06Z","created_at":"2026-01-02T03:04:05Z"}`)
}

func TestAgentRegistrationJSONContract(t *testing.T) {
	registration := AgentRegistration{
		ID: "agent-1", Name: "edge-1", Hostname: "host-1", OS: "linux",
		Arch: "amd64", Version: "1.2.3",
	}

	assertJSON(t, registration, `{"id":"agent-1","name":"edge-1","hostname":"host-1","os":"linux","arch":"amd64","version":"1.2.3"}`)
}

func TestAgentTelemetryJSONContract(t *testing.T) {
	telemetry := AgentTelemetry{
		AgentID:   "agent-1",
		Timestamp: time.Date(2026, time.January, 2, 3, 4, 5, 0, time.UTC),
		Metrics:   map[string]any{"cpu_percent": 42.5, "healthy": true},
	}

	assertJSON(t, telemetry, `{"agent_id":"agent-1","timestamp":"2026-01-02T03:04:05Z","metrics":{"cpu_percent":42.5,"healthy":true}}`)
}

func assertJSON(t *testing.T, value any, want string) {
	t.Helper()
	got, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal JSON: %v", err)
	}
	if string(got) != want {
		t.Fatalf("JSON = %s, want %s", got, want)
	}
}
