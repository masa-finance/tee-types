// Package types contains the minimal shared type definitions for tee-worker and tee-indexer
package types

// No imports needed

// JobArguments represents arguments passed to a job
type JobArguments map[string]interface{}

// Job represents a task to be executed by a worker
type Job struct {
	Type      string       `json:"type"`
	Arguments JobArguments `json:"arguments"`
	UUID      string       `json:"-"`
	Nonce     string       `json:"quote"`
	WorkerID  string       `json:"worker_id"`
}

// JobResult represents the result of executing a job
type JobResult struct {
	Error      string `json:"error"`
	Data       []byte `json:"data"`
	Job        Job    `json:"job"`
	NextCursor string `json:"next_cursor"`
}

// JobRequest represents a request to execute a job
type JobRequest struct {
	EncryptedJob string `json:"encrypted_job"`
}

// Common job type constants
const (
	// WebScraperType represents the job type for web scraping
	WebScraperType = "webscraper"
	
	// TwitterScraperType represents standard Twitter scraping jobs
	TwitterScraperType = "twitter"
)
