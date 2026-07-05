# \IntegrationKindsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIntegrationKindDatabaseSettings**](IntegrationKindsAPI.md#GetIntegrationKindDatabaseSettings) | **Get** /integration-kinds/{id}/database-settings | Get database settings
[**ListIntegrationKindDatabaseMachineTypes**](IntegrationKindsAPI.md#ListIntegrationKindDatabaseMachineTypes) | **Get** /integration-kinds/{id}/database-machine-types | List database machine types
[**ListIntegrationKindDatabaseRegions**](IntegrationKindsAPI.md#ListIntegrationKindDatabaseRegions) | **Get** /integration-kinds/{id}/database-regions | List database regions
[**ListIntegrationKindDatabaseTypes**](IntegrationKindsAPI.md#ListIntegrationKindDatabaseTypes) | **Get** /integration-kinds/{id}/database-types | List database types
[**ListIntegrationKindDatabaseVersions**](IntegrationKindsAPI.md#ListIntegrationKindDatabaseVersions) | **Get** /integration-kinds/{id}/database-versions | List database versions



## GetIntegrationKindDatabaseSettings

> map[string]interface{} GetIntegrationKindDatabaseSettings(ctx, id).DbType(dbType).Execute()

Get database settings



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
	dbType := "dbType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.GetIntegrationKindDatabaseSettings(context.Background(), id).DbType(dbType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.GetIntegrationKindDatabaseSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationKindDatabaseSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.GetIntegrationKindDatabaseSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationKindDatabaseSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKindDatabaseMachineTypes

> []map[string]interface{} ListIntegrationKindDatabaseMachineTypes(ctx, id).DbType(dbType).Version(version).Ha(ha).Region(region).Zone(zone).Execute()

List database machine types



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
	dbType := "dbType_example" // string | 
	version := "version_example" // string | 
	ha := true // bool |  (optional) (default to false)
	region := "region_example" // string |  (optional)
	zone := "zone_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.ListIntegrationKindDatabaseMachineTypes(context.Background(), id).DbType(dbType).Version(version).Ha(ha).Region(region).Zone(zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.ListIntegrationKindDatabaseMachineTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKindDatabaseMachineTypes`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.ListIntegrationKindDatabaseMachineTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKindDatabaseMachineTypesRequest struct via the builder pattern


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

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKindDatabaseRegions

> []map[string]interface{} ListIntegrationKindDatabaseRegions(ctx, id).DbType(dbType).Version(version).Ha(ha).Execute()

List database regions



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
	dbType := "dbType_example" // string | 
	version := "version_example" // string | 
	ha := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.ListIntegrationKindDatabaseRegions(context.Background(), id).DbType(dbType).Version(version).Ha(ha).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.ListIntegrationKindDatabaseRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKindDatabaseRegions`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.ListIntegrationKindDatabaseRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKindDatabaseRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 
 **version** | **string** |  | 
 **ha** | **bool** |  | [default to false]

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKindDatabaseTypes

> []DatabaseType ListIntegrationKindDatabaseTypes(ctx, id).Execute()

List database types



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
	resp, r, err := apiClient.IntegrationKindsAPI.ListIntegrationKindDatabaseTypes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.ListIntegrationKindDatabaseTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKindDatabaseTypes`: []DatabaseType
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.ListIntegrationKindDatabaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKindDatabaseTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DatabaseType**](DatabaseType.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKindDatabaseVersions

> []DatabaseVersion ListIntegrationKindDatabaseVersions(ctx, id).DbType(dbType).Execute()

List database versions



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
	dbType := "dbType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKindsAPI.ListIntegrationKindDatabaseVersions(context.Background(), id).DbType(dbType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKindsAPI.ListIntegrationKindDatabaseVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKindDatabaseVersions`: []DatabaseVersion
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKindsAPI.ListIntegrationKindDatabaseVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKindDatabaseVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbType** | **string** |  | 

### Return type

[**[]DatabaseVersion**](DatabaseVersion.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

