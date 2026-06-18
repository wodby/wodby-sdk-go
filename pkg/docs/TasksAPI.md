# \TasksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTask**](TasksAPI.md#CancelTask) | **Post** /tasks/{id}/cancel | Cancel task
[**GetTask**](TasksAPI.md#GetTask) | **Get** /tasks/{id} | Get task
[**ListTasks**](TasksAPI.md#ListTasks) | **Get** /tasks | List tasks
[**RepeatTask**](TasksAPI.md#RepeatTask) | **Post** /tasks/{id}/repeat | Repeat task



## CancelTask

> OperationResult CancelTask(ctx, id).Execute()

Cancel task

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
	resp, r, err := apiClient.TasksAPI.CancelTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.CancelTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTask`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.CancelTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, id).Execute()

Get task

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
	resp, r, err := apiClient.TasksAPI.GetTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> TasksResponse ListTasks(ctx).Scope(scope).OrgId(orgId).ProjectIds(projectIds).WithoutOrigin(withoutOrigin).Statuses(statuses).Search(search).AppId(appId).AppInstanceId(appInstanceId).StackId(stackId).DatabaseId(databaseId).ClusterId(clusterId).ServiceId(serviceId).IntegrationId(integrationId).ProviderId(providerId).Page(page).PageSize(pageSize).Execute()

List tasks

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
	scope := "scope_example" // string |  (optional)
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)
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
	resp, r, err := apiClient.TasksAPI.ListTasks(context.Background()).Scope(scope).OrgId(orgId).ProjectIds(projectIds).WithoutOrigin(withoutOrigin).Statuses(statuses).Search(search).AppId(appId).AppInstanceId(appInstanceId).StackId(stackId).DatabaseId(databaseId).ClusterId(clusterId).ServiceId(serviceId).IntegrationId(integrationId).ProviderId(providerId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: TasksResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** |  | 
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
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

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepeatTask

> OperationResult RepeatTask(ctx, id).RepeatTaskRequest(repeatTaskRequest).Execute()

Repeat task

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
	repeatTaskRequest := *openapiclient.NewRepeatTaskRequest(false) // RepeatTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.RepeatTask(context.Background(), id).RepeatTaskRequest(repeatTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.RepeatTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepeatTask`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.RepeatTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepeatTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repeatTaskRequest** | [**RepeatTaskRequest**](RepeatTaskRequest.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

