# \AppRoutesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppRoutesGet**](AppRoutesAPI.md#AppRoutesGet) | **Get** /app-routes | List app routes
[**AppRoutesIdDelete**](AppRoutesAPI.md#AppRoutesIdDelete) | **Delete** /app-routes/{id} | Delete app route
[**AppRoutesIdGet**](AppRoutesAPI.md#AppRoutesIdGet) | **Get** /app-routes/{id} | Get app route
[**AppRoutesIdPut**](AppRoutesAPI.md#AppRoutesIdPut) | **Put** /app-routes/{id} | Update app route
[**AppRoutesPost**](AppRoutesAPI.md#AppRoutesPost) | **Post** /app-routes | Create app route



## AppRoutesGet

> []AppRoute AppRoutesGet(ctx).AppInstanceId(appInstanceId).Execute()

List app routes

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
	appInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRoutesAPI.AppRoutesGet(context.Background()).AppInstanceId(appInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.AppRoutesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppRoutesGet`: []AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.AppRoutesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppRoutesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 

### Return type

[**[]AppRoute**](AppRoute.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRoutesIdDelete

> OperationResult AppRoutesIdDelete(ctx, id).Execute()

Delete app route

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
	resp, r, err := apiClient.AppRoutesAPI.AppRoutesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.AppRoutesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppRoutesIdDelete`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.AppRoutesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRoutesIdDeleteRequest struct via the builder pattern


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


## AppRoutesIdGet

> AppRoute AppRoutesIdGet(ctx, id).Execute()

Get app route

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
	resp, r, err := apiClient.AppRoutesAPI.AppRoutesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.AppRoutesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppRoutesIdGet`: AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.AppRoutesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRoutesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppRoute**](AppRoute.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRoutesIdPut

> AppRoute AppRoutesIdPut(ctx, id).UpdateAppRouteInput(updateAppRouteInput).Execute()

Update app route

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
	updateAppRouteInput := *openapiclient.NewUpdateAppRouteInput() // UpdateAppRouteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRoutesAPI.AppRoutesIdPut(context.Background(), id).UpdateAppRouteInput(updateAppRouteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.AppRoutesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppRoutesIdPut`: AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.AppRoutesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRoutesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppRouteInput** | [**UpdateAppRouteInput**](UpdateAppRouteInput.md) |  | 

### Return type

[**AppRoute**](AppRoute.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRoutesPost

> AppRoute AppRoutesPost(ctx).NewAppRouteInput(newAppRouteInput).Execute()

Create app route

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
	newAppRouteInput := *openapiclient.NewNewAppRouteInput(int32(123), false, false, int32(123), "Host_example") // NewAppRouteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRoutesAPI.AppRoutesPost(context.Background()).NewAppRouteInput(newAppRouteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.AppRoutesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppRoutesPost`: AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.AppRoutesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppRoutesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppRouteInput** | [**NewAppRouteInput**](NewAppRouteInput.md) |  | 

### Return type

[**AppRoute**](AppRoute.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

