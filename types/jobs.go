package types

type Capability string

// JobType represents the type of job that can be executed
type JobType string

// Job type constants - centralized from tee-indexer and tee-worker
const (
	// Web scraping job type
	WebJob JobType = "web"

	// Telemetry job type for worker monitoring and stats
	TelemetryJob JobType = "telemetry"

	// TikTok transcription job type
	TiktokJob JobType = "tiktok"

	// Twitter job types
	TwitterJob           JobType = "twitter"            // General Twitter scraping (uses best available auth)
	TwitterCredentialJob JobType = "twitter-credential" // Twitter scraping with credentials
	TwitterApiJob        JobType = "twitter-api"        // Twitter scraping with API keys
	TwitterApifyJob      JobType = "twitter-apify"      // Twitter scraping with Apify

	// LinkedIn job types
	LinkedInJob JobType = "linkedin" // LinkedIn scraping and profile operations

)

// Capability constants - typed to prevent typos and enable discoverability
const (
	// Web scraping capabilities
	CapWebScraper Capability = "scraper"

	// Telemetry capabilities
	CapTelemetry Capability = "telemetry"

	// TikTok capabilities
	CapTiktokTranscription Capability = "transcription"

	// Twitter capabilities
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

	// LinkedIn capabilities
	CapGetProfile Capability = "getprofile"
)

// Capability group constants for easy reuse
var (
	// AlwaysAvailableWebCaps are web capabilities always available
	AlwaysAvailableWebCaps = []Capability{CapWebScraper}

	// AlwaysAvailableTelemetryCaps are telemetry capabilities always available
	AlwaysAvailableTelemetryCaps = []Capability{CapTelemetry}

	// AlwaysAvailableTiktokCaps are TikTok capabilities always available
	AlwaysAvailableTiktokCaps = []Capability{CapTiktokTranscription}

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

	// LinkedInCaps are LinkedIn capabilities (basic set for future implementation)
	LinkedInCaps = []Capability{CapSearchByQuery, CapGetProfile}

	// AlwaysAvailableCapabilities defines the job capabilities that are always available regardless of configuration
	AlwaysAvailableCapabilities = WorkerCapabilities{
		WebJob:       AlwaysAvailableWebCaps,
		TelemetryJob: AlwaysAvailableTelemetryCaps,
		TiktokJob:    AlwaysAvailableTiktokCaps,
	}
)

// String returns the string representation of the JobType
func (j JobType) String() string {
	return string(j)
}

// WorkerCapabilities represents all capabilities available on a worker
// Maps JobType to the list of capabilities available for that job type
type WorkerCapabilities map[JobType][]Capability
