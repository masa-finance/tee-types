package types

// Key represents an encryption key
type Key struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

// KeyResponse represents a response when requesting a key
type KeyResponse struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}
