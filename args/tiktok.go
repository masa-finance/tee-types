package args

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	teetypes "github.com/masa-finance/tee-types/types"
)

// TikTokTranscriptionArguments defines args for TikTok transcriptions
type TikTokTranscriptionArguments struct {
	VideoURL string `json:"video_url"`
	Language string `json:"language,omitempty"` // e.g., "eng-US"
}

// UnmarshalJSON implements custom JSON unmarshaling with validation
func (t *TikTokTranscriptionArguments) UnmarshalJSON(data []byte) error {
	type Alias TikTokTranscriptionArguments
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal TikTok arguments: %w", err)
	}

	// Normalize language to lowercase if provided
	if t.Language != "" {
		t.Language = strings.ToLower(t.Language)
	}

	return t.Validate()
}

// Validate validates the TikTok arguments
func (t *TikTokTranscriptionArguments) Validate() error {
	if t.VideoURL == "" {
		return fmt.Errorf("video_url is required")
	}

	// Validate URL format
	parsedURL, err := url.Parse(t.VideoURL)
	if err != nil {
		return fmt.Errorf("invalid video_url format: %w", err)
	}

	// Basic TikTok URL validation
	if !t.IsTikTokURL(parsedURL) {
		return fmt.Errorf("URL must be a valid TikTok video URL")
	}

	// Validate language format if provided
	if t.Language != "" {
		if err := t.validateLanguageCode(); err != nil {
			return err
		}
	}

	return nil
}

// GetCapability returns the capability for TikTok operations (always transcription)
func (t *TikTokTranscriptionArguments) GetCapability() teetypes.Capability {
	return teetypes.CapTiktokTranscription
}

// IsTikTokURL validates if the URL is a TikTok URL
func (t *TikTokTranscriptionArguments) IsTikTokURL(parsedURL *url.URL) bool {
	host := strings.ToLower(parsedURL.Host)
	return host == "tiktok.com" ||
		host == "www.tiktok.com" ||
		host == "vm.tiktok.com" ||
		strings.HasSuffix(host, ".tiktok.com")
}

// HasLanguagePreference returns true if a language preference is specified
func (t *TikTokTranscriptionArguments) HasLanguagePreference() bool {
	return t.Language != ""
}

// GetLanguageCode returns the language code, defaulting to "en-us" if not specified
func (t *TikTokTranscriptionArguments) GetLanguageCode() string {
	if t.Language == "" {
		return "en-us"
	}
	return t.Language
}

// validateLanguageCode validates the language code format
func (t *TikTokTranscriptionArguments) validateLanguageCode() error {
	// Basic validation for language codes like "en-us", "es-es", etc.
	parts := strings.Split(t.Language, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid language format '%s', expected format: 'lang-region' (e.g., 'en-us')", t.Language)
	}

	if len(parts[0]) != 2 || len(parts[1]) != 2 {
		return fmt.Errorf("invalid language format '%s', expected 2-letter language and region codes", t.Language)
	}

	return nil
}
