# \ServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetService**](ServicesAPI.md#GetService) | **Get** /services/{id} | Get service
[**GetServiceByName**](ServicesAPI.md#GetServiceByName) | **Get** /services/by-name/{name} | Get service by name
[**GetServiceRevision**](ServicesAPI.md#GetServiceRevision) | **Get** /service-revisions/{id} | Get service revision
[**ListServiceLinkCandidates**](ServicesAPI.md#ListServiceLinkCandidates) | **Get** /services/{name}/options/link-candidates | List service link candidates
[**ListServices**](ServicesAPI.md#ListServices) | **Get** /services | List services



## GetService

> Service GetService(ctx, id).Execute()

Get service

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
	resp, r, err := apiClient.ServicesAPI.GetService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Service**](Service.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceByName

> Service GetServiceByName(ctx, name).RevNumber(revNumber).Execute()

Get service by name

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
	resp, r, err := apiClient.ServicesAPI.GetServiceByName(context.Background(), name).RevNumber(revNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetServiceByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceByName`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetServiceByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revNumber** | **int32** |  | 

### Return type

[**Service**](Service.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceRevision

> ServiceRevision GetServiceRevision(ctx, id).Execute()

Get service revision

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
	resp, r, err := apiClient.ServicesAPI.GetServiceRevision(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetServiceRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceRevision`: ServiceRevision
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetServiceRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceRevision**](ServiceRevision.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceLinkCandidates

> []map[string]interface{} ListServiceLinkCandidates(ctx, name).Execute()

List service link candidates

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
	resp, r, err := apiClient.ServicesAPI.ListServiceLinkCandidates(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ListServiceLinkCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceLinkCandidates`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ListServiceLinkCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceLinkCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ServicesResponse ListServices(ctx).OrgId(orgId).ProjectIds(projectIds).Search(search).Page(page).PageSize(pageSize).Execute()

List services

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
	resp, r, err := apiClient.ServicesAPI.ListServices(context.Background()).OrgId(orgId).ProjectIds(projectIds).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ListServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServices`: ServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ListServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **search** | **string** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**ServicesResponse**](ServicesResponse.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

