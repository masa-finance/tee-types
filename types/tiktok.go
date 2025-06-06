// Package types provides shared types between tee-worker and tee-indexer
package types

// TikTokTranscriptionResult defines the structure of the result data for a TikTok transcription
type TikTokTranscriptionResult struct {
	TranscriptionText string `json:"transcription_text"`
	DetectedLanguage  string `json:"detected_language,omitempty"`
	VideoTitle        string `json:"video_title,omitempty"`
	OriginalURL       string `json:"original_url"`
	ThumbnailURL      string `json:"thumbnail_url,omitempty"`
}