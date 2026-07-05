# \AppDeploymentsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppDeployment**](AppDeploymentsAPI.md#CreateAppDeployment) | **Post** /app-deployments | Create deployment
[**CreateAppDeploymentFromCi**](AppDeploymentsAPI.md#CreateAppDeploymentFromCi) | **Post** /app-deployments/from-ci | Create deployment from CI
[**GetAppDeployment**](AppDeploymentsAPI.md#GetAppDeployment) | **Get** /app-deployments/{id} | Get deployment
[**ListAppDeployments**](AppDeploymentsAPI.md#ListAppDeployments) | **Get** /app-deployments | List app deployments
[**RedeployAppDeployment**](AppDeploymentsAPI.md#RedeployAppDeployment) | **Post** /app-deployments/{id}/redeploy | Redeploy deployment



## CreateAppDeployment

> AppDeployment CreateAppDeployment(ctx).CreateDeploymentRequest(createDeploymentRequest).Execute()

Create deployment



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
	createDeploymentRequest := *openapiclient.NewCreateDeploymentRequest([]openapiclient.AppServiceDeploymentRequest{*openapiclient.NewAppServiceDeploymentRequest(int32(123), false)}) // CreateDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDeploymentsAPI.CreateAppDeployment(context.Background()).CreateDeploymentRequest(createDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.CreateAppDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppDeployment`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.CreateAppDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeploymentRequest** | [**CreateDeploymentRequest**](CreateDeploymentRequest.md) |  | 

### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppDeploymentFromCi

> AppDeployment CreateAppDeploymentFromCi(ctx).DeploymentFromCIInput(deploymentFromCIInput).Execute()

Create deployment from CI



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
	deploymentFromCIInput := *openapiclient.NewDeploymentFromCIInput(int32(123), []openapiclient.ServiceDeploymentInput{*openapiclient.NewServiceDeploymentInput("Name_example", "Image_example")}, false) // DeploymentFromCIInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDeploymentsAPI.CreateAppDeploymentFromCi(context.Background()).DeploymentFromCIInput(deploymentFromCIInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.CreateAppDeploymentFromCi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppDeploymentFromCi`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.CreateAppDeploymentFromCi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppDeploymentFromCiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentFromCIInput** | [**DeploymentFromCIInput**](DeploymentFromCIInput.md) |  | 

### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [ciAccessTokenHeader](../README.md#ciAccessTokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDeployment

> AppDeployment GetAppDeployment(ctx, id).Execute()

Get deployment



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
	resp, r, err := apiClient.AppDeploymentsAPI.GetAppDeployment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.GetAppDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppDeployment`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.GetAppDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppDeployments

> AppDeploymentsResponse ListAppDeployments(ctx).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()

List app deployments



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
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDeploymentsAPI.ListAppDeployments(context.Background()).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.ListAppDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppDeployments`: AppDeploymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.ListAppDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**AppDeploymentsResponse**](AppDeploymentsResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployAppDeployment

> AppDeployment RedeployAppDeployment(ctx, id).Execute()

Redeploy deployment



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
	resp, r, err := apiClient.AppDeploymentsAPI.RedeployAppDeployment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDeploymentsAPI.RedeployAppDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployAppDeployment`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppDeploymentsAPI.RedeployAppDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployAppDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

