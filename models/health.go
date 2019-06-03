package models

// Health response used by the health check
type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// NewHealth returns a Health object
func NewHealth(status string, version string) Health {
	return Health{
		Status:  status,
		Version: version,
	}
}
