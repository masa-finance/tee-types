package args

import (
	"encoding/json"
	"errors"
	"fmt"

	teetypes "github.com/masa-finance/tee-types/types"
)

var (
	ErrLLMDatasetIdRequired = errors.New("dataset id is required")
	ErrLLMPromptRequired    = errors.New("prompt is required")
	ErrLLMMaxTokensNegative = errors.New("max tokens must be non-negative")
)

const (
	llmDefaultMaxTokens       = 300
	llmDefaultTemperature     = "0.1"
	llmDefaultMultipleColumns = false
	llmDefaultModel           = "gemini-1.5-flash-8b"
)

type LLMProcessorArguments struct {
	QueryType   string `json:"type"`
	DatasetId   string `json:"dataset_id"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature string `json:"temperature"`
}

// UnmarshalJSON implements custom JSON unmarshaling with validation
func (l *LLMProcessorArguments) UnmarshalJSON(data []byte) error {
	// Prevent infinite recursion (you call json.Unmarshal which then calls `UnmarshalJSON`, which then calls `json.Unmarshal`...)
	type Alias LLMProcessorArguments
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(l),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal web arguments: %w", err)
	}

	l.setDefaultValues()

	return l.Validate()
}

func (l *LLMProcessorArguments) setDefaultValues() {
	if l.MaxTokens == 0 {
		l.MaxTokens = llmDefaultMaxTokens
	}
	if l.Temperature == "" {
		l.Temperature = llmDefaultTemperature
	}
}

func (l *LLMProcessorArguments) Validate() error {
	if l.DatasetId == "" {
		return ErrLLMDatasetIdRequired
	}
	if l.Prompt == "" {
		return ErrLLMPromptRequired
	}
	if l.MaxTokens < 0 {
		return fmt.Errorf("%w: got %v", ErrLLMMaxTokensNegative, l.MaxTokens)
	}
	return nil
}

func (l *LLMProcessorArguments) ValidateForJobType(jobType teetypes.JobType) error {
	if err := l.Validate(); err != nil {
		return err
	}

	// Validate QueryType against job-specific capabilities
	return jobType.ValidateCapability(l.GetCapability())
}

// GetCapability returns the capability for web operations (always scraper)
func (l *LLMProcessorArguments) GetCapability() teetypes.Capability {
	return teetypes.CapDatasetProcessor
}

func (l LLMProcessorArguments) ToLLMProcessorRequest() teetypes.LLMProcessorRequest {
	return teetypes.LLMProcessorRequest{
		InputDatasetId:  l.DatasetId,
		Prompt:          l.Prompt,
		MaxTokens:       l.MaxTokens,
		Temperature:     l.Temperature,
		MultipleColumns: llmDefaultMultipleColumns,
		Model:           llmDefaultModel,
	}
}
