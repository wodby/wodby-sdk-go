# \ImportsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImport**](ImportsAPI.md#CreateImport) | **Post** /imports | Create import
[**GetImport**](ImportsAPI.md#GetImport) | **Get** /imports/{id} | Get import
[**ListImports**](ImportsAPI.md#ListImports) | **Get** /imports | List imports



## CreateImport

> OperationResult CreateImport(ctx).CreateImportInput(createImportInput).Execute()

Create import



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
	createImportInput := *openapiclient.NewCreateImportInput(*openapiclient.NewImportInput("Source_example")) // CreateImportInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportsAPI.CreateImport(context.Background()).CreateImportInput(createImportInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.CreateImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImport`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.CreateImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImportInput** | [**CreateImportInput**](CreateImportInput.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImport

> Import GetImport(ctx, id).Execute()

Get import



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
	resp, r, err := apiClient.ImportsAPI.GetImport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.GetImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImport`: Import
	fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.GetImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Import**](Import.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImports

> []Import ListImports(ctx).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).Execute()

List imports



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
	appInstanceId := int32(56) // int32 |  (optional)
	appServiceId := int32(56) // int32 |  (optional)
	databaseId := int32(56) // int32 |  (optional)
	databaseDbId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportsAPI.ListImports(context.Background()).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImports`: []Import
	fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **appServiceId** | **int32** |  | 
 **databaseId** | **int32** |  | 
 **databaseDbId** | **int32** |  | 

### Return type

[**[]Import**](Import.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

