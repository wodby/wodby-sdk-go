# \AppServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppService**](AppServicesAPI.md#GetAppService) | **Get** /app-services/{id} | Get app service
[**ListAppServices**](AppServicesAPI.md#ListAppServices) | **Get** /app-services | List app services
[**RunAppServiceAction**](AppServicesAPI.md#RunAppServiceAction) | **Post** /app-services/{id}/actions/{name} | Run app service action
[**UpdateAppService**](AppServicesAPI.md#UpdateAppService) | **Put** /app-services/{id} | Update app service



## GetAppService

> AppService GetAppService(ctx, id).Execute()

Get app service



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
	resp, r, err := apiClient.AppServicesAPI.GetAppService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.GetAppService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppService`: AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.GetAppService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServices

> []AppService ListAppServices(ctx).AppInstanceId(appInstanceId).Execute()

List app services



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServices(context.Background()).AppInstanceId(appInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServices`: []AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 

### Return type

[**[]AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAppServiceAction

> OperationResult RunAppServiceAction(ctx, id, name).Execute()

Run app service action



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.RunAppServiceAction(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.RunAppServiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunAppServiceAction`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.RunAppServiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunAppServiceActionRequest struct via the builder pattern


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


## UpdateAppService

> AppService UpdateAppService(ctx, id).AppServiceInput(appServiceInput).Execute()

Update app service



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
	appServiceInput := *openapiclient.NewAppServiceInput() // AppServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppService(context.Background(), id).AppServiceInput(appServiceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppService`: AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appServiceInput** | [**AppServiceInput**](AppServiceInput.md) |  | 

### Return type

[**AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

