// Package crypto contains cryptographic interfaces and helpers for secure operations
package crypto

// KeyType defines the type of a key in the keyring
type KeyType string

const (
	// SealingKeyType is used for sealing operations
	SealingKeyType KeyType = "sealing"
	
	// EncryptionKeyType is used for encryption operations 
	EncryptionKeyType KeyType = "encryption"
	
	// SigningKeyType is used for signing operations
	SigningKeyType KeyType = "signing"
)

// Key represents a cryptographic key in the keyring
type Key struct {
	ID   string  `json:"id"`
	Type KeyType `json:"type"`
	Data []byte  `json:"data"`
}

// KeyRing defines the interface for key management
// This follows the modern approach of using a KeyRing instead of single keys
type KeyRing interface {
	// AddKey adds a key to the keyring
	AddKey(key Key) error
	
	// GetKey retrieves a key from the keyring by ID
	GetKey(id string) (Key, error)
	
	// GetKeysByType retrieves all keys of a specific type
	GetKeysByType(keyType KeyType) ([]Key, error)
	
	// RemoveKey removes a key from the keyring by ID
	RemoveKey(id string) error
	
	// Seal seals data using the keyring
	Seal(data []byte) ([]byte, error)
	
	// Unseal unseals data using the keyring
	Unseal(data []byte) ([]byte, error)
	
	// SealWithKey seals data using a specific key from the keyring
	SealWithKey(keyID string, data []byte) ([]byte, error)
	
	// UnsealWithKey unseals data using a specific key from the keyring
	UnsealWithKey(keyID string, data []byte) ([]byte, error)
}

// TEEOptions defines options for TEE operations
type TEEOptions struct {
	// IsTestEnvironment indicates if this is a test environment
	IsTestEnvironment bool
	
	// KeysDirectory specifies the directory for storing keys
	KeysDirectory string
}
