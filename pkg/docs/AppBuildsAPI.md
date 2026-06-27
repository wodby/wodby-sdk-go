# \AppBuildsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppBuild**](AppBuildsAPI.md#CreateAppBuild) | **Post** /app-builds | Create build
[**CreateAppBuildFromCi**](AppBuildsAPI.md#CreateAppBuildFromCi) | **Post** /app-builds/from-ci | Create build from CI
[**DeployAppBuild**](AppBuildsAPI.md#DeployAppBuild) | **Post** /app-builds/{id}/deploy | Deploy build
[**GetAppBuild**](AppBuildsAPI.md#GetAppBuild) | **Get** /app-builds/{id} | Get build
[**GetAppBuildConfig**](AppBuildsAPI.md#GetAppBuildConfig) | **Get** /app-builds/{id}/config | Get build config
[**GetAppBuildDockerRegistryCredentials**](AppBuildsAPI.md#GetAppBuildDockerRegistryCredentials) | **Get** /app-builds/{id}/docker-registry-credentials | Get Docker registry credentials for build
[**ListAppBuilds**](AppBuildsAPI.md#ListAppBuilds) | **Get** /app-builds | List app builds



## CreateAppBuild

> []AppBuild CreateAppBuild(ctx).CreateBuildRequest(createBuildRequest).Execute()

Create build

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
	createBuildRequest := *openapiclient.NewCreateBuildRequest([]int32{int32(123)}) // CreateBuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppBuildsAPI.CreateAppBuild(context.Background()).CreateBuildRequest(createBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.CreateAppBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppBuild`: []AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.CreateAppBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuildRequest** | [**CreateBuildRequest**](CreateBuildRequest.md) |  | 

### Return type

[**[]AppBuild**](AppBuild.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppBuildFromCi

> AppBuild CreateAppBuildFromCi(ctx).NewBuildFromCIInput(newBuildFromCIInput).Execute()

Create build from CI

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
	newBuildFromCIInput := *openapiclient.NewNewBuildFromCIInput(int32(123), "GitCommitSHA_example", "GitRef_example", "GitRefType_example", int32(123), "BuildId_example", "Provider_example") // NewBuildFromCIInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppBuildsAPI.CreateAppBuildFromCi(context.Background()).NewBuildFromCIInput(newBuildFromCIInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.CreateAppBuildFromCi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppBuildFromCi`: AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.CreateAppBuildFromCi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppBuildFromCiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBuildFromCIInput** | [**NewBuildFromCIInput**](NewBuildFromCIInput.md) |  | 

### Return type

[**AppBuild**](AppBuild.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployAppBuild

> AppDeployment DeployAppBuild(ctx, id).Execute()

Deploy build

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
	resp, r, err := apiClient.AppBuildsAPI.DeployAppBuild(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.DeployAppBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployAppBuild`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.DeployAppBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployAppBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDeployment**](AppDeployment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppBuild

> AppBuild GetAppBuild(ctx, id).Execute()

Get build

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
	resp, r, err := apiClient.AppBuildsAPI.GetAppBuild(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.GetAppBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppBuild`: AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.GetAppBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppBuild**](AppBuild.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [ciAccessTokenHeader](../README.md#ciAccessTokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppBuildConfig

> AppBuildConfig GetAppBuildConfig(ctx, id).Execute()

Get build config

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
	resp, r, err := apiClient.AppBuildsAPI.GetAppBuildConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.GetAppBuildConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppBuildConfig`: AppBuildConfig
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.GetAppBuildConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppBuildConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppBuildConfig**](AppBuildConfig.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [ciAccessTokenHeader](../README.md#ciAccessTokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppBuildDockerRegistryCredentials

> DockerRegistryCredentials GetAppBuildDockerRegistryCredentials(ctx, id).Execute()

Get Docker registry credentials for build

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
	resp, r, err := apiClient.AppBuildsAPI.GetAppBuildDockerRegistryCredentials(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.GetAppBuildDockerRegistryCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppBuildDockerRegistryCredentials`: DockerRegistryCredentials
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.GetAppBuildDockerRegistryCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppBuildDockerRegistryCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerRegistryCredentials**](DockerRegistryCredentials.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [ciAccessTokenHeader](../README.md#ciAccessTokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppBuilds

> AppBuildsResponse ListAppBuilds(ctx).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()

List app builds

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
	resp, r, err := apiClient.AppBuildsAPI.ListAppBuilds(context.Background()).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.ListAppBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppBuilds`: AppBuildsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.ListAppBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**AppBuildsResponse**](AppBuildsResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

