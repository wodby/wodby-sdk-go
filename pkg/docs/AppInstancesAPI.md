# \AppInstancesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppAccess**](AppInstancesAPI.md#CreateAppAccess) | **Post** /app-instance-accesses/{id} | Create app instance access
[**CreateAppInstance**](AppInstancesAPI.md#CreateAppInstance) | **Post** /app-instances | Create app instance
[**DeleteAppAccess**](AppInstancesAPI.md#DeleteAppAccess) | **Delete** /app-accesses/{id} | Delete app access
[**DeleteAppInstance**](AppInstancesAPI.md#DeleteAppInstance) | **Delete** /app-instances/{id} | Delete app instance
[**GetAppInstance**](AppInstancesAPI.md#GetAppInstance) | **Get** /app-instances/{id} | Get app instance
[**GetAppInstanceAccess**](AppInstancesAPI.md#GetAppInstanceAccess) | **Get** /app-instance-accesses/{id} | Get app instance access
[**GetAppInstanceByName**](AppInstancesAPI.md#GetAppInstanceByName) | **Get** /app-instances/by-name/{appName}/{instanceName} | Get app instance by app and instance name
[**GetAppInstanceCICDSettings**](AppInstancesAPI.md#GetAppInstanceCICDSettings) | **Get** /app-instances/cicd-settings/{id} | Get app instance CI/CD settings
[**GetAppInstanceStackUpgradeChangelog**](AppInstancesAPI.md#GetAppInstanceStackUpgradeChangelog) | **Get** /app-instance-stack-upgrade-changelogs/{id} | Preview app instance stack upgrade
[**ListAppAccessCleanups**](AppInstancesAPI.md#ListAppAccessCleanups) | **Get** /app-access-cleanups | List app-access cleanups
[**ListAppInstances**](AppInstancesAPI.md#ListAppInstances) | **Get** /app-instances | List app instances
[**PreflightAppAccess**](AppInstancesAPI.md#PreflightAppAccess) | **Post** /app-accesses/actions/preflight | Preflight app instance access
[**RetryAppAccessCleanup**](AppInstancesAPI.md#RetryAppAccessCleanup) | **Post** /app-access-cleanups/{id}/actions/retry | Retry app-access cleanup
[**UpdateAppAccess**](AppInstancesAPI.md#UpdateAppAccess) | **Put** /app-accesses/{id} | Update app access
[**UpdateAppInstance**](AppInstancesAPI.md#UpdateAppInstance) | **Put** /app-instances/{id} | Update app instance
[**UpdateAppInstanceCICDSettings**](AppInstancesAPI.md#UpdateAppInstanceCICDSettings) | **Put** /app-instances/cicd-settings/{id} | Update app instance CI/CD settings
[**UpdateAppInstanceMaintenanceMode**](AppInstancesAPI.md#UpdateAppInstanceMaintenanceMode) | **Put** /app-instances/{id}/actions/maintenance-mode | Update app instance maintenance mode
[**UpdateAppInstanceSettings**](AppInstancesAPI.md#UpdateAppInstanceSettings) | **Put** /app-instances/settings/{id} | Update app instance settings
[**UpgradeAppInstanceStack**](AppInstancesAPI.md#UpgradeAppInstanceStack) | **Post** /app-instances/{id}/actions/upgrade-stack | Upgrade app instance stack



## CreateAppAccess

> AppAccessOperationResult CreateAppAccess(ctx, id).NewAppAccessInput(newAppAccessInput).Execute()

Create app instance access



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
	newAppAccessInput := *openapiclient.NewNewAppAccessInput(int32(123), "Mode_example", "Scope_example", []openapiclient.AppAccessEndpointInput{*openapiclient.NewAppAccessEndpointInput(int32(123), false)}) // NewAppAccessInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.CreateAppAccess(context.Background(), id).NewAppAccessInput(newAppAccessInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.CreateAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppAccess`: AppAccessOperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.CreateAppAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAppAccessInput** | [**NewAppAccessInput**](NewAppAccessInput.md) |  | 

### Return type

[**AppAccessOperationResult**](AppAccessOperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppInstance

> AppInstance CreateAppInstance(ctx).NewAppInstanceInput(newAppInstanceInput).Execute()

Create app instance



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
	newAppInstanceInput := *openapiclient.NewNewAppInstanceInput(int32(123), "InstanceName_example", int32(123), int32(123)) // NewAppInstanceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.CreateAppInstance(context.Background()).NewAppInstanceInput(newAppInstanceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.CreateAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppInstance`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.CreateAppInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppInstanceInput** | [**NewAppInstanceInput**](NewAppInstanceInput.md) |  | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppAccess

> OperationResult DeleteAppAccess(ctx, id).Execute()

Delete app access



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
	resp, r, err := apiClient.AppInstancesAPI.DeleteAppAccess(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.DeleteAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppAccess`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.DeleteAppAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppInstance

> OperationResult DeleteAppInstance(ctx, id).Force(force).Execute()

Delete app instance



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
	force := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.DeleteAppInstance(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.DeleteAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppInstance`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.DeleteAppInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstance

> AppInstance GetAppInstance(ctx, id).Execute()

Get app instance



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
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstance`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstanceAccess

> AppAccess GetAppInstanceAccess(ctx, id).Execute()

Get app instance access



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
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstanceAccess(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstanceAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstanceAccess`: AppAccess
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstanceAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppAccess**](AppAccess.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstanceByName

> AppInstance GetAppInstanceByName(ctx, appName, instanceName).OrgId(orgId).Execute()

Get app instance by app and instance name



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
	appName := "appName_example" // string | 
	instanceName := "instanceName_example" // string | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstanceByName(context.Background(), appName, instanceName).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstanceByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstanceByName`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstanceByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 
**instanceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstanceCICDSettings

> AppInstanceCICDSettings GetAppInstanceCICDSettings(ctx, id).Execute()

Get app instance CI/CD settings



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
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstanceCICDSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstanceCICDSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstanceCICDSettings`: AppInstanceCICDSettings
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstanceCICDSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceCICDSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppInstanceCICDSettings**](AppInstanceCICDSettings.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstanceStackUpgradeChangelog

> AppInstanceStackUpgradeChangelog GetAppInstanceStackUpgradeChangelog(ctx, id).Execute()

Preview app instance stack upgrade



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
	resp, r, err := apiClient.AppInstancesAPI.GetAppInstanceStackUpgradeChangelog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.GetAppInstanceStackUpgradeChangelog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstanceStackUpgradeChangelog`: AppInstanceStackUpgradeChangelog
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.GetAppInstanceStackUpgradeChangelog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstanceStackUpgradeChangelogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppInstanceStackUpgradeChangelog**](AppInstanceStackUpgradeChangelog.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppAccessCleanups

> []AppAccessCleanup ListAppAccessCleanups(ctx).AppInstanceId(appInstanceId).IntegrationId(integrationId).Execute()

List app-access cleanups



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
	integrationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.ListAppAccessCleanups(context.Background()).AppInstanceId(appInstanceId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.ListAppAccessCleanups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppAccessCleanups`: []AppAccessCleanup
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.ListAppAccessCleanups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppAccessCleanupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **integrationId** | **int32** |  | 

### Return type

[**[]AppAccessCleanup**](AppAccessCleanup.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppInstances

> []AppInstance ListAppInstances(ctx).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()

List app instances



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
	appId := int32(56) // int32 |  (optional)
	clusterId := int32(56) // int32 |  (optional)
	clusterApp := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.ListAppInstances(context.Background()).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.ListAppInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppInstances`: []AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.ListAppInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **appId** | **int32** |  | 
 **clusterId** | **int32** |  | 
 **clusterApp** | **bool** |  | 

### Return type

[**[]AppInstance**](AppInstance.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreflightAppAccess

> ValidationResult PreflightAppAccess(ctx).NewAppInstanceAccessInput(newAppInstanceAccessInput).Execute()

Preflight app instance access



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
	newAppInstanceAccessInput := *openapiclient.NewNewAppInstanceAccessInput(int32(123), "Mode_example", "Scope_example") // NewAppInstanceAccessInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.PreflightAppAccess(context.Background()).NewAppInstanceAccessInput(newAppInstanceAccessInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.PreflightAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreflightAppAccess`: ValidationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.PreflightAppAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreflightAppAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppInstanceAccessInput** | [**NewAppInstanceAccessInput**](NewAppInstanceAccessInput.md) |  | 

### Return type

[**ValidationResult**](ValidationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryAppAccessCleanup

> OperationResult RetryAppAccessCleanup(ctx, id).Execute()

Retry app-access cleanup



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
	resp, r, err := apiClient.AppInstancesAPI.RetryAppAccessCleanup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.RetryAppAccessCleanup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryAppAccessCleanup`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.RetryAppAccessCleanup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryAppAccessCleanupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppAccess

> AppAccessOperationResult UpdateAppAccess(ctx, id).UpdateAppAccessInput(updateAppAccessInput).Execute()

Update app access



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
	updateAppAccessInput := *openapiclient.NewUpdateAppAccessInput("Scope_example", []openapiclient.AppAccessEndpointInput{*openapiclient.NewAppAccessEndpointInput(int32(123), false)}) // UpdateAppAccessInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.UpdateAppAccess(context.Background(), id).UpdateAppAccessInput(updateAppAccessInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpdateAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppAccess`: AppAccessOperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpdateAppAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppAccessInput** | [**UpdateAppAccessInput**](UpdateAppAccessInput.md) |  | 

### Return type

[**AppAccessOperationResult**](AppAccessOperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppInstance

> AppInstance UpdateAppInstance(ctx, id).UpdateTitleRequest(updateTitleRequest).Execute()

Update app instance



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
	resp, r, err := apiClient.AppInstancesAPI.UpdateAppInstance(context.Background(), id).UpdateTitleRequest(updateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpdateAppInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppInstance`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpdateAppInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTitleRequest** | [**UpdateTitleRequest**](UpdateTitleRequest.md) |  | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppInstanceCICDSettings

> AppInstanceCICDSettings UpdateAppInstanceCICDSettings(ctx, id).AppInstanceCICDSettingsInput(appInstanceCICDSettingsInput).Execute()

Update app instance CI/CD settings



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
	appInstanceCICDSettingsInput := *openapiclient.NewAppInstanceCICDSettingsInput(int32(123), int32(123)) // AppInstanceCICDSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.UpdateAppInstanceCICDSettings(context.Background(), id).AppInstanceCICDSettingsInput(appInstanceCICDSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpdateAppInstanceCICDSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppInstanceCICDSettings`: AppInstanceCICDSettings
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpdateAppInstanceCICDSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceCICDSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appInstanceCICDSettingsInput** | [**AppInstanceCICDSettingsInput**](AppInstanceCICDSettingsInput.md) |  | 

### Return type

[**AppInstanceCICDSettings**](AppInstanceCICDSettings.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppInstanceMaintenanceMode

> OperationResult UpdateAppInstanceMaintenanceMode(ctx, id).AppInstanceMaintenanceModeInput(appInstanceMaintenanceModeInput).Execute()

Update app instance maintenance mode



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
	appInstanceMaintenanceModeInput := *openapiclient.NewAppInstanceMaintenanceModeInput(false) // AppInstanceMaintenanceModeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.UpdateAppInstanceMaintenanceMode(context.Background(), id).AppInstanceMaintenanceModeInput(appInstanceMaintenanceModeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpdateAppInstanceMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppInstanceMaintenanceMode`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpdateAppInstanceMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appInstanceMaintenanceModeInput** | [**AppInstanceMaintenanceModeInput**](AppInstanceMaintenanceModeInput.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppInstanceSettings

> AppInstance UpdateAppInstanceSettings(ctx, id).AppInstanceSettingsInput(appInstanceSettingsInput).Execute()

Update app instance settings



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
	appInstanceSettingsInput := *openapiclient.NewAppInstanceSettingsInput() // AppInstanceSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.UpdateAppInstanceSettings(context.Background(), id).AppInstanceSettingsInput(appInstanceSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpdateAppInstanceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppInstanceSettings`: AppInstance
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpdateAppInstanceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppInstanceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appInstanceSettingsInput** | [**AppInstanceSettingsInput**](AppInstanceSettingsInput.md) |  | 

### Return type

[**AppInstance**](AppInstance.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeAppInstanceStack

> OperationResult UpgradeAppInstanceStack(ctx, id).AppInstanceStackUpgradeInput(appInstanceStackUpgradeInput).Execute()

Upgrade app instance stack



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
	appInstanceStackUpgradeInput := *openapiclient.NewAppInstanceStackUpgradeInput(false, false, false, false, false, false, false, false, false, false, false, false) // AppInstanceStackUpgradeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstancesAPI.UpgradeAppInstanceStack(context.Background(), id).AppInstanceStackUpgradeInput(appInstanceStackUpgradeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstancesAPI.UpgradeAppInstanceStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeAppInstanceStack`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppInstancesAPI.UpgradeAppInstanceStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeAppInstanceStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appInstanceStackUpgradeInput** | [**AppInstanceStackUpgradeInput**](AppInstanceStackUpgradeInput.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

