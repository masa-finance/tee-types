# tee-types

A shared type definitions package for Masa Finance TEE projects.

## Minimal Sharing Approach

This package follows a minimalist approach, sharing only the essential types needed for the interface between tee-worker and tee-indexer. This approach reduces coupling between the services while ensuring consistent communication.

Each service should implement their own internal types that extend or build upon these shared types as needed.

## Structure

*WIP*

## Usage

To use this package in your project, add it as a dependency:

```bash
go get github.com/masa-finance/tee-types
```

Then import the required packages:

```go
import "github.com/masa-finance/tee-types/types"
```

## Types Included

### Core Types (`types/job.go`)

- `Job`: Represents a task to be executed by a worker
  ```go
  type Job struct {
      Type      string       `json:"type"`
      Arguments JobArguments `json:"arguments"`
      UUID      string       `json:"-"`
      Nonce     string       `json:"quote"`
      WorkerID  string       `json:"worker_id"`
  }
  ```

- `JobArguments`: Map type for job arguments with unmarshal utility
  ```go
  type JobArguments map[string]interface{}
  ```

- `JobResult`: Represents the result of executing a job
- `JobRequest`: Represents a request to execute a job
- Common job type constants (Web, Twitter)

### Twitter Types (`types/twitter.go`)

- `TwitterSearchParams`: Parameters for Twitter searches
  ```go
  type TwitterSearchParams struct {
      ScraperType            string                `json:"type"`      // Type of search
      TwitterSearchArguments `json:"arguments"`    // Search arguments
  }
  ```

- `TwitterSearchArguments`: Arguments for Twitter searches
  ```go
  type TwitterSearchArguments struct {
      Query      string `json:"query"`       // Username or search query
      QueryType  string `json:"type"`        // Optional, type of search
      StartTime  string `json:"start_time"`  // Optional ISO timestamp
      EndTime    string `json:"end_time"`    // Optional ISO timestamp
      MaxResults int    `json:"max_results"` // Optional, max number of results
  }
  ```

### Cryptographic Types (`types/crypto.go`)

- `EncryptedRequest`: For secure job requests
- `Key`: Basic key representation

## Contributing

When adding new features to tee-worker or tee-indexer, consider whether the types should be added to this shared package first.