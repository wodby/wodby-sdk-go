# \AppPortsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppPort**](AppPortsAPI.md#GetAppPort) | **Get** /app-ports/{id} | Get app port
[**ListAppPorts**](AppPortsAPI.md#ListAppPorts) | **Get** /app-ports | List app ports



## GetAppPort

> AppPort GetAppPort(ctx, id).Execute()

Get app port



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
	resp, r, err := apiClient.AppPortsAPI.GetAppPort(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.GetAppPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppPort`: AppPort
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.GetAppPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppPort**](AppPort.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppPorts

> []AppPort ListAppPorts(ctx).AppInstanceId(appInstanceId).Execute()

List app ports



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
	resp, r, err := apiClient.AppPortsAPI.ListAppPorts(context.Background()).AppInstanceId(appInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.ListAppPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppPorts`: []AppPort
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.ListAppPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 

### Return type

[**[]AppPort**](AppPort.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

