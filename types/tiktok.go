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

type TikTokSearchByQueryResult struct {
	URL                   string            `json:"url"`
	ID                    string            `json:"id"`
	Desc                  string            `json:"desc"`
	CreateTime            string            `json:"create_time"`
	ScheduleTime          int64             `json:"schedule_time"`
	Video                 TikTokVideo       `json:"video"`
	Author                string            `json:"author"`
	Music                 TikTokMusic       `json:"music"`
	Challenges            []any             `json:"challenges"` // we don't have examples of this data yet...
	Stats                 TikTokStats       `json:"stats"`
	IsActivityItem        bool              `json:"is_activity_item"`
	DuetInfo              TikTokDuetInfo    `json:"duet_info"`
	WarnInfo              []any             `json:"warn_info"` // we don't have examples of this data yet...
	OriginalItem          bool              `json:"original_item"`
	OfficalItem           bool              `json:"offical_item"`
	TextExtra             []TikTokTextExtra `json:"text_extra"`
	Secret                bool              `json:"secret"`
	ForFriend             bool              `json:"for_friend"`
	Digged                bool              `json:"digged"`
	ItemCommentStatus     int               `json:"item_comment_status"`
	ShowNotPass           bool              `json:"show_not_pass"`
	VL1                   bool              `json:"vl1"`
	TakeDown              int               `json:"take_down"`
	ItemMute              bool              `json:"item_mute"`
	EffectStickers        []any             `json:"effect_stickers"` // we don't have examples of this data yet...
	AuthorStats           TikTokAuthorStats `json:"author_stats"`
	PrivateItem           bool              `json:"private_item"`
	DuetEnabled           bool              `json:"duet_enabled"`
	StitchEnabled         bool              `json:"stitch_enabled"`
	StickersOnItem        []any             `json:"stickers_on_item"` // we don't have examples of this data yet...
	IsAd                  bool              `json:"is_ad"`
	ShareEnabled          bool              `json:"share_enabled"`
	Comments              []any             `json:"comments"` // we don't have examples of this data yet...
	DuetDisplay           int               `json:"duet_display"`
	StitchDisplay         int               `json:"stitch_display"`
	IndexEnabled          bool              `json:"index_enabled"`
	DiversificationLabels []string          `json:"diversification_labels"`
	AdAuthorization       bool              `json:"ad_authorization"`
	AdLabelVersion        int               `json:"ad_label_version"`
	LocationCreated       string            `json:"location_created"`
	Nickname              string            `json:"nickname"`
	AuthorID              string            `json:"author_id"`
	AuthorSecID           string            `json:"author_sec_id"`
	AvatarThumb           string            `json:"avatar_thumb"`
	DownloadSetting       int               `json:"download_setting"`
	AuthorPrivate         bool              `json:"author_private"`
}

type TikTokSearchByTrending struct {
	CountryCode string `json:"country_code"`
	Cover       string `json:"cover"`
	Duration    int    `json:"duration"`
	ID          string `json:"id"`
	ItemID      string `json:"item_id"`
	ItemURL     string `json:"item_url"`
	Region      string `json:"region"`
	Title       string `json:"title"`
}

type TikTokVideo struct {
	ID            string              `json:"id"`
	Height        int                 `json:"height"`
	Width         int                 `json:"width"`
	Duration      int                 `json:"duration"`
	Ratio         string              `json:"ratio"`
	Cover         string              `json:"cover"`
	OriginCover   string              `json:"origin_cover"`
	DynamicCover  string              `json:"dynamic_cover"`
	PlayAddr      string              `json:"play_addr"`
	DownloadAddr  string              `json:"download_addr"`
	ShareCover    []string            `json:"share_cover"`
	ReflowCover   string              `json:"reflow_cover"`
	Bitrate       int                 `json:"bitrate"`
	EncodedType   string              `json:"encoded_type"`
	Format        string              `json:"format"`
	VideoQuality  string              `json:"video_quality"`
	EncodeUserTag string              `json:"encode_user_tag"`
	CodecType     string              `json:"codec_type"`
	Definition    string              `json:"definition"`
	SubtitleInfos []any               `json:"subtitle_infos"` // we don't have examples of this data yet...
	ZoomCover     TikTokZoomCover     `json:"zoom_cover"`
	VolumeInfo    TikTokVolumeInfo    `json:"volume_info"`
	BitrateInfo   []TikTokBitrateInfo `json:"bitrate_info"`
}

type TikTokZoomCover struct {
	Cover240 string `json:"240"`
	Cover480 string `json:"480"`
	Cover720 string `json:"720"`
	Cover960 string `json:"960"`
}

type TikTokVolumeInfo struct {
	Loudness float64 `json:"loudness"`
	Peak     float64 `json:"peak"`
}

type TikTokBitrateInfo struct {
	GearName    string         `json:"gear_name"`
	Bitrate     int            `json:"bitrate"`
	QualityType int            `json:"quality_type"`
	PlayAddr    TikTokPlayAddr `json:"play_addr"`
	CodecType   string         `json:"codec_type"`
}

type TikTokPlayAddr struct {
	Uri      string   `json:"uri"`
	UrlList  []string `json:"url_list"`
	DataSize string   `json:"data_size"`
	UrlKey   string   `json:"url_key"`
	FileHash string   `json:"file_hash"`
	FileCs   string   `json:"file_cs"`
}

type TikTokMusic struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	PlayURL            string `json:"play_url"`
	CoverLarge         string `json:"cover_large"`
	CoverMedium        string `json:"cover_medium"`
	CoverThumb         string `json:"cover_thumb"`
	AuthorName         string `json:"author_name"`
	Original           bool   `json:"original"`
	Duration           int    `json:"duration"`
	Album              string `json:"album"`
	ScheduleSearchTime int64  `json:"schedule_search_time"`
}

type TikTokStats struct {
	DiggCount    int64 `json:"digg_count"`
	ShareCount   int64 `json:"share_count"`
	CommentCount int64 `json:"comment_count"`
	PlayCount    int64 `json:"play_count"`
}

type TikTokDuetInfo struct {
	DuetFromID string `json:"duet_from_id"`
}

type TikTokTextExtra struct {
	AwemeID      string `json:"aweme_id"`
	Start        int    `json:"start"`
	End          int    `json:"end"`
	HashtagID    string `json:"hashtag_id"`
	HashtagName  string `json:"hashtag_name"`
	Type         int    `json:"type"`
	SubType      int    `json:"sub_type"`
	UserID       string `json:"user_id"`
	IsCommerce   bool   `json:"is_commerce"`
	UserUniqueID string `json:"user_unique_id"`
	SecUID       string `json:"sec_uid"`
}

type TikTokAuthorStats struct {
	FollowerCount  int64 `json:"follower_count"`
	FollowingCount int64 `json:"following_count"`
	Heart          int64 `json:"heart"`
	HeartCount     int64 `json:"heart_count"`
	VideoCount     int64 `json:"video_count"`
	DiggCount      int64 `json:"digg_count"`
}
