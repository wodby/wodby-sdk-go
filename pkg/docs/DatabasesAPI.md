# \DatabasesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabasesAPI.md#CreateDatabase) | **Post** /databases | Create database
[**DeleteDatabase**](DatabasesAPI.md#DeleteDatabase) | **Delete** /databases/{id} | Delete database
[**GetDatabase**](DatabasesAPI.md#GetDatabase) | **Get** /databases/{id} | Get database
[**GetDatabaseByName**](DatabasesAPI.md#GetDatabaseByName) | **Get** /databases/by-name/{name} | Get database by name
[**ListDatabases**](DatabasesAPI.md#ListDatabases) | **Get** /databases | List databases
[**UpdateDatabase**](DatabasesAPI.md#UpdateDatabase) | **Put** /databases/{id} | Update database



## CreateDatabase

> Database CreateDatabase(ctx).NewDatabaseInput(newDatabaseInput).Execute()

Create database

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
	newDatabaseInput := *openapiclient.NewNewDatabaseInput(int32(123), int32(123), "Name_example", "Title_example", int32(123), "Type_example", "Version_example", "MachineType_example") // NewDatabaseInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.CreateDatabase(context.Background()).NewDatabaseInput(newDatabaseInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDatabaseInput** | [**NewDatabaseInput**](NewDatabaseInput.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> OperationResult DeleteDatabase(ctx, id).Execute()

Delete database

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
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabase(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabase`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDatabase

> Database GetDatabase(ctx, id).Execute()

Get database

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
	resp, r, err := apiClient.DatabasesAPI.GetDatabase(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabase`: Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Database**](Database.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseByName

> Database GetDatabaseByName(ctx, name).OrgId(orgId).Execute()

Get database by name

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
	resp, r, err := apiClient.DatabasesAPI.GetDatabaseByName(context.Background(), name).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseByName`: Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** |  | 

### Return type

[**Database**](Database.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> []Database ListDatabases(ctx).OrgId(orgId).ProjectIds(projectIds).Kind(kind).Execute()

List databases

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
	kind := "kind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.ListDatabases(context.Background()).OrgId(orgId).ProjectIds(projectIds).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabases`: []Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** |  | 
 **projectIds** | **string** | Comma-separated project ids | 
 **kind** | **string** |  | 

### Return type

[**[]Database**](Database.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabase

> Database UpdateDatabase(ctx, id).UpdateTitleRequest(updateTitleRequest).Execute()

Update database

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
	resp, r, err := apiClient.DatabasesAPI.UpdateDatabase(context.Background(), id).UpdateTitleRequest(updateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDatabase`: Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTitleRequest** | [**UpdateTitleRequest**](UpdateTitleRequest.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

