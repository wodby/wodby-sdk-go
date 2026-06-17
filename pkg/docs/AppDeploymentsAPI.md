# \AppDeploymentsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppDeploymentsFromCiPost**](AppDeploymentsAPI.md#AppDeploymentsFromCiPost) | **Post** /app-deployments/from-ci | Create deployment from CI
[**AppDeploymentsGet**](AppDeploymentsAPI.md#AppDeploymentsGet) | **Get** /app-deployments | List app deployments
[**AppDeploymentsIdGet**](AppDeploymentsAPI.md#AppDeploymentsIdGet) | **Get** /app-deployments/{id} | Get deployment
[**AppDeploymentsIdRedeployPost**](AppDeploymentsAPI.md#AppDeploymentsIdRedeployPost) | **Post** /app-deployments/{id}/redeploy | Redeploy deployment
[**AppDeploymentsPost**](AppDeploymentsAPI.md#AppDeploymentsPost) | **Post** /app-deployments | Create deployment



## AppDeploymentsFromCiPost

> AppDeployment AppDeploymentsFromCiPost(ctx).DeploymentFromCIInput(deploymentFromCIInput).Execute()

Create deployment from CI

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
	deploymentFromCIInput := *openapiclient.NewDeploymentFromCIInput(int32(123), []openapiclient.ServiceDeploymentInput{*openapiclient.NewServiceDeploymentInput("Name_example", "Image_example")}, false) // DeploymentFromCIInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDeploymentsAPI.AppDeploymentsFromCiPost(context.Background()).DeploymentFromCIInput(deploymentFromCIInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.AppDeploymentsFromCiPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDeploymentsFromCiPost`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.AppDeploymentsFromCiPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDeploymentsFromCiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentFromCIInput** | [**DeploymentFromCIInput**](DeploymentFromCIInput.md) |  | 

### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDeploymentsGet

> AppDeploymentsResponse AppDeploymentsGet(ctx).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()

List app deployments

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
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDeploymentsAPI.AppDeploymentsGet(context.Background()).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.AppDeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDeploymentsGet`: AppDeploymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.AppDeploymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDeploymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**AppDeploymentsResponse**](AppDeploymentsResponse.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDeploymentsIdGet

> AppDeployment AppDeploymentsIdGet(ctx, id).Execute()

Get deployment

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
	resp, r, err := apiClient.AppDeploymentsAPI.AppDeploymentsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.AppDeploymentsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDeploymentsIdGet`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.AppDeploymentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDeploymentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDeploymentsIdRedeployPost

> AppDeployment AppDeploymentsIdRedeployPost(ctx, id).Execute()

Redeploy deployment

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
	resp, r, err := apiClient.AppDeploymentsAPI.AppDeploymentsIdRedeployPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.AppDeploymentsIdRedeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDeploymentsIdRedeployPost`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.AppDeploymentsIdRedeployPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDeploymentsIdRedeployPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDeploymentsPost

> AppDeployment AppDeploymentsPost(ctx).CreateDeploymentRequest(createDeploymentRequest).Execute()

Create deployment

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
	createDeploymentRequest := *openapiclient.NewCreateDeploymentRequest([]openapiclient.AppServiceDeploymentRequest{*openapiclient.NewAppServiceDeploymentRequest(int32(123), false)}) // CreateDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDeploymentsAPI.AppDeploymentsPost(context.Background()).CreateDeploymentRequest(createDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.AppDeploymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDeploymentsPost`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.AppDeploymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDeploymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeploymentRequest** | [**CreateDeploymentRequest**](CreateDeploymentRequest.md) |  | 

### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

