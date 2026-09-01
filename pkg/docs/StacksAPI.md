# \StacksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStackEnvVar**](StacksAPI.md#CreateStackEnvVar) | **Post** /stacks/{id}/configuration/env-vars | Create stack env var
[**CreateStackFromManifest**](StacksAPI.md#CreateStackFromManifest) | **Post** /stacks/actions/create-from-manifest | Create stack from manifest
[**DeleteStack**](StacksAPI.md#DeleteStack) | **Delete** /stacks/{id} | Delete stack
[**DeleteStackEnvVar**](StacksAPI.md#DeleteStackEnvVar) | **Delete** /stack-env-vars/{id} | Delete stack env var
[**DuplicateStack**](StacksAPI.md#DuplicateStack) | **Post** /stacks/{id}/actions/duplicate | Duplicate stack
[**GetStack**](StacksAPI.md#GetStack) | **Get** /stacks/{id} | Get stack
[**GetStackByName**](StacksAPI.md#GetStackByName) | **Get** /stacks/by-name/{name} | Get stack by name
[**GetStackOriginSyncChangelog**](StacksAPI.md#GetStackOriginSyncChangelog) | **Get** /stack-origin-sync-changelogs/{id} | Preview stack origin synchronization
[**GetStackRevision**](StacksAPI.md#GetStackRevision) | **Get** /stack-revisions/{id} | Get stack revision
[**GetStackServiceUpdateChangelog**](StacksAPI.md#GetStackServiceUpdateChangelog) | **Get** /stack-service-update-changelogs/{id} | Preview stack service revision updates
[**ImportStacks**](StacksAPI.md#ImportStacks) | **Post** /stacks/actions/import | Import stacks from Git
[**ListPublicStacks**](StacksAPI.md#ListPublicStacks) | **Get** /catalog/stacks | List public catalog stacks
[**ListStackEnvVars**](StacksAPI.md#ListStackEnvVars) | **Get** /stack-revisions/{id}/env-vars | List stack env vars
[**ListStackRevisionServices**](StacksAPI.md#ListStackRevisionServices) | **Get** /stack-revisions/{id}/services | List stack services
[**ListStacks**](StacksAPI.md#ListStacks) | **Get** /stacks | List stacks
[**PublishStackDraft**](StacksAPI.md#PublishStackDraft) | **Post** /stacks/{id}/actions/publish-draft | Publish stack draft
[**ScaffoldStackFromHelmChart**](StacksAPI.md#ScaffoldStackFromHelmChart) | **Post** /stacks/actions/scaffold-from-helm-chart | Scaffold stack from Helm chart
[**SyncStackWithOrigin**](StacksAPI.md#SyncStackWithOrigin) | **Post** /stacks/{id}/actions/sync-origin | Sync stack with origin
[**UpdateStack**](StacksAPI.md#UpdateStack) | **Put** /stacks/{id} | Rename stack
[**UpdateStackEnvVar**](StacksAPI.md#UpdateStackEnvVar) | **Put** /stack-env-vars/{id} | Update stack env var
[**UpdateStackFromGit**](StacksAPI.md#UpdateStackFromGit) | **Post** /stacks/{id}/actions/update-from-git | Update stack from git
[**UpdateStackServiceRevisions**](StacksAPI.md#UpdateStackServiceRevisions) | **Post** /stacks/{id}/actions/update-service-revisions | Update stack service revisions
[**UpdateStackSettings**](StacksAPI.md#UpdateStackSettings) | **Put** /stacks/settings/{id} | Update stack settings
[**ValidateStackManifest**](StacksAPI.md#ValidateStackManifest) | **Post** /stacks/actions/validate-manifest | Validate stack manifest



## CreateStackEnvVar

> StackEnvVar CreateStackEnvVar(ctx, id).NewStackEnvVarInput(newStackEnvVarInput).Execute()

Create stack env var



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	newStackEnvVarInput := *openapiclient.NewNewStackEnvVarInput("Name_example", "Value_example", false) // NewStackEnvVarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.CreateStackEnvVar(context.Background(), id).NewStackEnvVarInput(newStackEnvVarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.CreateStackEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackEnvVar`: StackEnvVar
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.CreateStackEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStackEnvVarInput** | [**NewStackEnvVarInput**](NewStackEnvVarInput.md) |  | 

### Return type

[**StackEnvVar**](StackEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackFromManifest

> Stack CreateStackFromManifest(ctx).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()

Create stack from manifest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	manifestFromYAMLInput := *openapiclient.NewManifestFromYAMLInput("ManifestYaml_example") // ManifestFromYAMLInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.CreateStackFromManifest(context.Background()).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.CreateStackFromManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackFromManifest`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.CreateStackFromManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackFromManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manifestFromYAMLInput** | [**ManifestFromYAMLInput**](ManifestFromYAMLInput.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStack

> OperationResult DeleteStack(ctx, id).Execute()

Delete stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.DeleteStack(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.DeleteStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStack`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.DeleteStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStackEnvVar

> OperationResult DeleteStackEnvVar(ctx, id).Execute()

Delete stack env var



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.DeleteStackEnvVar(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.DeleteStackEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackEnvVar`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.DeleteStackEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuplicateStack

> Stack DuplicateStack(ctx, id).DuplicateStackRequest(duplicateStackRequest).Execute()

Duplicate stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	duplicateStackRequest := *openapiclient.NewDuplicateStackRequest() // DuplicateStackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.DuplicateStack(context.Background(), id).DuplicateStackRequest(duplicateStackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.DuplicateStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicateStack`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.DuplicateStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicateStackRequest** | [**DuplicateStackRequest**](DuplicateStackRequest.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStack

> Stack GetStack(ctx, id).Execute()

Get stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.GetStack(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.GetStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStack`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.GetStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackByName

> Stack GetStackByName(ctx, name).RevNumber(revNumber).Execute()

Get stack by name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	name := "name_example" // string | 
	revNumber := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.GetStackByName(context.Background(), name).RevNumber(revNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.GetStackByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackByName`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.GetStackByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revNumber** | **int32** |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackOriginSyncChangelog

> StackOriginSyncChangelog GetStackOriginSyncChangelog(ctx, id).Execute()

Preview stack origin synchronization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.GetStackOriginSyncChangelog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.GetStackOriginSyncChangelog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackOriginSyncChangelog`: StackOriginSyncChangelog
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.GetStackOriginSyncChangelog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackOriginSyncChangelogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackOriginSyncChangelog**](StackOriginSyncChangelog.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackRevision

> StackRevision GetStackRevision(ctx, id).Execute()

Get stack revision



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.GetStackRevision(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.GetStackRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackRevision`: StackRevision
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.GetStackRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackRevision**](StackRevision.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackServiceUpdateChangelog

> []StackServiceUpdateChangelog GetStackServiceUpdateChangelog(ctx, id).StackServiceId(stackServiceId).Execute()

Preview stack service revision updates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	stackServiceId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.GetStackServiceUpdateChangelog(context.Background(), id).StackServiceId(stackServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.GetStackServiceUpdateChangelog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackServiceUpdateChangelog`: []StackServiceUpdateChangelog
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.GetStackServiceUpdateChangelog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackServiceUpdateChangelogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackServiceId** | **int32** |  | 

### Return type

[**[]StackServiceUpdateChangelog**](StackServiceUpdateChangelog.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportStacks

> OperationResult ImportStacks(ctx).ImportCatalogFromGitInput(importCatalogFromGitInput).Execute()

Import stacks from Git



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	importCatalogFromGitInput := *openapiclient.NewImportCatalogFromGitInput(int32(123), "RemoteGitRepoId_example", "GitRef_example", "GitRefType_example") // ImportCatalogFromGitInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ImportStacks(context.Background()).ImportCatalogFromGitInput(importCatalogFromGitInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ImportStacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportStacks`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ImportStacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportStacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importCatalogFromGitInput** | [**ImportCatalogFromGitInput**](ImportCatalogFromGitInput.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicStacks

> []Stack ListPublicStacks(ctx).Execute()

List public catalog stacks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ListPublicStacks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ListPublicStacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicStacks`: []Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ListPublicStacks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPublicStacksRequest struct via the builder pattern


### Return type

[**[]Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackEnvVars

> []StackEnvVar ListStackEnvVars(ctx, id).Execute()

List stack env vars



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ListStackEnvVars(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ListStackEnvVars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackEnvVars`: []StackEnvVar
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ListStackEnvVars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackEnvVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackEnvVar**](StackEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackRevisionServices

> []StackService ListStackRevisionServices(ctx, id).Execute()

List stack services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ListStackRevisionServices(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ListStackRevisionServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackRevisionServices`: []StackService
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ListStackRevisionServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackRevisionServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackService**](StackService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStacks

> StacksResponse ListStacks(ctx).OrgId(orgId).ProjectIds(projectIds).Search(search).Page(page).PageSize(pageSize).Execute()

List stacks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)
	projectIds := "projectIds_example" // string | Comma-separated project ids (optional)
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ListStacks(context.Background()).OrgId(orgId).ProjectIds(projectIds).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ListStacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStacks`: StacksResponse
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ListStacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **search** | **string** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**StacksResponse**](StacksResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishStackDraft

> Stack PublishStackDraft(ctx, id).Execute()

Publish stack draft



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.PublishStackDraft(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.PublishStackDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishStackDraft`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.PublishStackDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishStackDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaffoldStackFromHelmChart

> HelmChartStackScaffoldResponse ScaffoldStackFromHelmChart(ctx).HelmChartStackScaffoldInput(helmChartStackScaffoldInput).Execute()

Scaffold stack from Helm chart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	helmChartStackScaffoldInput := *openapiclient.NewHelmChartStackScaffoldInput(*openapiclient.NewHelmChartInput("Chart_example")) // HelmChartStackScaffoldInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ScaffoldStackFromHelmChart(context.Background()).HelmChartStackScaffoldInput(helmChartStackScaffoldInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ScaffoldStackFromHelmChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScaffoldStackFromHelmChart`: HelmChartStackScaffoldResponse
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ScaffoldStackFromHelmChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScaffoldStackFromHelmChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmChartStackScaffoldInput** | [**HelmChartStackScaffoldInput**](HelmChartStackScaffoldInput.md) |  | 

### Return type

[**HelmChartStackScaffoldResponse**](HelmChartStackScaffoldResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncStackWithOrigin

> Stack SyncStackWithOrigin(ctx, id).StackSyncOptionsInput(stackSyncOptionsInput).Execute()

Sync stack with origin



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	stackSyncOptionsInput := *openapiclient.NewStackSyncOptionsInput() // StackSyncOptionsInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.SyncStackWithOrigin(context.Background(), id).StackSyncOptionsInput(stackSyncOptionsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.SyncStackWithOrigin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncStackWithOrigin`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.SyncStackWithOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncStackWithOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackSyncOptionsInput** | [**StackSyncOptionsInput**](StackSyncOptionsInput.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStack

> Stack UpdateStack(ctx, id).UpdateStackRequest(updateStackRequest).Execute()

Rename stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	updateStackRequest := *openapiclient.NewUpdateStackRequest("Name_example", "Title_example") // UpdateStackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.UpdateStack(context.Background(), id).UpdateStackRequest(updateStackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.UpdateStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStack`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.UpdateStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackRequest** | [**UpdateStackRequest**](UpdateStackRequest.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackEnvVar

> StackEnvVar UpdateStackEnvVar(ctx, id).UpdateStackEnvVarInput(updateStackEnvVarInput).Execute()

Update stack env var



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	updateStackEnvVarInput := *openapiclient.NewUpdateStackEnvVarInput("Value_example", false) // UpdateStackEnvVarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.UpdateStackEnvVar(context.Background(), id).UpdateStackEnvVarInput(updateStackEnvVarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.UpdateStackEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackEnvVar`: StackEnvVar
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.UpdateStackEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackEnvVarInput** | [**UpdateStackEnvVarInput**](UpdateStackEnvVarInput.md) |  | 

### Return type

[**StackEnvVar**](StackEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackFromGit

> OperationResult UpdateStackFromGit(ctx, id).UpdateStackFromGitRequest(updateStackFromGitRequest).Execute()

Update stack from git



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	updateStackFromGitRequest := *openapiclient.NewUpdateStackFromGitRequest("GitRef_example", "GitRefType_example") // UpdateStackFromGitRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.UpdateStackFromGit(context.Background(), id).UpdateStackFromGitRequest(updateStackFromGitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.UpdateStackFromGit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackFromGit`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.UpdateStackFromGit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackFromGitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackFromGitRequest** | [**UpdateStackFromGitRequest**](UpdateStackFromGitRequest.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackServiceRevisions

> OperationResult UpdateStackServiceRevisions(ctx, id).Scope(scope).Execute()

Update stack service revisions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	scope := "scope_example" // string | Limits the update to all services or stateless services. Defaults to all. (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.UpdateStackServiceRevisions(context.Background(), id).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.UpdateStackServiceRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackServiceRevisions`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.UpdateStackServiceRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackServiceRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** | Limits the update to all services or stateless services. Defaults to all. | [default to &quot;all&quot;]

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackSettings

> Stack UpdateStackSettings(ctx, id).StackSettingsInput(stackSettingsInput).Execute()

Update stack settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	id := int32(56) // int32 | 
	stackSettingsInput := *openapiclient.NewStackSettingsInput() // StackSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.UpdateStackSettings(context.Background(), id).StackSettingsInput(stackSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.UpdateStackSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackSettings`: Stack
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.UpdateStackSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackSettingsInput** | [**StackSettingsInput**](StackSettingsInput.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateStackManifest

> ManifestValidationResponse ValidateStackManifest(ctx).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()

Validate stack manifest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	manifestFromYAMLInput := *openapiclient.NewManifestFromYAMLInput("ManifestYaml_example") // ManifestFromYAMLInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ValidateStackManifest(context.Background()).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ValidateStackManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateStackManifest`: ManifestValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ValidateStackManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateStackManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manifestFromYAMLInput** | [**ManifestFromYAMLInput**](ManifestFromYAMLInput.md) |  | 

### Return type

[**ManifestValidationResponse**](ManifestValidationResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

