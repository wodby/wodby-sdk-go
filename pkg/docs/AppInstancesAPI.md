# \AppInstancesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppInstance**](AppInstancesAPI.md#CreateAppInstance) | **Post** /app-instances | Create app instance
[**DeleteAppInstance**](AppInstancesAPI.md#DeleteAppInstance) | **Delete** /app-instances/{id} | Delete app instance
[**GetAppInstance**](AppInstancesAPI.md#GetAppInstance) | **Get** /app-instances/{id} | Get app instance
[**GetAppInstanceByName**](AppInstancesAPI.md#GetAppInstanceByName) | **Get** /app-instances/by-name/{appName}/{instanceName} | Get app instance by app and instance name
[**ListAppInstances**](AppInstancesAPI.md#ListAppInstances) | **Get** /app-instances | List app instances
[**UpdateAppInstance**](AppInstancesAPI.md#UpdateAppInstance) | **Put** /app-instances/{id} | Update app instance



## CreateAppInstance

> AppInstance CreateAppInstance(ctx).NewAppInstanceInput(newAppInstanceInput).Execute()

Create app instance

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
	newAppInstanceInput := *openapiclient.NewNewAppInstanceInput(int32(123), "InstanceName_example", "InstanceTitle_example", "Domain_example", int32(123), []openapiclient.CreateAppServiceInput{*openapiclient.NewCreateAppServiceInput(int32(123), false)}, int32(123)) // NewAppInstanceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.CreateAppInstance(context.Background()).NewAppInstanceInput(newAppInstanceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.CreateAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppInstance`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.CreateAppInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppInstanceInput** | [**NewAppInstanceInput**](NewAppInstanceInput.md) |  | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppInstance

> OperationResult DeleteAppInstance(ctx, id).Force(force).Execute()

Delete app instance

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
	force := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.DeleteAppInstance(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.DeleteAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppInstance`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.DeleteAppInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]

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


## GetAppInstance

> AppInstance GetAppInstance(ctx, id).Execute()

Get app instance

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
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstance`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstanceByName

> AppInstance GetAppInstanceByName(ctx, appName, instanceName).OrgId(orgId).Execute()

Get app instance by app and instance name

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
	appName := "appName_example" // string | 
	instanceName := "instanceName_example" // string | 
	orgId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstanceByName(context.Background(), appName, instanceName).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstanceByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstanceByName`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstanceByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 
**instanceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgId** | **int32** |  | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppInstances

> []AppInstance ListAppInstances(ctx).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()

List app instances

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
	orgId := int32(56) // int32 | 
	projectIds := "projectIds_example" // string | Comma-separated project ids (optional)
	appId := int32(56) // int32 |  (optional)
	clusterId := int32(56) // int32 |  (optional)
	clusterApp := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.ListAppInstances(context.Background()).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.ListAppInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppInstances`: []AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.ListAppInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **appId** | **int32** |  | 
 **clusterId** | **int32** |  | 
 **clusterApp** | **bool** |  | 

### Return type

[**[]AppInstance**](AppInstance.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppInstance

> AppInstance UpdateAppInstance(ctx, id).UpdateTitleRequest(updateTitleRequest).Execute()

Update app instance

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
	updateTitleRequest := *openapiclient.NewUpdateTitleRequest("Title_example") // UpdateTitleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.UpdateAppInstance(context.Background(), id).UpdateTitleRequest(updateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpdateAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppInstance`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpdateAppInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTitleRequest** | [**UpdateTitleRequest**](UpdateTitleRequest.md) |  | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

