package args_test

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/masa-finance/tee-types/args"
	"github.com/masa-finance/tee-types/types"
)

var _ = Describe("RedditArguments", func() {
	Describe("Unmarshalling", func() {
		It("should set default values", func() {
			redditArgs := &args.RedditArguments{}
			jsonData := `{"type": "searchposts", "queries": ["test"]}`
			err := json.Unmarshal([]byte(jsonData), redditArgs)
			Expect(err).ToNot(HaveOccurred())
			Expect(redditArgs.MaxItems).To(Equal(uint(10)))
			Expect(redditArgs.MaxPosts).To(Equal(uint(10)))
			Expect(redditArgs.MaxComments).To(Equal(uint(10)))
			Expect(redditArgs.MaxCommunities).To(Equal(uint(2)))
			Expect(redditArgs.MaxUsers).To(Equal(uint(2)))
			Expect(redditArgs.Sort).To(Equal(types.RedditSortNew))
			Expect(redditArgs.MaxResults).To(Equal(redditArgs.MaxItems))
		})

		It("should override default values", func() {
			redditArgs := &args.RedditArguments{}
			jsonData := `{"type": "searchposts", "queries": ["test"], "max_items": 20, "sort": "top"}`
			err := json.Unmarshal([]byte(jsonData), redditArgs)
			Expect(err).ToNot(HaveOccurred())
			Expect(redditArgs.MaxItems).To(Equal(uint(20)))
			Expect(redditArgs.Sort).To(Equal(types.RedditSortTop))
			Expect(redditArgs.MaxResults).To(Equal(uint(20)))
		})
	})

	Describe("Validation", func() {
		It("should succeed with valid arguments", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditSearchPosts,
				Queries:   []string{"test"},
				Sort:      types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should succeed with valid scrapeurls arguments", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditScrapeUrls,
				URLs: []types.RedditStartURL{
					{URL: "https://www.reddit.com/r/golang/", Method: "GET"},
				},
				Sort: types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should fail with an invalid type", func() {
			redditArgs := &args.RedditArguments{
				QueryType: "invalidtype",
				Queries:   []string{"test"},
				Sort:      types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditInvalidType))
		})

		It("should fail with an invalid sort", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditSearchPosts,
				Queries:   []string{"test"},
				Sort:      "invalidsort",
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditInvalidSort))
		})

		It("should fail if the after time is in the future", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditSearchPosts,
				Queries:   []string{"test"},
				Sort:      types.RedditSortNew,
				After:     time.Now().Add(24 * time.Hour),
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditTimeInTheFuture))
		})

		It("should fail if queries are not provided for searchposts", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditSearchPosts,
				Sort:      types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditNoQueries))
		})

		It("should fail if urls are not provided for scrapeurls", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditScrapeUrls,
				Sort:      types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditNoUrls))
		})

		It("should fail if queries are provided for scrapeurls", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditScrapeUrls,
				Queries:   []string{"test"},
				URLs: []types.RedditStartURL{
					{URL: "https://www.reddit.com/r/golang/", Method: "GET"},
				},
				Sort: types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditQueriesNotAllowed))
		})

		It("should fail if urls are provided for searchposts", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditSearchPosts,
				Queries:   []string{"test"},
				URLs: []types.RedditStartURL{
					{URL: "https://www.reddit.com/r/golang/", Method: "GET"},
				},
				Sort: types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(MatchError(args.ErrRedditUrlsNotAllowed))
		})

		It("should fail with an invalid URL", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditScrapeUrls,
				URLs: []types.RedditStartURL{
					{URL: "ht tp://invalid-url.com", Method: "GET"},
				},
				Sort: types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not a valid URL"))
		})

		It("should fail with an invalid domain", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditScrapeUrls,
				URLs: []types.RedditStartURL{
					{URL: "https://www.google.com", Method: "GET"},
				},
				Sort: types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid Reddit URL"))
		})

		It("should fail with an invalid HTTP method", func() {
			redditArgs := &args.RedditArguments{
				QueryType: types.RedditScrapeUrls,
				URLs: []types.RedditStartURL{
					{URL: "https://www.reddit.com/r/golang/", Method: "INVALID"},
				},
				Sort: types.RedditSortNew,
			}
			err := redditArgs.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not a valid HTTP method"))
		})
	})
})
