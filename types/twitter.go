// Package types provides shared types between tee-worker and tee-indexer
package types

import "time"

type TweetResult struct {
	ID             int64 `json:"id"`
	TweetID        string
	ConversationID string
	UserID         string
	Text           string
	CreatedAt      time.Time
	Timestamp      int64

	ThreadCursor struct {
		FocalTweetID string
		ThreadID     string
		Cursor       string
		CursorType   string
	}
	IsQuoted     bool
	IsPin        bool
	IsReply      bool
	IsRetweet    bool
	IsSelfThread bool
	Likes        int
	Hashtags     []string
	HTML         string
	Replies      int
	Retweets     int
	URLs         []string
	Username     string

	Photos []Photo

	// Video type.
	Videos []Video

	RetweetedStatusID string
	Views             int
	SensitiveContent  bool

	// from twitterx
	AuthorID          string
	PublicMetrics     PublicMetrics
	PossiblySensitive bool
	Lang              string
	NewestID          string
	OldestID          string
	ResultCount       int

	Error error
}

type PublicMetrics struct {
	RetweetCount    int
	ReplyCount      int
	LikeCount       int
	QuoteCount      int
	BookmarkCount   int
	ImpressionCount int
}
type Photo struct {
	ID  string
	URL string
}

type Video struct {
	ID      string
	Preview string
	URL     string
	HLSURL  string
}

type ProfileResultApify struct {
	ID                             int64           `json:"id"`
	IDStr                          string          `json:"id_str"`
	Name                           string          `json:"name"`
	ScreenName                     string          `json:"screen_name"`
	Location                       string          `json:"location"`
	Description                    string          `json:"description"`
	URL                            *string         `json:"url"`
	Entities                       ProfileEntities `json:"entities"`
	Protected                      bool            `json:"protected"`
	FollowersCount                 int             `json:"followers_count"`
	FastFollowersCount             int             `json:"fast_followers_count"`
	NormalFollowersCount           int             `json:"normal_followers_count"`
	FriendsCount                   int             `json:"friends_count"`
	ListedCount                    int             `json:"listed_count"`
	CreatedAt                      string          `json:"created_at"`
	FavouritesCount                int             `json:"favourites_count"`
	UTCOffset                      *int            `json:"utc_offset"`
	TimeZone                       *string         `json:"time_zone"`
	GeoEnabled                     bool            `json:"geo_enabled"`
	Verified                       bool            `json:"verified"`
	StatusesCount                  int             `json:"statuses_count"`
	MediaCount                     int             `json:"media_count"`
	Lang                           *string         `json:"lang"`
	ContributorsEnabled            bool            `json:"contributors_enabled"`
	IsTranslator                   bool            `json:"is_translator"`
	IsTranslationEnabled           bool            `json:"is_translation_enabled"`
	ProfileBackgroundColor         string          `json:"profile_background_color"`
	ProfileBackgroundImageURL      *string         `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS *string         `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool            `json:"profile_background_tile"`
	ProfileImageURL                string          `json:"profile_image_url"`
	ProfileImageURLHTTPS           string          `json:"profile_image_url_https"`
	ProfileLinkColor               string          `json:"profile_link_color"`
	ProfileSidebarBorderColor      string          `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string          `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string          `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool            `json:"profile_use_background_image"`
	HasExtendedProfile             bool            `json:"has_extended_profile"`
	DefaultProfile                 bool            `json:"default_profile"`
	DefaultProfileImage            bool            `json:"default_profile_image"`
	PinnedTweetIDs                 []int64         `json:"pinned_tweet_ids"`
	PinnedTweetIDsStr              []string        `json:"pinned_tweet_ids_str"`
	HasCustomTimelines             bool            `json:"has_custom_timelines"`
	CanMediaTag                    bool            `json:"can_media_tag"`
	FollowedBy                     bool            `json:"followed_by"`
	Following                      bool            `json:"following"`
	LiveFollowing                  bool            `json:"live_following"`
	FollowRequestSent              bool            `json:"follow_request_sent"`
	Notifications                  bool            `json:"notifications"`
	Muting                         bool            `json:"muting"`
	Blocking                       bool            `json:"blocking"`
	BlockedBy                      bool            `json:"blocked_by"`
	AdvertiserAccountType          string          `json:"advertiser_account_type"`
	AdvertiserAccountServiceLevels []string        `json:"advertiser_account_service_levels"`
	AnalyticsType                  string          `json:"analytics_type"`
	BusinessProfileState           string          `json:"business_profile_state"`
	TranslatorType                 string          `json:"translator_type"`
	WithheldInCountries            []string        `json:"withheld_in_countries"`
	RequireSomeConsent             bool            `json:"require_some_consent"`
	Type                           string          `json:"type"`
	TargetUsername                 string          `json:"target_username"`
	Email                          *string         `json:"email"`
}

type ProfileEntities struct {
	URL         *URLEntities `json:"url,omitempty"`
	Description *URLEntities `json:"description,omitempty"`
}

type URLEntities struct {
	URLs []URLEntity `json:"urls,omitempty"`
}

type URLEntity struct {
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	Indices     []int  `json:"indices"`
}
