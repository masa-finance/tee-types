package job

// TelemetryJobType represents the job type for telemetry operations
const TelemetryJobType = "telemetry"

// TelemetryConfiguration defines configuration for telemetry jobs
type TelemetryConfiguration struct {
	StatsInterval int `json:"stats_interval"`
}

// TelemetryResult represents the result of a telemetry operation
type TelemetryResult struct {
	Stats map[string]uint `json:"stats"`
}
