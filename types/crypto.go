package types

// EncryptedRequest represents an encrypted request
type EncryptedRequest struct {
	EncryptedResult  string `json:"encrypted_result"`
	EncryptedRequest string `json:"encrypted_request"`
}

// Key represents a cryptographic key
type Key struct {
	ID   string `json:"id"`
	Data []byte `json:"data,omitempty"`
}
