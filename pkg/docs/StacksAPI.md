# \StacksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStack**](StacksAPI.md#GetStack) | **Get** /stacks/{id} | Get stack
[**GetStackByName**](StacksAPI.md#GetStackByName) | **Get** /stacks/by-name/{name} | Get stack by name
[**GetStackRevision**](StacksAPI.md#GetStackRevision) | **Get** /stack-revisions/{id} | Get stack revision
[**ListStackRevisionServices**](StacksAPI.md#ListStackRevisionServices) | **Get** /stack-revisions/{id}/services | List stack services
[**ListStacks**](StacksAPI.md#ListStacks) | **Get** /stacks | List stacks



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := int32(56) // int32 | 
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
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **search** | **string** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**StacksResponse**](StacksResponse.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

