# Wodby 1.0 SDK for Go

[![Build](https://github.com/wodby/wodby-sdk-go/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/wodby/wodby-sdk-go/actions/workflows/build.yml?query=branch%3Amaster)

Go client for the Wodby 1.0 public API. This branch maintains SDK 3.x.

## Version compatibility

| Wodby platform | SDK version | Branch | API reference |
| --- | --- | --- | --- |
| Wodby 1.0 | 3.x | [master](https://github.com/wodby/wodby-sdk-go/tree/master) | [Wodby 1.0 API](https://wodby.com/docs/1.0/api/) |
| Wodby 2.0 | 4.x | [2.0](https://github.com/wodby/wodby-sdk-go/tree/2.0) | [Wodby 2.0 API](https://wodby.com/docs/2.0/api/) |

Choose the SDK major version for your Wodby platform. Upgrading from SDK 3.x to 4.x changes the target platform to Wodby 2.0.

## Requirements

Go 1.26 or newer. Uses standard-library context and golang.org/x/oauth2 v0.37.0. These minimums apply to the next 3.x release; existing releases are unchanged.

## Documentation

* [API reference](https://wodby.com/docs/1.0/api/)
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
