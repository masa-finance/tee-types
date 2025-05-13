// Package types defines common type definitions used across tee services
package types

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/rand"
)

// JobArguments represents arguments passed to a job
type JobArguments map[string]interface{}

// Unmarshal unmarshals job arguments into the supplied interface
func (ja JobArguments) Unmarshal(i interface{}) error {
	dat, err := json.Marshal(ja)
	if err != nil {
		return err
	}
	return json.Unmarshal(dat, i)
}

// Job represents a task to be executed by a worker
type Job struct {
	Type      string       `json:"type"`
	Arguments JobArguments `json:"arguments"`
	UUID      string       `json:"-"`
	Nonce     string       `json:"quote"`
	WorkerID  string       `json:"worker_id"`
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
// Note: This method will need to be adjusted when used in actual implementations
// to use the appropriate sealing mechanism.
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

// JobResponse represents a response to a job submission
type JobResponse struct {
	UID string `json:"uid"`
}

// JobResult represents the result of executing a job
type JobResult struct {
	Error      string `json:"error"`
	Data       []byte `json:"data"`
	Job        Job    `json:"job"`
	NextCursor string `json:"next_cursor"`
}

// Success returns true if the job was successful.
func (jr JobResult) Success() bool {
	return jr.Error == ""
}

// Unmarshal unmarshals the job result data.
func (jr JobResult) Unmarshal(i interface{}) error {
	return json.Unmarshal(jr.Data, i)
}

// JobRequest represents a request to execute a job
type JobRequest struct {
	EncryptedJob string `json:"encrypted_job"`
}

// JobConfiguration represents configuration for a job
type JobConfiguration map[string]interface{}

// Unmarshal unmarshals the job configuration into the supplied interface.
func (jc JobConfiguration) Unmarshal(v interface{}) error {
	data, err := json.Marshal(jc)
	if err != nil {
		return fmt.Errorf("error marshalling job configuration: %w", err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("error unmarshalling job configuration: %w", err)
	}

	return nil
}
