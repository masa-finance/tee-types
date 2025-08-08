package types

import (
	"fmt"
	"slices"

	"github.com/masa-finance/tee-types/pkg/util"
)

type JobType string
type Capability string
type WorkerCapabilities map[JobType][]Capability

// String returns the string representation of the JobType
func (j JobType) String() string {
	return string(j)
}

// ValidateCapability validates that a capability is supported for this job type
func (j JobType) ValidateCapability(capability Capability) error {
	validCaps, exists := JobCapabilityMap[j]
	if !exists {
		return fmt.Errorf("unknown job type: %s", j)
	}

	if !slices.Contains(validCaps, capability) {
		return fmt.Errorf("capability '%s' is not valid for job type '%s'. Valid capabilities: %v",
			capability, j, validCaps)
	}

	return nil
}

// combineCapabilities combines multiple capability slices and ensures uniqueness
func combineCapabilities(capSlices ...[]Capability) []Capability {
	caps := util.NewSet[Capability]()
	for _, capSlice := range capSlices {
		caps.Add(capSlice...)
	}
	return caps.Items()
}

// Job type constants - centralized from tee-indexer and tee-worker
const (
	WebJob               JobType = "web"
	TelemetryJob         JobType = "telemetry"
	TiktokJob            JobType = "tiktok"
	TwitterJob           JobType = "twitter"            // General Twitter scraping (uses best available auth for capability)
	TwitterCredentialJob JobType = "twitter-credential" // Twitter scraping with credentials
	TwitterApiJob        JobType = "twitter-api"        // Twitter scraping with API keys
	TwitterApifyJob      JobType = "twitter-apify"      // Twitter scraping with Apify
	LinkedInJob          JobType = "linkedin"           // LinkedIn scraping, keeping for unmarshalling logic
)

// Capability constants - typed to prevent typos and enable discoverability
const (
	CapScraper             Capability = "scraper"
	CapTelemetry           Capability = "telemetry"
	CapTranscription       Capability = "transcription"
	CapSearchByQuery       Capability = "searchbyquery"
	CapSearchByFullArchive Capability = "searchbyfullarchive"
	CapSearchByProfile     Capability = "searchbyprofile"
	CapGetById             Capability = "getbyid"
	CapGetReplies          Capability = "getreplies"
	CapGetRetweeters       Capability = "getretweeters"
	CapGetTweets           Capability = "gettweets"
	CapGetMedia            Capability = "getmedia"
	CapGetHomeTweets       Capability = "gethometweets"
	CapGetForYouTweets     Capability = "getforyoutweets"
	CapGetProfileById      Capability = "getprofilebyid"
	CapGetTrends           Capability = "gettrends"
	CapGetFollowing        Capability = "getfollowing"
	CapGetFollowers        Capability = "getfollowers"
	CapGetSpace            Capability = "getspace"
	CapGetProfile          Capability = "getprofile" // LinkedIn get profile capability
	CapEmpty               Capability = ""
)

// Capability group constants for easy reuse
var (
	AlwaysAvailableWebCaps       = []Capability{CapScraper, CapEmpty}
	AlwaysAvailableTelemetryCaps = []Capability{CapTelemetry, CapEmpty}
	AlwaysAvailableTiktokCaps    = []Capability{CapTranscription, CapEmpty}
	AlwaysAvailableLinkedInCaps  = []Capability{CapSearchByQuery, CapGetProfile, CapEmpty}

	// AlwaysAvailableCapabilities defines the job capabilities that are always available regardless of configuration
	AlwaysAvailableCapabilities = WorkerCapabilities{
		WebJob:       AlwaysAvailableWebCaps,
		TelemetryJob: AlwaysAvailableTelemetryCaps,
		TiktokJob:    AlwaysAvailableTiktokCaps,
	}

	// TwitterCredentialCaps are all Twitter capabilities available with credential-based auth
	TwitterCredentialCaps = []Capability{
		CapSearchByQuery, CapSearchByProfile,
		CapGetById, CapGetReplies, CapGetRetweeters, CapGetTweets, CapGetMedia,
		CapGetHomeTweets, CapGetForYouTweets, CapGetProfileById,
		CapGetTrends, CapGetFollowing, CapGetFollowers, CapGetSpace,
		CapEmpty,
	}

	// TwitterAPICaps are basic Twitter capabilities available with API keys
	TwitterAPICaps = []Capability{CapSearchByQuery, CapGetById, CapGetProfileById, CapEmpty}

	// TwitterApifyCaps are Twitter capabilities available with Apify
	TwitterApifyCaps = []Capability{CapGetFollowers, CapGetFollowing, CapEmpty}
)

// JobCapabilityMap defines which capabilities are valid for each job type
var JobCapabilityMap = map[JobType][]Capability{
	// Twitter job types and their valid capabilities
	TwitterJob: combineCapabilities(
		TwitterCredentialCaps,
		TwitterAPICaps,
		TwitterApifyCaps,
		[]Capability{CapSearchByFullArchive}, // Elevated API capability
	),
	TwitterCredentialJob: TwitterCredentialCaps,
	TwitterApiJob: combineCapabilities(
		TwitterAPICaps,
		[]Capability{CapSearchByFullArchive}, // Elevated API capability
	),
	TwitterApifyJob: TwitterApifyCaps,

	// Web job capabilities
	WebJob: AlwaysAvailableWebCaps,

	// TikTok job capabilities
	TiktokJob: AlwaysAvailableTiktokCaps,

	// Telemetry job capabilities
	TelemetryJob: AlwaysAvailableTelemetryCaps,
}

// if no capability is specified, use the default capability for the job type
var JobDefaultCapabilityMap = map[JobType]Capability{
	TwitterJob:           CapSearchByQuery,
	TwitterCredentialJob: CapSearchByQuery,
	TwitterApiJob:        CapSearchByQuery,
	TwitterApifyJob:      CapGetFollowers,
	WebJob:               CapScraper,
	TiktokJob:            CapTranscription,
	TelemetryJob:         CapTelemetry,
}
