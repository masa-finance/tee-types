package types

// EncryptedRequest represents an encrypted request and result
type EncryptedRequest struct {
	EncryptedResult  string `json:"encrypted_result"`
	EncryptedRequest string `json:"encrypted_request"`
}

// JobError represents an error response for a job
type JobError struct {
	Error string `json:"error"`
}
