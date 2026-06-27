# \EnvsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnv**](EnvsAPI.md#CreateEnv) | **Post** /envs | Create env
[**DeleteEnv**](EnvsAPI.md#DeleteEnv) | **Delete** /envs/{id} | Delete env
[**GetEnv**](EnvsAPI.md#GetEnv) | **Get** /envs/{id} | Get env
[**GetEnvByName**](EnvsAPI.md#GetEnvByName) | **Get** /envs/by-name/{name} | Get env by name
[**ListEnvs**](EnvsAPI.md#ListEnvs) | **Get** /envs | List envs
[**UpdateEnv**](EnvsAPI.md#UpdateEnv) | **Put** /envs/{id} | Update env



## CreateEnv

> Env CreateEnv(ctx).CreateEnvRequest(createEnvRequest).Execute()

Create env



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
	createEnvRequest := *openapiclient.NewCreateEnvRequest("Name_example", "Title_example", "Type_example") // CreateEnvRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvsAPI.CreateEnv(context.Background()).CreateEnvRequest(createEnvRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvsAPI.CreateEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnv`: Env
	fmt.Fprintf(os.Stdout, "Response from `EnvsAPI.CreateEnv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEnvRequest** | [**CreateEnvRequest**](CreateEnvRequest.md) |  | 

### Return type

[**Env**](Env.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnv

> OperationResult DeleteEnv(ctx, id).Execute()

Delete env



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
	resp, r, err := apiClient.EnvsAPI.DeleteEnv(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvsAPI.DeleteEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnv`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `EnvsAPI.DeleteEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvRequest struct via the builder pattern


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


## GetEnv

> Env GetEnv(ctx, id).Execute()

Get env



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
	resp, r, err := apiClient.EnvsAPI.GetEnv(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvsAPI.GetEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnv`: Env
	fmt.Fprintf(os.Stdout, "Response from `EnvsAPI.GetEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Env**](Env.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvByName

> Env GetEnvByName(ctx, name).OrgId(orgId).Execute()

Get env by name



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
	name := "name_example" // string | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvsAPI.GetEnvByName(context.Background(), name).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvsAPI.GetEnvByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvByName`: Env
	fmt.Fprintf(os.Stdout, "Response from `EnvsAPI.GetEnvByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**Env**](Env.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvs

> []Env ListEnvs(ctx).OrgId(orgId).Execute()

List envs



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
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvsAPI.ListEnvs(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvsAPI.ListEnvs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvs`: []Env
	fmt.Fprintf(os.Stdout, "Response from `EnvsAPI.ListEnvs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnvsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**[]Env**](Env.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnv

> Env UpdateEnv(ctx, id).UpdateEnvRequest(updateEnvRequest).Execute()

Update env



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
	updateEnvRequest := *openapiclient.NewUpdateEnvRequest("Name_example", "Title_example", "Type_example") // UpdateEnvRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvsAPI.UpdateEnv(context.Background(), id).UpdateEnvRequest(updateEnvRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvsAPI.UpdateEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnv`: Env
	fmt.Fprintf(os.Stdout, "Response from `EnvsAPI.UpdateEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnvRequest** | [**UpdateEnvRequest**](UpdateEnvRequest.md) |  | 

### Return type

[**Env**](Env.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

