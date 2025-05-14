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