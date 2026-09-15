# Wodby SDK Go


The Wodby SDK for Go makes it easy for developers to access Wodby in their Go code. You can get started in minutes by installing the SDK with the language package manager.

---

* [Documentation](#documentation)
* [Basic usage](#basic-usage)

## Requirements

Go 1.26 or newer. Uses standard-library context and golang.org/x/oauth2 v0.37.0. These minimums apply to the next 3.x release; existing releases are unchanged.

SDK 3.x targets Wodby 1. SDK 4.x targets Wodby 2.

## Documentation

* [API reference](https://wodby.com/docs/1.0/api)
* [Automatically generated documentation](pkg/README.md)

## Basic usage

Import the client using its versioned module path:

```go
import client "github.com/wodby/wodby-sdk-go/v3/pkg"
```

When upgrading, replace the previous import path with this `/v3` path and run `go mod tidy`. The next release will also publish an annotated `v3.X.Y` module tag alongside its `3.X.Y` release tag. Existing 3.x tags predate module support.

## Development

```sh
go test -race ./...
go vet ./...
```

Tests use mocks or a local HTTP server and do not require an API key.

Regenerate with `make codegen` (Docker and Java 17 image required). The existing Swagger generator stays pinned to preserve the client API; the post-generation script reapplies modern dependency compatibility.
