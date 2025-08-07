package args

import (
	"encoding/json"
	"fmt"
	"net/url"

	teetypes "github.com/masa-finance/tee-types/types"
	"github.com/masa-finance/tee-types/pkg/util"
)

type WebSearchArguments struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
	Depth    int    `json:"depth"`
	MaxDepth int    `json:"max_depth"`
}

// UnmarshalJSON implements custom JSON unmarshaling with validation
func (w *WebSearchArguments) UnmarshalJSON(data []byte) error {
	type Alias WebSearchArguments
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(w),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal Web arguments: %w", err)
	}

	return w.Validate()
}

// Validate validates the Web arguments
func (w *WebSearchArguments) Validate() error {
	if w.URL == "" {
		return fmt.Errorf("url is required")
	}

	// Validate URL format
	parsedURL, err := url.Parse(w.URL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	// Ensure URL has a scheme
	if parsedURL.Scheme == "" {
		return fmt.Errorf("URL must include a scheme (http:// or https://)")
	}

	if w.MaxDepth < 0 {
		return fmt.Errorf("max_depth must be non-negative, got: %d", w.MaxDepth)
	}

	if w.Depth < 0 {
		return fmt.Errorf("depth must be non-negative, got: %d", w.Depth)
	}

	if w.Depth > w.MaxDepth && w.MaxDepth > 0 {
		return fmt.Errorf("depth (%d) cannot exceed max_depth (%d)", w.Depth, w.MaxDepth)
	}

	return nil
}

// ValidateForJobType validates Web arguments for a specific job type
func (w *WebSearchArguments) ValidateForJobType(jobType teetypes.JobType) error {
	if err := w.Validate(); err != nil {
		return err
	}

	// Validate capability against job-specific capabilities
	return jobType.ValidateCapability(w.GetCapability())
}

// GetCapability returns the capability for web operations (always scraper)
func (w *WebSearchArguments) GetCapability() teetypes.Capability {
	return teetypes.CapScraper
}

// IsDeepScrape returns true if this is a deep scraping operation
func (w *WebSearchArguments) IsDeepScrape() bool {
	return w.MaxDepth > 1 || w.Depth > 0
}

// HasSelector returns true if a CSS selector is specified
func (w *WebSearchArguments) HasSelector() bool {
	return w.Selector != ""
}

// GetEffectiveMaxDepth returns the effective maximum depth for scraping
func (w *WebSearchArguments) GetEffectiveMaxDepth() int {
	return util.Max(w.MaxDepth, 1)
}
