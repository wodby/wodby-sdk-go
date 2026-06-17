# \IntegrationKindsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationKindsIdDatabaseMachineTypesGet**](IntegrationKindsAPI.md#IntegrationKindsIdDatabaseMachineTypesGet) | **Get** /integration-kinds/{id}/database-machine-types | List database machine types
[**IntegrationKindsIdDatabaseRegionsGet**](IntegrationKindsAPI.md#IntegrationKindsIdDatabaseRegionsGet) | **Get** /integration-kinds/{id}/database-regions | List database regions
[**IntegrationKindsIdDatabaseSettingsGet**](IntegrationKindsAPI.md#IntegrationKindsIdDatabaseSettingsGet) | **Get** /integration-kinds/{id}/database-settings | Get database settings
[**IntegrationKindsIdDatabaseTypesGet**](IntegrationKindsAPI.md#IntegrationKindsIdDatabaseTypesGet) | **Get** /integration-kinds/{id}/database-types | List database types
[**IntegrationKindsIdDatabaseVersionsGet**](IntegrationKindsAPI.md#IntegrationKindsIdDatabaseVersionsGet) | **Get** /integration-kinds/{id}/database-versions | List database versions



## IntegrationKindsIdDatabaseMachineTypesGet

> []map[string]interface{} IntegrationKindsIdDatabaseMachineTypesGet(ctx, id).DbType(dbType).Version(version).Ha(ha).Region(region).Zone(zone).Execute()

List database machine types

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
	dbType := "dbType_example" // string | 
	version := "version_example" // string | 
	ha := true // bool |  (optional) (default to false)
	region := "region_example" // string |  (optional)
	zone := "zone_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.IntegrationKindsIdDatabaseMachineTypesGet(context.Background(), id).DbType(dbType).Version(version).Ha(ha).Region(region).Zone(zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.IntegrationKindsIdDatabaseMachineTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationKindsIdDatabaseMachineTypesGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.IntegrationKindsIdDatabaseMachineTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationKindsIdDatabaseMachineTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 
 **version** | **string** |  | 
 **ha** | **bool** |  | [default to false]
 **region** | **string** |  | 
 **zone** | **string** |  | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationKindsIdDatabaseRegionsGet

> []map[string]interface{} IntegrationKindsIdDatabaseRegionsGet(ctx, id).DbType(dbType).Version(version).Ha(ha).Execute()

List database regions

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
	dbType := "dbType_example" // string | 
	version := "version_example" // string | 
	ha := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.IntegrationKindsIdDatabaseRegionsGet(context.Background(), id).DbType(dbType).Version(version).Ha(ha).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.IntegrationKindsIdDatabaseRegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationKindsIdDatabaseRegionsGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.IntegrationKindsIdDatabaseRegionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationKindsIdDatabaseRegionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 
 **version** | **string** |  | 
 **ha** | **bool** |  | [default to false]

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationKindsIdDatabaseSettingsGet

> map[string]interface{} IntegrationKindsIdDatabaseSettingsGet(ctx, id).DbType(dbType).Execute()

Get database settings

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
	dbType := "dbType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.IntegrationKindsIdDatabaseSettingsGet(context.Background(), id).DbType(dbType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.IntegrationKindsIdDatabaseSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationKindsIdDatabaseSettingsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.IntegrationKindsIdDatabaseSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationKindsIdDatabaseSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationKindsIdDatabaseTypesGet

> []DatabaseType IntegrationKindsIdDatabaseTypesGet(ctx, id).Execute()

List database types

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
	resp, r, err := apiClient.IntegrationKindsAPI.IntegrationKindsIdDatabaseTypesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.IntegrationKindsIdDatabaseTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationKindsIdDatabaseTypesGet`: []DatabaseType
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.IntegrationKindsIdDatabaseTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationKindsIdDatabaseTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DatabaseType**](DatabaseType.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationKindsIdDatabaseVersionsGet

> []DatabaseVersion IntegrationKindsIdDatabaseVersionsGet(ctx, id).DbType(dbType).Execute()

List database versions

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
	dbType := "dbType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.IntegrationKindsIdDatabaseVersionsGet(context.Background(), id).DbType(dbType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.IntegrationKindsIdDatabaseVersionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationKindsIdDatabaseVersionsGet`: []DatabaseVersion
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.IntegrationKindsIdDatabaseVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationKindsIdDatabaseVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 

### Return type

[**[]DatabaseVersion**](DatabaseVersion.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

