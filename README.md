# tee-types

A shared type definitions package for Masa Finance TEE projects.

## Overview

This package provides essential type definitions for communication between tee-worker and tee-indexer services. It supports various social media platforms including LinkedIn, Twitter, and TikTok, enabling both search and detailed profile fetching operations.

## Minimal Sharing Approach

This package follows a minimalist approach, sharing only the essential types needed for the interface between tee-worker and tee-indexer. This approach reduces coupling between the services while ensuring consistent communication.

Each service should implement their own internal types that extend or build upon these shared types as needed.

## Structure

### Arguments (`args/`)
- `linkedin.go` - LinkedIn operation arguments supporting both search and profile fetching
- `twitter.go` - Twitter-specific arguments
- `tiktok.go` - TikTok-specific arguments
- `web.go` - General web scraping arguments

### Types (`types/`)
- `linkedin.go` - LinkedIn result types including search results and detailed profile data
- `twitter.go` - Twitter result structures
- `tiktok.go` - TikTok result structures

## LinkedIn Support

### Search Operations
Use `LinkedInArguments` (or the deprecated `LinkedInSearchArguments`) for profile searches:

```go
args := &args.LinkedInArguments{
    QueryType: "searchbyquery",
    Query: "software engineer",
    MaxResults: 10,
    Start: 0,
}
```

### Profile Fetching
Use `LinkedInArguments` with `PublicIdentifier` for detailed profile retrieval:

```go
args := &args.LinkedInArguments{
    QueryType: "getprofile",
    PublicIdentifier: "john-doe-123",
}
```

### Result Types
- `LinkedInProfileResult` - Basic profile information from search results
- `LinkedInFullProfileResult` - Comprehensive profile data including experience, education, and skills

## Usage

To use this package in your project, add it as a dependency:

```bash
go get github.com/masa-finance/tee-types@v1.0.0
```

Then import the required packages:

```go
import (
    "github.com/masa-finance/tee-types/args"
    "github.com/masa-finance/tee-types/types"
)
```

## Backward Compatibility

The package maintains full backward compatibility. Existing code using `LinkedInSearchArguments` will continue to work, though migration to `LinkedInArguments` is recommended for future compatibility.

## Releases

- **v1.0.0** - Initial release with LinkedIn profile fetching support