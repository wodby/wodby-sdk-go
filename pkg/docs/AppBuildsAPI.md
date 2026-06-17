# \AppBuildsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppBuildsFromCiPost**](AppBuildsAPI.md#AppBuildsFromCiPost) | **Post** /app-builds/from-ci | Create build from CI
[**AppBuildsGet**](AppBuildsAPI.md#AppBuildsGet) | **Get** /app-builds | List app builds
[**AppBuildsIdDeployPost**](AppBuildsAPI.md#AppBuildsIdDeployPost) | **Post** /app-builds/{id}/deploy | Deploy build
[**AppBuildsIdDockerRegistryCredentialsGet**](AppBuildsAPI.md#AppBuildsIdDockerRegistryCredentialsGet) | **Get** /app-builds/{id}/docker-registry-credentials | Get Docker registry credentials for build
[**AppBuildsIdGet**](AppBuildsAPI.md#AppBuildsIdGet) | **Get** /app-builds/{id} | Get build
[**AppBuildsIdVoidPost**](AppBuildsAPI.md#AppBuildsIdVoidPost) | **Post** /app-builds/{id}/void | Void build images
[**AppBuildsPost**](AppBuildsAPI.md#AppBuildsPost) | **Post** /app-builds | Create build



## AppBuildsFromCiPost

> AppBuild AppBuildsFromCiPost(ctx).NewBuildFromCIInput(newBuildFromCIInput).Execute()

Create build from CI

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
	newBuildFromCIInput := *openapiclient.NewNewBuildFromCIInput(int32(123), "GitCommitSHA_example", "GitRef_example", "GitRefType_example", int32(123), "BuildID_example", "Provider_example") // NewBuildFromCIInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsFromCiPost(context.Background()).NewBuildFromCIInput(newBuildFromCIInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsFromCiPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsFromCiPost`: AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsFromCiPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsFromCiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBuildFromCIInput** | [**NewBuildFromCIInput**](NewBuildFromCIInput.md) |  | 

### Return type

[**AppBuild**](AppBuild.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppBuildsGet

> AppBuildsResponse AppBuildsGet(ctx).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()

List app builds

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
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsGet(context.Background()).AppInstanceId(appInstanceId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsGet`: AppBuildsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**AppBuildsResponse**](AppBuildsResponse.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppBuildsIdDeployPost

> AppDeployment AppBuildsIdDeployPost(ctx, id).Execute()

Deploy build

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
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsIdDeployPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsIdDeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsIdDeployPost`: AppDeployment
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsIdDeployPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsIdDeployPostRequest struct via the builder pattern


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


## AppBuildsIdDockerRegistryCredentialsGet

> DockerRegistryCredentials AppBuildsIdDockerRegistryCredentialsGet(ctx, id).Execute()

Get Docker registry credentials for build

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
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsIdDockerRegistryCredentialsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsIdDockerRegistryCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsIdDockerRegistryCredentialsGet`: DockerRegistryCredentials
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsIdDockerRegistryCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsIdDockerRegistryCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerRegistryCredentials**](DockerRegistryCredentials.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppBuildsIdGet

> AppBuild AppBuildsIdGet(ctx, id).Execute()

Get build

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
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsIdGet`: AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppBuild**](AppBuild.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppBuildsIdVoidPost

> AppBuild AppBuildsIdVoidPost(ctx, id).Execute()

Void build images

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
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsIdVoidPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsIdVoidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsIdVoidPost`: AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsIdVoidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsIdVoidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppBuild**](AppBuild.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppBuildsPost

> []AppBuild AppBuildsPost(ctx).CreateBuildRequest(createBuildRequest).Execute()

Create build

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
	createBuildRequest := *openapiclient.NewCreateBuildRequest() // CreateBuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppBuildsAPI.AppBuildsPost(context.Background()).CreateBuildRequest(createBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuildsAPI.AppBuildsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppBuildsPost`: []AppBuild
	fmt.Fprintf(os.Stdout, "Response from `AppBuildsAPI.AppBuildsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppBuildsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuildRequest** | [**CreateBuildRequest**](CreateBuildRequest.md) |  | 

### Return type

[**[]AppBuild**](AppBuild.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

