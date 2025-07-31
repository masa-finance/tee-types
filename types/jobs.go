package types

type JobType string
type Capability string
type WorkerCapabilities map[JobType][]Capability

// String returns the string representation of the JobType
func (j JobType) String() string {
	return string(j)
}

// combineCapabilities combines multiple capability slices and ensures uniqueness
func combineCapabilities(capSlices ...[]Capability) []Capability {
	seen := make(map[Capability]bool)
	var result []Capability

	for _, capSlice := range capSlices {
		for _, cap := range capSlice {
			if !seen[cap] {
				seen[cap] = true
				result = append(result, cap)
			}
		}
	}

	return result
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
)

// Capability group constants for easy reuse
var (
	AlwaysAvailableWebCaps       = []Capability{CapScraper}
	AlwaysAvailableTelemetryCaps = []Capability{CapTelemetry}
	AlwaysAvailableTiktokCaps    = []Capability{CapTranscription}

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
	}

	// TwitterAPICaps are basic Twitter capabilities available with API keys
	TwitterAPICaps = []Capability{CapSearchByQuery, CapGetById, CapGetProfileById}

	// TwitterApifyCaps are Twitter capabilities available with Apify
	TwitterApifyCaps = []Capability{CapGetFollowers, CapGetFollowing}
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
