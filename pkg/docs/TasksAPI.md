# \TasksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksGet**](TasksAPI.md#TasksGet) | **Get** /tasks | List tasks
[**TasksIdCancelPost**](TasksAPI.md#TasksIdCancelPost) | **Post** /tasks/{id}/cancel | Cancel task
[**TasksIdGet**](TasksAPI.md#TasksIdGet) | **Get** /tasks/{id} | Get task
[**TasksIdRepeatPost**](TasksAPI.md#TasksIdRepeatPost) | **Post** /tasks/{id}/repeat | Repeat task



## TasksGet

> TasksResponse TasksGet(ctx).Scope(scope).OrgId(orgId).ProjectIds(projectIds).WithoutOrigin(withoutOrigin).Statuses(statuses).Search(search).AppId(appId).AppInstanceId(appInstanceId).StackId(stackId).DatabaseId(databaseId).ClusterId(clusterId).ServiceId(serviceId).IntegrationId(integrationId).ProviderId(providerId).Page(page).PageSize(pageSize).Execute()

List tasks

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
	scope := "scope_example" // string |  (optional)
	orgId := int32(56) // int32 |  (optional)
	projectIds := "projectIds_example" // string | Comma-separated project ids (optional)
	withoutOrigin := true // bool |  (optional)
	statuses := "statuses_example" // string | Comma-separated task statuses (optional)
	search := "search_example" // string |  (optional)
	appId := int32(56) // int32 |  (optional)
	appInstanceId := int32(56) // int32 |  (optional)
	stackId := int32(56) // int32 |  (optional)
	databaseId := int32(56) // int32 |  (optional)
	clusterId := int32(56) // int32 |  (optional)
	serviceId := int32(56) // int32 |  (optional)
	integrationId := int32(56) // int32 |  (optional)
	providerId := int32(56) // int32 |  (optional)
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksGet(context.Background()).Scope(scope).OrgId(orgId).ProjectIds(projectIds).WithoutOrigin(withoutOrigin).Statuses(statuses).Search(search).AppId(appId).AppInstanceId(appInstanceId).StackId(stackId).DatabaseId(databaseId).ClusterId(clusterId).ServiceId(serviceId).IntegrationId(integrationId).ProviderId(providerId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksGet`: TasksResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** |  | 
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **withoutOrigin** | **bool** |  | 
 **statuses** | **string** | Comma-separated task statuses | 
 **search** | **string** |  | 
 **appId** | **int32** |  | 
 **appInstanceId** | **int32** |  | 
 **stackId** | **int32** |  | 
 **databaseId** | **int32** |  | 
 **clusterId** | **int32** |  | 
 **serviceId** | **int32** |  | 
 **integrationId** | **int32** |  | 
 **providerId** | **int32** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**TasksResponse**](TasksResponse.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdCancelPost

> OperationResult TasksIdCancelPost(ctx, id).Execute()

Cancel task

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
	resp, r, err := apiClient.TasksAPI.TasksIdCancelPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdCancelPost`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdGet

> Task TasksIdGet(ctx, id).Execute()

Get task

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
	resp, r, err := apiClient.TasksAPI.TasksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdGet`: Task
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdRepeatPost

> OperationResult TasksIdRepeatPost(ctx, id).RepeatTaskRequest(repeatTaskRequest).Execute()

Repeat task

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
	repeatTaskRequest := *openapiclient.NewRepeatTaskRequest(false) // RepeatTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksIdRepeatPost(context.Background(), id).RepeatTaskRequest(repeatTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksIdRepeatPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdRepeatPost`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksIdRepeatPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdRepeatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repeatTaskRequest** | [**RepeatTaskRequest**](RepeatTaskRequest.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

