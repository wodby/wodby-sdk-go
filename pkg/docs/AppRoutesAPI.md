# \AppRoutesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppRoute**](AppRoutesAPI.md#CreateAppRoute) | **Post** /app-routes | Create app route
[**DeleteAppRoute**](AppRoutesAPI.md#DeleteAppRoute) | **Delete** /app-routes/{id} | Delete app route
[**GetAppRoute**](AppRoutesAPI.md#GetAppRoute) | **Get** /app-routes/{id} | Get app route
[**ListAppRoutes**](AppRoutesAPI.md#ListAppRoutes) | **Get** /app-routes | List app routes
[**UpdateAppRoute**](AppRoutesAPI.md#UpdateAppRoute) | **Put** /app-routes/{id} | Update app route



## CreateAppRoute

> AppRoute CreateAppRoute(ctx).NewAppRouteInput(newAppRouteInput).Execute()

Create app route



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
	newAppRouteInput := *openapiclient.NewNewAppRouteInput(int32(123), false, false, int32(123), "Host_example") // NewAppRouteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRoutesAPI.CreateAppRoute(context.Background()).NewAppRouteInput(newAppRouteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.CreateAppRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppRoute`: AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.CreateAppRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppRouteInput** | [**NewAppRouteInput**](NewAppRouteInput.md) |  | 

### Return type

[**AppRoute**](AppRoute.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppRoute

> OperationResult DeleteAppRoute(ctx, id).Execute()

Delete app route



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
	resp, r, err := apiClient.AppRoutesAPI.DeleteAppRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.DeleteAppRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppRoute`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.DeleteAppRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRouteRequest struct via the builder pattern


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


## GetAppRoute

> AppRoute GetAppRoute(ctx, id).Execute()

Get app route



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
	resp, r, err := apiClient.AppRoutesAPI.GetAppRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.GetAppRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppRoute`: AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.GetAppRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppRoute**](AppRoute.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppRoutes

> []AppRoute ListAppRoutes(ctx).AppInstanceId(appInstanceId).Execute()

List app routes



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
	appInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRoutesAPI.ListAppRoutes(context.Background()).AppInstanceId(appInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.ListAppRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppRoutes`: []AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.ListAppRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 

### Return type

[**[]AppRoute**](AppRoute.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppRoute

> AppRoute UpdateAppRoute(ctx, id).UpdateAppRouteInput(updateAppRouteInput).Execute()

Update app route



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
	updateAppRouteInput := *openapiclient.NewUpdateAppRouteInput() // UpdateAppRouteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRoutesAPI.UpdateAppRoute(context.Background(), id).UpdateAppRouteInput(updateAppRouteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRoutesAPI.UpdateAppRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppRoute`: AppRoute
	fmt.Fprintf(os.Stdout, "Response from `AppRoutesAPI.UpdateAppRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppRouteInput** | [**UpdateAppRouteInput**](UpdateAppRouteInput.md) |  | 

### Return type

[**AppRoute**](AppRoute.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

