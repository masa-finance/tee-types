package args

// TikTokTranscriptionArguments defines args for TikTok transcriptions
type TikTokTranscriptionArguments struct {
	VideoURL string `json:"video_url"`
	Language string `json:"language,omitempty"` // e.g., "eng-US"
}
