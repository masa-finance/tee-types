package types

type Capability string

// JobType represents the type of job that can be executed
type JobType string

// Job type constants - centralized from tee-indexer and tee-worker
const (
	// Web scraping job type
	WebJob JobType = "web-scraper"

	// Telemetry job type for worker monitoring and stats
	TelemetryJob JobType = "telemetry"

	// TikTok transcription job type
	TiktokJob JobType = "tiktok-transcription"

	// Twitter job types
	TwitterJob           JobType = "twitter-scraper"            // General Twitter scraping (uses best available auth)
	TwitterCredentialJob JobType = "twitter-credential-scraper" // Twitter scraping with credentials
	TwitterApiJob        JobType = "twitter-api-scraper"        // Twitter scraping with API keys

	// Unknown/invalid job type
	UnknownJob JobType = ""
)

// String returns the string representation of the JobType
func (j JobType) String() string {
	return string(j)
}

// JobCapability represents the capabilities of a specific job type
type JobCapability struct {
	JobType      string       `json:"job_type"`
	Capabilities []Capability `json:"capabilities"`
}

// WorkerCapabilities represents all capabilities available on a worker
type WorkerCapabilities []JobCapability
