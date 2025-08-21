package args

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/masa-finance/tee-types/pkg/util"
	teetypes "github.com/masa-finance/tee-types/types"
)

var (
	ErrRedditInvalidType       = errors.New("invalid type")
	ErrRedditInvalidSort       = errors.New("invalid sort")
	ErrRedditTimeInTheFuture   = errors.New("after field is in the future")
	ErrRedditNoQueries         = errors.New("queries must be provided for all query types except scrapeurls")
	ErrRedditNoUrls            = errors.New("urls must be provided for scrapeurls query type")
	ErrRedditQueriesNotAllowed = errors.New("the scrapeurls query type does not admit queries")
	ErrRedditUrlsNotAllowed    = errors.New("urls can only be provided for the scrapeurls query type")
)

const (
	// These reflect the default values in https://apify.com/trudax/reddit-scraper/input-schema
	redditDefaultMaxItems       = 10
	redditDefaultMaxPosts       = 10
	redditDefaultMaxComments    = 10
	redditDefaultMaxCommunities = 2
	redditDefaultMaxUsers       = 2
	redditDefaultSort           = teetypes.RedditSortNew
)

// RedditArguments defines args for Reddit scrapes
// see https://apify.com/trudax/reddit-scraper
type RedditArguments struct {
	QueryType      teetypes.RedditQueryType  `json:"type"`
	Queries        []string                  `json:"queries"`
	URLs           []teetypes.RedditStartURL `json:"urls"`
	Sort           teetypes.RedditSortType   `json:"sort"`
	IncludeNSFW    bool                      `json:"include_nsfw"`
	SkipPosts      bool                      `json:"skip_posts"`      // Valid only for searchusers
	After          time.Time                 `json:"after"`           // valid only for scrapeurls and searchposts
	MaxItems       uint                      `json:"max_items"`       // Max number of items to scrape (total), default 10
	MaxResults     uint                      `json:"max_results"`     // Max number of results per page, default MaxItems
	MaxPosts       uint                      `json:"max_posts"`       // Max number of posts per page, default 10
	MaxComments    uint                      `json:"max_comments"`    // Max number of comments per page, default 10
	MaxCommunities uint                      `json:"max_communities"` // Max number of communities per page, default 2
	MaxUsers       uint                      `json:"max_users"`       // Max number of users per page, default 2
	NextCursor     string                    `json:"next_cursor"`
}

func (r *RedditArguments) UnmarshalJSON(data []byte) error {
	type Alias RedditArguments

	// Set default values. They will be overridden if present in the JSON.
	r.MaxItems = redditDefaultMaxItems
	r.MaxPosts = redditDefaultMaxPosts
	r.MaxComments = redditDefaultMaxComments
	r.MaxCommunities = redditDefaultMaxCommunities
	r.MaxUsers = redditDefaultMaxUsers
	r.Sort = redditDefaultSort

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal Reddit arguments: %w", err)
	}

	if r.MaxResults == 0 {
		r.MaxResults = r.MaxItems
	}

	return r.Validate()
}

var allowedHttpMethods = util.NewSet("GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS")

const redditDomainSuffix = "reddit.com"

func (r *RedditArguments) Validate() error {
	var errs []error
	if !teetypes.AllRedditQueryTypes.Contains(r.QueryType) {
		errs = append(errs, ErrRedditInvalidType)
	}

	if !teetypes.AllRedditSortTypes.Contains(r.Sort) {
		errs = append(errs, ErrRedditInvalidSort)
	}

	if time.Now().Before(r.After) {
		errs = append(errs, ErrRedditTimeInTheFuture)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	if r.QueryType == teetypes.RedditScrapeUrls {
		if len(r.URLs) == 0 {
			errs = append(errs, ErrRedditNoUrls)
		}
		if len(r.Queries) > 0 {
			errs = append(errs, ErrRedditQueriesNotAllowed)
		}

		for _, q := range r.URLs {
			if !allowedHttpMethods.Contains(q.Method) {
				errs = append(errs, fmt.Errorf("%s is not a valid HTTP method", q.Method))
			}
			u, err := url.Parse(q.URL)
			if err != nil {
				errs = append(errs, fmt.Errorf("%s is not a valid URL", q.URL))
			} else {
				if !strings.HasSuffix(u.Host, redditDomainSuffix) {
					errs = append(errs, fmt.Errorf("invalid Reddit URL %s", q.URL))
				}
			}
		}
	} else {
		if len(r.Queries) == 0 {
			errs = append(errs, ErrRedditNoQueries)
		}
		if len(r.URLs) > 0 {
			errs = append(errs, ErrRedditUrlsNotAllowed)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// ValidateForJobType validates Twitter arguments for a specific job type
func (r *RedditArguments) ValidateForJobType(jobType teetypes.JobType) error {
	if err := r.Validate(); err != nil {
		return err
	}

	// Validate QueryType against job-specific capabilities
	return jobType.ValidateCapability(teetypes.Capability(r.QueryType))
}

// GetCapability returns the QueryType as a typed Capability
func (r *RedditArguments) GetCapability() teetypes.Capability {
	return teetypes.Capability(r.QueryType)
}
