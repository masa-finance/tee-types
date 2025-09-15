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
	ErrLLMMaxPagesNegative  = errors.New("max pages must be non-negative")
)

const (
	LLMDefaultMaxTokens       = 300
	LLMDefaultTemperature     = "0.1"
	LLMDefaultMultipleColumns = false
	LLMDefaultModel           = "gemini-1.5-flash-8b"
	LLMDefaultMaxPages        = 1
)

type LLMProcessorArguments struct {
	DatasetId   string `json:"dataset_id"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature string `json:"temperature"`
	Items       int    `json:"items"`
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
		return fmt.Errorf("failed to unmarshal llm arguments: %w", err)
	}

	l.setDefaultValues()

	return l.Validate()
}

func (l *LLMProcessorArguments) setDefaultValues() {
	if l.MaxTokens == 0 {
		l.MaxTokens = LLMDefaultMaxTokens
	}
	if l.Temperature == "" {
		l.Temperature = LLMDefaultTemperature
	}
	if l.Items == 0 {
		l.Items = LLMDefaultMaxPages
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
	if l.Items < 1 {
		return fmt.Errorf("%w: got %v", ErrLLMMaxPagesNegative, l.Items)
	}
	return nil
}

func (l LLMProcessorArguments) ToLLMProcessorRequest() teetypes.LLMProcessorRequest {
	return teetypes.LLMProcessorRequest{
		InputDatasetId:  l.DatasetId,
		Prompt:          l.Prompt,
		MaxTokens:       l.MaxTokens,
		Temperature:     l.Temperature,
		MultipleColumns: LLMDefaultMultipleColumns, // overrides default in actor API
		Model:           LLMDefaultModel,           // overrides default in actor API
	}
}
