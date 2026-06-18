# \ClustersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersAPI.md#CreateCluster) | **Post** /clusters | Create cluster
[**DeleteCluster**](ClustersAPI.md#DeleteCluster) | **Delete** /clusters/{id} | Delete cluster
[**GetCluster**](ClustersAPI.md#GetCluster) | **Get** /clusters/{id} | Get cluster
[**GetClusterByName**](ClustersAPI.md#GetClusterByName) | **Get** /clusters/by-name/{name} | Get cluster by name
[**ListClusters**](ClustersAPI.md#ListClusters) | **Get** /clusters | List clusters
[**UpdateCluster**](ClustersAPI.md#UpdateCluster) | **Put** /clusters/{id} | Update cluster



## CreateCluster

> Cluster CreateCluster(ctx).NewClusterInput(newClusterInput).Execute()

Create cluster

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
	newClusterInput := *openapiclient.NewNewClusterInput(int32(123), int32(123), "Name_example", "Title_example", false, false) // NewClusterInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.CreateCluster(context.Background()).NewClusterInput(newClusterInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newClusterInput** | [**NewClusterInput**](NewClusterInput.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> OperationResult DeleteCluster(ctx, id).Force(force).Execute()

Delete cluster

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
	resp, r, err := apiClient.ClustersAPI.DeleteCluster(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCluster`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


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


## GetCluster

> Cluster GetCluster(ctx, id).Execute()

Get cluster

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
	resp, r, err := apiClient.ClustersAPI.GetCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterByName

> Cluster GetClusterByName(ctx, name).OrgId(orgId).Execute()

Get cluster by name

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
	orgId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterByName(context.Background(), name).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterByName`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> []Cluster ListClusters(ctx).OrgId(orgId).ProjectIds(projectIds).IntegrationId(integrationId).Execute()

List clusters

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
	integrationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListClusters(context.Background()).OrgId(orgId).ProjectIds(projectIds).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: []Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **integrationId** | **int32** |  | 

### Return type

[**[]Cluster**](Cluster.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx, id).UpdateTitleRequest(updateTitleRequest).Execute()

Update cluster

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
	resp, r, err := apiClient.ClustersAPI.UpdateCluster(context.Background(), id).UpdateTitleRequest(updateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.UpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTitleRequest** | [**UpdateTitleRequest**](UpdateTitleRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

