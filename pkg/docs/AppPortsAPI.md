# \AppPortsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppPortSetting**](AppPortsAPI.md#DeleteAppPortSetting) | **Delete** /app-ports/{id}/settings/{name} | Delete app port setting
[**GetAppEndpoint**](AppPortsAPI.md#GetAppEndpoint) | **Get** /app-endpoints/{id} | Get app endpoint
[**GetAppPort**](AppPortsAPI.md#GetAppPort) | **Get** /app-ports/{id} | Get app port
[**ListAppEndpoints**](AppPortsAPI.md#ListAppEndpoints) | **Get** /app-endpoints | List app endpoints
[**ListAppPortSettings**](AppPortsAPI.md#ListAppPortSettings) | **Get** /app-ports/{id}/settings | List app port settings
[**ListAppPorts**](AppPortsAPI.md#ListAppPorts) | **Get** /app-ports | List app ports
[**PublishAppPort**](AppPortsAPI.md#PublishAppPort) | **Post** /app-ports/{id}/actions/publish | Publish app port
[**ReplaceAppPortSettings**](AppPortsAPI.md#ReplaceAppPortSettings) | **Put** /app-ports/{id}/settings | Replace app port settings
[**SetAppPortSetting**](AppPortsAPI.md#SetAppPortSetting) | **Put** /app-ports/{id}/settings/{name} | Set app port setting
[**UnpublishAppPort**](AppPortsAPI.md#UnpublishAppPort) | **Post** /app-ports/{id}/actions/unpublish | Unpublish app port
[**UpdateAppPort**](AppPortsAPI.md#UpdateAppPort) | **Put** /app-ports/{id} | Update app port



## DeleteAppPortSetting

> OperationResult DeleteAppPortSetting(ctx, id, name).Execute()

Delete app port setting



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
	name := openapiclient.AppPortSettingName("IDLE_TIMEOUT") // AppPortSettingName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPortsAPI.DeleteAppPortSetting(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.DeleteAppPortSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppPortSetting`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.DeleteAppPortSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | [**AppPortSettingName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppPortSettingRequest struct via the builder pattern


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


## GetAppEndpoint

> AppEndpoint GetAppEndpoint(ctx, id).Execute()

Get app endpoint



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
	resp, r, err := apiClient.AppPortsAPI.GetAppEndpoint(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.GetAppEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEndpoint`: AppEndpoint
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.GetAppEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEndpoint**](AppEndpoint.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppEndpoints

> []AppEndpoint ListAppEndpoints(ctx).AppEnvironmentId(appEnvironmentId).Execute()

List app endpoints



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
	appEnvironmentId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPortsAPI.ListAppEndpoints(context.Background()).AppEnvironmentId(appEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.ListAppEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppEndpoints`: []AppEndpoint
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.ListAppEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appEnvironmentId** | **int32** |  | 

### Return type

[**[]AppEndpoint**](AppEndpoint.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppPortSettings

> []AppPortSetting ListAppPortSettings(ctx, id).Execute()

List app port settings



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
	resp, r, err := apiClient.AppPortsAPI.ListAppPortSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.ListAppPortSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppPortSettings`: []AppPortSetting
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.ListAppPortSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppPortSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppPortSetting**](AppPortSetting.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishAppPort

> AppPort PublishAppPort(ctx, id).UpdateAppPortInput(updateAppPortInput).Execute()

Publish app port



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
	updateAppPortInput := *openapiclient.NewUpdateAppPortInput() // UpdateAppPortInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPortsAPI.PublishAppPort(context.Background(), id).UpdateAppPortInput(updateAppPortInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.PublishAppPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishAppPort`: AppPort
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.PublishAppPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishAppPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppPortInput** | [**UpdateAppPortInput**](UpdateAppPortInput.md) |  | 

### Return type

[**AppPort**](AppPort.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAppPortSettings

> []AppPortSetting ReplaceAppPortSettings(ctx, id).AppPortSettingsInput(appPortSettingsInput).Execute()

Replace app port settings



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
	appPortSettingsInput := *openapiclient.NewAppPortSettingsInput([]openapiclient.AppPortSettingInput{*openapiclient.NewAppPortSettingInput(openapiclient.AppPortSettingName("IDLE_TIMEOUT"), "Value_example")}) // AppPortSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPortsAPI.ReplaceAppPortSettings(context.Background(), id).AppPortSettingsInput(appPortSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.ReplaceAppPortSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceAppPortSettings`: []AppPortSetting
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.ReplaceAppPortSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAppPortSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appPortSettingsInput** | [**AppPortSettingsInput**](AppPortSettingsInput.md) |  | 

### Return type

[**[]AppPortSetting**](AppPortSetting.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAppPortSetting

> AppPortSetting SetAppPortSetting(ctx, id, name).SetStringValueInput(setStringValueInput).Execute()

Set app port setting



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
	name := openapiclient.AppPortSettingName("IDLE_TIMEOUT") // AppPortSettingName | 
	setStringValueInput := *openapiclient.NewSetStringValueInput("Value_example") // SetStringValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPortsAPI.SetAppPortSetting(context.Background(), id, name).SetStringValueInput(setStringValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.SetAppPortSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAppPortSetting`: AppPortSetting
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.SetAppPortSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | [**AppPortSettingName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAppPortSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setStringValueInput** | [**SetStringValueInput**](SetStringValueInput.md) |  | 

### Return type

[**AppPortSetting**](AppPortSetting.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishAppPort

> AppPort UnpublishAppPort(ctx, id).Execute()

Unpublish app port



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
	resp, r, err := apiClient.AppPortsAPI.UnpublishAppPort(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.UnpublishAppPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpublishAppPort`: AppPort
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.UnpublishAppPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishAppPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppPort**](AppPort.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppPort

> AppPort UpdateAppPort(ctx, id).UpdateAppPortInput(updateAppPortInput).Execute()

Update app port



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
	updateAppPortInput := *openapiclient.NewUpdateAppPortInput() // UpdateAppPortInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPortsAPI.UpdateAppPort(context.Background(), id).UpdateAppPortInput(updateAppPortInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPortsAPI.UpdateAppPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppPort`: AppPort
	fmt.Fprintf(os.Stdout, "Response from `AppPortsAPI.UpdateAppPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppPortInput** | [**UpdateAppPortInput**](UpdateAppPortInput.md) |  | 

### Return type

[**AppPort**](AppPort.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

