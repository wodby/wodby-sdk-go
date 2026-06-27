# \DatabasesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabasesAPI.md#CreateDatabase) | **Post** /databases | Create database
[**CreateDatabaseDB**](DatabasesAPI.md#CreateDatabaseDB) | **Post** /database-dbs | Create database DB
[**CreateDatabaseUser**](DatabasesAPI.md#CreateDatabaseUser) | **Post** /database-users | Create database user
[**DeleteDatabase**](DatabasesAPI.md#DeleteDatabase) | **Delete** /databases/{id} | Delete database
[**DeleteDatabaseDB**](DatabasesAPI.md#DeleteDatabaseDB) | **Delete** /database-dbs/{id} | Delete database DB
[**DeleteDatabaseUser**](DatabasesAPI.md#DeleteDatabaseUser) | **Delete** /database-users/{id} | Delete database user
[**GetDatabase**](DatabasesAPI.md#GetDatabase) | **Get** /databases/{id} | Get database
[**GetDatabaseByName**](DatabasesAPI.md#GetDatabaseByName) | **Get** /databases/by-name/{name} | Get database by name
[**GetDatabaseDB**](DatabasesAPI.md#GetDatabaseDB) | **Get** /database-dbs/{id} | Get database DB
[**ListDatabaseCharsets**](DatabasesAPI.md#ListDatabaseCharsets) | **Get** /databases/{id}/options/charsets | List database charsets
[**ListDatabaseDBs**](DatabasesAPI.md#ListDatabaseDBs) | **Get** /database-dbs | List database DBs
[**ListDatabaseUsers**](DatabasesAPI.md#ListDatabaseUsers) | **Get** /database-users | List database users
[**ListDatabases**](DatabasesAPI.md#ListDatabases) | **Get** /databases | List databases
[**UpdateDatabase**](DatabasesAPI.md#UpdateDatabase) | **Put** /databases/{id} | Update database
[**UpdateDatabaseUserDBs**](DatabasesAPI.md#UpdateDatabaseUserDBs) | **Put** /database-users/{id}/dbs | Update database user DB grants



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
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	newDatabaseInput := *openapiclient.NewNewDatabaseInput(int32(123), "Name_example", "Title_example", int32(123), "Type_example", "Version_example", "MachineType_example") // NewDatabaseInput | 

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

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseDB

> DatabaseDB CreateDatabaseDB(ctx).NewDatabaseDBInput(newDatabaseDBInput).Execute()

Create database DB



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
	newDatabaseDBInput := *openapiclient.NewNewDatabaseDBInput(int32(123), "Name_example") // NewDatabaseDBInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.CreateDatabaseDB(context.Background()).NewDatabaseDBInput(newDatabaseDBInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabaseDB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabaseDB`: DatabaseDB
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabaseDB`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseDBRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDatabaseDBInput** | [**NewDatabaseDBInput**](NewDatabaseDBInput.md) |  | 

### Return type

[**DatabaseDB**](DatabaseDB.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseUser

> DatabaseUser CreateDatabaseUser(ctx).NewDatabaseUserInput(newDatabaseUserInput).Execute()

Create database user



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
	newDatabaseUserInput := *openapiclient.NewNewDatabaseUserInput(int32(123), "Name_example", "Password_example") // NewDatabaseUserInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.CreateDatabaseUser(context.Background()).NewDatabaseUserInput(newDatabaseUserInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabaseUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabaseUser`: DatabaseUser
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabaseUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDatabaseUserInput** | [**NewDatabaseUserInput**](NewDatabaseUserInput.md) |  | 

### Return type

[**DatabaseUser**](DatabaseUser.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

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
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
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

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseDB

> OperationResult DeleteDatabaseDB(ctx, id).Execute()

Delete database DB



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
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabaseDB(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseDB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabaseDB`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabaseDB`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseDBRequest struct via the builder pattern


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


## DeleteDatabaseUser

> OperationResult DeleteDatabaseUser(ctx, id).Execute()

Delete database user



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
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabaseUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabaseUser`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabaseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseUserRequest struct via the builder pattern


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
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
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

[apiKeyHeader](../README.md#apiKeyHeader)

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
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	name := "name_example" // string | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

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

 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**Database**](Database.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseDB

> DatabaseDB GetDatabaseDB(ctx, id).Execute()

Get database DB



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
	resp, r, err := apiClient.DatabasesAPI.GetDatabaseDB(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseDB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseDB`: DatabaseDB
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseDB`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseDBRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseDB**](DatabaseDB.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseCharsets

> []DatabaseCharset ListDatabaseCharsets(ctx, id).Execute()

List database charsets



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
	resp, r, err := apiClient.DatabasesAPI.ListDatabaseCharsets(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabaseCharsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseCharsets`: []DatabaseCharset
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabaseCharsets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseCharsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DatabaseCharset**](DatabaseCharset.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseDBs

> []DatabaseDB ListDatabaseDBs(ctx).DatabaseId(databaseId).Execute()

List database DBs



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
	databaseId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.ListDatabaseDBs(context.Background()).DatabaseId(databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabaseDBs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseDBs`: []DatabaseDB
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabaseDBs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseDBsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseId** | **int32** |  | 

### Return type

[**[]DatabaseDB**](DatabaseDB.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseUsers

> []DatabaseUser ListDatabaseUsers(ctx).DatabaseId(databaseId).Execute()

List database users



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
	databaseId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.ListDatabaseUsers(context.Background()).DatabaseId(databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseUsers`: []DatabaseUser
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseId** | **int32** |  | 

### Return type

[**[]DatabaseUser**](DatabaseUser.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

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
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)
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
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **kind** | **string** |  | 

### Return type

[**[]Database**](Database.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

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
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
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

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseUserDBs

> DatabaseUser UpdateDatabaseUserDBs(ctx, id).UpdateDatabaseUserDBsInput(updateDatabaseUserDBsInput).Execute()

Update database user DB grants



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
	updateDatabaseUserDBsInput := *openapiclient.NewUpdateDatabaseUserDBsInput([]int32{int32(123)}) // UpdateDatabaseUserDBsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.UpdateDatabaseUserDBs(context.Background(), id).UpdateDatabaseUserDBsInput(updateDatabaseUserDBsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabaseUserDBs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDatabaseUserDBs`: DatabaseUser
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabaseUserDBs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseUserDBsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabaseUserDBsInput** | [**UpdateDatabaseUserDBsInput**](UpdateDatabaseUserDBsInput.md) |  | 

### Return type

[**DatabaseUser**](DatabaseUser.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

