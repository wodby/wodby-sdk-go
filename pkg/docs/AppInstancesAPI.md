# \AppInstancesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppInstancesByNameAppNameInstanceNameGet**](AppInstancesAPI.md#AppInstancesByNameAppNameInstanceNameGet) | **Get** /app-instances/by-name/{appName}/{instanceName} | Get app instance by app and instance name
[**AppInstancesGet**](AppInstancesAPI.md#AppInstancesGet) | **Get** /app-instances | List app instances
[**AppInstancesIdDelete**](AppInstancesAPI.md#AppInstancesIdDelete) | **Delete** /app-instances/{id} | Delete app instance
[**AppInstancesIdGet**](AppInstancesAPI.md#AppInstancesIdGet) | **Get** /app-instances/{id} | Get app instance
[**AppInstancesIdPut**](AppInstancesAPI.md#AppInstancesIdPut) | **Put** /app-instances/{id} | Update app instance
[**AppInstancesPost**](AppInstancesAPI.md#AppInstancesPost) | **Post** /app-instances | Create app instance



## AppInstancesByNameAppNameInstanceNameGet

> AppInstance AppInstancesByNameAppNameInstanceNameGet(ctx, appName, instanceName).OrgId(orgId).Execute()

Get app instance by app and instance name

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
	appName := "appName_example" // string | 
	instanceName := "instanceName_example" // string | 
	orgId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.AppInstancesByNameAppNameInstanceNameGet(context.Background(), appName, instanceName).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.AppInstancesByNameAppNameInstanceNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInstancesByNameAppNameInstanceNameGet`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.AppInstancesByNameAppNameInstanceNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 
**instanceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppInstancesByNameAppNameInstanceNameGetRequest struct via the builder pattern


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


## AppInstancesGet

> []AppInstance AppInstancesGet(ctx).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()

List app instances

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
	appId := int32(56) // int32 |  (optional)
	clusterId := int32(56) // int32 |  (optional)
	clusterApp := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.AppInstancesGet(context.Background()).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.AppInstancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInstancesGet`: []AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.AppInstancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppInstancesGetRequest struct via the builder pattern


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


## AppInstancesIdDelete

> OperationResult AppInstancesIdDelete(ctx, id).Force(force).Execute()

Delete app instance

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
	force := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.AppInstancesIdDelete(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.AppInstancesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInstancesIdDelete`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.AppInstancesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppInstancesIdDeleteRequest struct via the builder pattern


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


## AppInstancesIdGet

> AppInstance AppInstancesIdGet(ctx, id).Execute()

Get app instance

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
	resp, r, err := apiClient.AppInstancesAPI.AppInstancesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.AppInstancesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInstancesIdGet`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.AppInstancesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppInstancesIdGetRequest struct via the builder pattern


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


## AppInstancesIdPut

> AppInstance AppInstancesIdPut(ctx, id).UpdateTitleRequest(updateTitleRequest).Execute()

Update app instance

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
	updateTitleRequest := *openapiclient.NewUpdateTitleRequest("Title_example") // UpdateTitleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.AppInstancesIdPut(context.Background(), id).UpdateTitleRequest(updateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.AppInstancesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInstancesIdPut`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.AppInstancesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppInstancesIdPutRequest struct via the builder pattern


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


## AppInstancesPost

> AppInstance AppInstancesPost(ctx).NewAppInstanceInput(newAppInstanceInput).Execute()

Create app instance

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
	newAppInstanceInput := *openapiclient.NewNewAppInstanceInput(int32(123), "InstanceName_example", "InstanceTitle_example", "Domain_example", int32(123), []openapiclient.NewAppServiceInput{*openapiclient.NewNewAppServiceInput(int32(123), false)}, int32(123)) // NewAppInstanceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.AppInstancesPost(context.Background()).NewAppInstanceInput(newAppInstanceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.AppInstancesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInstancesPost`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.AppInstancesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppInstancesPostRequest struct via the builder pattern


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

