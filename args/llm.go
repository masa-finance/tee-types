package args

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	teetypes "github.com/masa-finance/tee-types/types"
)

var (
	ErrLLMDatasetIdRequired = errors.New("dataset id is required")
	ErrLLMPromptRequired    = errors.New("prompt is required")
)

const (
	LLMDefaultMaxTokens       uint    = 300
	LLMDefaultTemperature     float64 = 0.1
	LLMDefaultMultipleColumns bool    = false
	LLMDefaultModel           string  = "gemini-1.5-flash-8b"
	LLMDefaultItems           uint    = 1
)

type LLMProcessorArguments struct {
	DatasetId   string  `json:"dataset_id"`
	Prompt      string  `json:"prompt"`
	MaxTokens   uint    `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	Items       uint    `json:"items"`
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
	if l.Temperature == 0 {
		l.Temperature = LLMDefaultTemperature
	}
	if l.MaxTokens == 0 {
		l.MaxTokens = LLMDefaultMaxTokens
	}
	if l.Items == 0 {
		l.Items = LLMDefaultItems
	}
}

func (l *LLMProcessorArguments) Validate() error {
	if l.DatasetId == "" {
		return ErrLLMDatasetIdRequired
	}
	if l.Prompt == "" {
		return ErrLLMPromptRequired
	}
	return nil
}

func (l LLMProcessorArguments) ToLLMProcessorRequest() teetypes.LLMProcessorRequest {
	return teetypes.LLMProcessorRequest{
		InputDatasetId:  l.DatasetId,
		Prompt:          l.Prompt,
		MaxTokens:       l.MaxTokens,
		Temperature:     strconv.FormatFloat(l.Temperature, 'f', -1, 64),
		MultipleColumns: LLMDefaultMultipleColumns, // overrides default in actor API
		Model:           LLMDefaultModel,           // overrides default in actor API
	}
}
