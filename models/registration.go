package models

// AgentRegistration represents the identity information
// sent by an OpsNexus agent when it registers with the backend.
type AgentRegistration struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Version  string `json:"version"`
}
