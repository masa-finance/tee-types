// Package job contains the core job types shared across TEE services
package job

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/rand"
)

// Version information
const (
	VersionMajor = 0
	VersionMinor = 1
	VersionPatch = 0
)

// GetVersion returns the semantic version string
func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}

// Arguments represents arguments passed to a job
type Arguments map[string]interface{}

// Unmarshal unmarshals job arguments into the supplied interface
func (ja Arguments) Unmarshal(i interface{}) error {
	dat, err := json.Marshal(ja)
	if err != nil {
		return err
	}
	return json.Unmarshal(dat, i)
}

// Job represents a task to be executed by a worker
type Job struct {
	Type      string    `json:"type"`
	Arguments Arguments `json:"arguments"`
	UUID      string    `json:"-"`
	Nonce     string    `json:"quote"`
	WorkerID  string    `json:"worker_id"`
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GenerateJobSignature generates a signature for the job.
// Note: This method is a placeholder. Each service will need to implement
// its own version with the appropriate sealing mechanism.
func (job *Job) GenerateJobSignature() (string, error) {
	dat, err := json.Marshal(job)
	if err != nil {
		return "", err
	}

	checksum := sha256.New()
	checksum.Write(dat)

	job.Nonce = fmt.Sprintf("%s-%s", string(checksum.Sum(nil)), randStringRunes(99))

	return job.Nonce, nil
}

// Response represents a response to a job submission
type Response struct {
	UID string `json:"uid"`
}

// Result represents the result of executing a job
type Result struct {
	Error      string `json:"error"`
	Data       []byte `json:"data"`
	Job        Job    `json:"job"`
	NextCursor string `json:"next_cursor"`
}

// Success returns true if the job was successful.
func (jr Result) Success() bool {
	return jr.Error == ""
}

// Unmarshal unmarshals the job result data.
func (jr Result) Unmarshal(i interface{}) error {
	return json.Unmarshal(jr.Data, i)
}

// Request represents a request to execute a job
type Request struct {
	EncryptedJob string `json:"encrypted_job"`
}

// Configuration represents configuration for a job
type Configuration map[string]interface{}

// Unmarshal unmarshals the job configuration into the supplied interface.
func (jc Configuration) Unmarshal(v interface{}) error {
	data, err := json.Marshal(jc)
	if err != nil {
		return fmt.Errorf("error marshalling job configuration: %w", err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("error unmarshalling job configuration: %w", err)
	}

	return nil
}

// Parameters defines the base interface for job parameters
type Parameters interface {
	// GetIdentifier returns a unique identifier for the job parameters
	GetIdentifier() string
}

// Status defines the base interface for job status
type Status interface {
	// GetError returns the job error
	GetError() string
	
	// SetError sets the job error
	SetError(err string)
	
	// GetStatus returns the job status
	GetStatus() string
	
	// SetStatus sets the job status
	SetStatus(status string)
}
