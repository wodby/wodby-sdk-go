# Wodby 2.0 SDK for Go

[![Build](https://github.com/wodby/wodby-sdk-go/actions/workflows/build.yml/badge.svg?branch=2.0)](https://github.com/wodby/wodby-sdk-go/actions/workflows/build.yml?query=branch%3A2.0)

Go client for the Wodby 2.0 public API. This branch maintains SDK 4.x.

## Version compatibility

| Wodby platform | SDK version | Branch | API reference |
| --- | --- | --- | --- |
| Wodby 1.0 | 3.x | [master](https://github.com/wodby/wodby-sdk-go/tree/master) | [Wodby 1.0 API](https://wodby.com/docs/1.0/api/) |
| Wodby 2.0 | 4.x | [2.0](https://github.com/wodby/wodby-sdk-go/tree/2.0) | [Wodby 2.0 API](https://wodby.com/docs/2.0/api/) |

Choose the SDK major version for your Wodby platform. Upgrading from SDK 3.x to 4.x changes the target platform to Wodby 2.0.

## Package

- [Go package documentation](https://pkg.go.dev/github.com/wodby/wodby-sdk-go/v4/pkg)
- Go package: `github.com/wodby/wodby-sdk-go/v4/pkg`

## Documentation

- [API reference](https://wodby.com/docs/2.0/api/)
- [OpenAPI schema](https://github.com/wodby/backend-api/blob/2.0/schema/openapi.yaml)
- [Generated SDK documentation](pkg/docs)

## Install

Install the versioned module package:

```bash
go get github.com/wodby/wodby-sdk-go/v4/pkg
```

## Authentication

Wodby API requests use an API key in the `X-API-KEY` header.
