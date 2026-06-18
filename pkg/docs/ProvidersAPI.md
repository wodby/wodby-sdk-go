# \ProvidersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProvider**](ProvidersAPI.md#GetProvider) | **Get** /providers/{id} | Get provider
[**GetProviderByName**](ProvidersAPI.md#GetProviderByName) | **Get** /providers/by-name/{name} | Get provider by name
[**GetProviderRevision**](ProvidersAPI.md#GetProviderRevision) | **Get** /provider-revisions/{id} | Get provider revision
[**ListProviders**](ProvidersAPI.md#ListProviders) | **Get** /providers | List providers



## GetProvider

> Provider GetProvider(ctx, id).Execute()

Get provider

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
	resp, r, err := apiClient.ProvidersAPI.GetProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Provider**](Provider.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderByName

> Provider GetProviderByName(ctx, name).Execute()

Get provider by name

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.GetProviderByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProviderByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderByName`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProviderByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Provider**](Provider.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderRevision

> ProviderRevision GetProviderRevision(ctx, id).Execute()

Get provider revision

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
	resp, r, err := apiClient.ProvidersAPI.GetProviderRevision(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProviderRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderRevision`: ProviderRevision
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProviderRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderRevision**](ProviderRevision.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProviders

> ProvidersResponse ListProviders(ctx).OrgId(orgId).ProjectIds(projectIds).ExcludePublic(excludePublic).Search(search).Page(page).PageSize(pageSize).Execute()

List providers

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
	excludePublic := true // bool |  (optional)
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ListProviders(context.Background()).OrgId(orgId).ProjectIds(projectIds).ExcludePublic(excludePublic).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: ProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **excludePublic** | **bool** |  | 
 **search** | **string** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**ProvidersResponse**](ProvidersResponse.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

