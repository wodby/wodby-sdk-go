# \StackServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStackService**](StackServicesAPI.md#CreateStackService) | **Post** /stack-services | Create stack service
[**DeleteStackService**](StackServicesAPI.md#DeleteStackService) | **Delete** /stack-services/{id} | Delete stack service
[**ListStackServices**](StackServicesAPI.md#ListStackServices) | **Get** /stack-services | List stack services
[**UpdateStackService**](StackServicesAPI.md#UpdateStackService) | **Put** /stack-services/{id} | Update stack service



## CreateStackService

> StackService CreateStackService(ctx).NewStackServiceInput(newStackServiceInput).Execute()

Create stack service



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
	newStackServiceInput := *openapiclient.NewNewStackServiceInput(int32(123), int32(123), "Name_example", "Title_example", false, int32(123)) // NewStackServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackService(context.Background()).NewStackServiceInput(newStackServiceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackService`: StackService
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newStackServiceInput** | [**NewStackServiceInput**](NewStackServiceInput.md) |  | 

### Return type

[**StackService**](StackService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStackService

> OperationResult DeleteStackService(ctx, id).Execute()

Delete stack service



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackService`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceRequest struct via the builder pattern


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


## ListStackServices

> []StackService ListStackServices(ctx).StackRevId(stackRevId).Execute()

List stack services



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
	stackRevId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.ListStackServices(context.Background()).StackRevId(stackRevId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServices`: []StackService
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStackServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackRevId** | **int32** |  | 

### Return type

[**[]StackService**](StackService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackService

> StackService UpdateStackService(ctx, id).StackServiceInput(stackServiceInput).Execute()

Update stack service



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
	stackServiceInput := *openapiclient.NewStackServiceInput() // StackServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.UpdateStackService(context.Background(), id).StackServiceInput(stackServiceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.UpdateStackService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackService`: StackService
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.UpdateStackService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackServiceInput** | [**StackServiceInput**](StackServiceInput.md) |  | 

### Return type

[**StackService**](StackService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

