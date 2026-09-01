# \AppEnvironmentsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppAccess**](AppEnvironmentsAPI.md#CreateAppAccess) | **Post** /app-environment-accesses/{id} | Create app environment access
[**CreateAppEnvironment**](AppEnvironmentsAPI.md#CreateAppEnvironment) | **Post** /app-environments | Create app environment
[**DeleteAppAccess**](AppEnvironmentsAPI.md#DeleteAppAccess) | **Delete** /app-accesses/{id} | Delete app access
[**DeleteAppEnvironment**](AppEnvironmentsAPI.md#DeleteAppEnvironment) | **Delete** /app-environments/{id} | Delete app environment
[**GetAppEnvironment**](AppEnvironmentsAPI.md#GetAppEnvironment) | **Get** /app-environments/{id} | Get app environment
[**GetAppEnvironmentAccess**](AppEnvironmentsAPI.md#GetAppEnvironmentAccess) | **Get** /app-environment-accesses/{id} | Get app environment access
[**GetAppEnvironmentByName**](AppEnvironmentsAPI.md#GetAppEnvironmentByName) | **Get** /app-environments/by-name/{appName}/{environmentName} | Get app environment by name
[**GetAppEnvironmentCICDSettings**](AppEnvironmentsAPI.md#GetAppEnvironmentCICDSettings) | **Get** /app-environments/cicd-settings/{id} | Get app environment CI/CD settings
[**GetAppEnvironmentStackUpgradeChangelog**](AppEnvironmentsAPI.md#GetAppEnvironmentStackUpgradeChangelog) | **Get** /app-environment-stack-upgrade-changelogs/{id} | Preview app environment stack upgrade
[**ListAppAccessCleanups**](AppEnvironmentsAPI.md#ListAppAccessCleanups) | **Get** /app-access-cleanups | List app-access cleanups
[**ListAppEnvironments**](AppEnvironmentsAPI.md#ListAppEnvironments) | **Get** /app-environments | List app environments
[**PreflightAppAccess**](AppEnvironmentsAPI.md#PreflightAppAccess) | **Post** /app-accesses/actions/preflight | Preflight app environment access
[**ReconcileAppEnvironmentStack**](AppEnvironmentsAPI.md#ReconcileAppEnvironmentStack) | **Post** /app-environments/{id}/actions/reconcile-stack | Reconcile app environment stack
[**RetryAppAccessCleanup**](AppEnvironmentsAPI.md#RetryAppAccessCleanup) | **Post** /app-access-cleanups/{id}/actions/retry | Retry app-access cleanup
[**UpdateAppAccess**](AppEnvironmentsAPI.md#UpdateAppAccess) | **Put** /app-accesses/{id} | Update app access
[**UpdateAppEnvironment**](AppEnvironmentsAPI.md#UpdateAppEnvironment) | **Put** /app-environments/{id} | Update app environment
[**UpdateAppEnvironmentCICDSettings**](AppEnvironmentsAPI.md#UpdateAppEnvironmentCICDSettings) | **Put** /app-environments/cicd-settings/{id} | Update app environment CI/CD settings
[**UpdateAppEnvironmentMaintenanceMode**](AppEnvironmentsAPI.md#UpdateAppEnvironmentMaintenanceMode) | **Put** /app-environments/{id}/actions/maintenance-mode | Update app environment maintenance mode
[**UpdateAppEnvironmentSettings**](AppEnvironmentsAPI.md#UpdateAppEnvironmentSettings) | **Put** /app-environments/settings/{id} | Update app environment settings
[**UpgradeAppEnvironmentStack**](AppEnvironmentsAPI.md#UpgradeAppEnvironmentStack) | **Post** /app-environments/{id}/actions/upgrade-stack | Upgrade app environment stack



## CreateAppAccess

> AppAccessOperationResult CreateAppAccess(ctx, id).NewAppAccessInput(newAppAccessInput).Execute()

Create app environment access



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
	resp, r, err := apiClient.AppEnvironmentsAPI.CreateAppAccess(context.Background(), id).NewAppAccessInput(newAppAccessInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.CreateAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppAccess`: AppAccessOperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.CreateAppAccess`: %v\n", resp)
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


## CreateAppEnvironment

> AppEnvironment CreateAppEnvironment(ctx).NewAppEnvironmentInput(newAppEnvironmentInput).Execute()

Create app environment



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
	newAppEnvironmentInput := *openapiclient.NewNewAppEnvironmentInput(int32(123), "EnvironmentName_example", "EnvironmentType_example", int32(123)) // NewAppEnvironmentInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.CreateAppEnvironment(context.Background()).NewAppEnvironmentInput(newAppEnvironmentInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.CreateAppEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppEnvironment`: AppEnvironment
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.CreateAppEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppEnvironmentInput** | [**NewAppEnvironmentInput**](NewAppEnvironmentInput.md) |  | 

### Return type

[**AppEnvironment**](AppEnvironment.md)

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
	resp, r, err := apiClient.AppEnvironmentsAPI.DeleteAppAccess(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.DeleteAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppAccess`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.DeleteAppAccess`: %v\n", resp)
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


## DeleteAppEnvironment

> OperationResult DeleteAppEnvironment(ctx, id).Force(force).Execute()

Delete app environment



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
	resp, r, err := apiClient.AppEnvironmentsAPI.DeleteAppEnvironment(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.DeleteAppEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppEnvironment`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.DeleteAppEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppEnvironmentRequest struct via the builder pattern


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


## GetAppEnvironment

> AppEnvironment GetAppEnvironment(ctx, id).Execute()

Get app environment



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
	resp, r, err := apiClient.AppEnvironmentsAPI.GetAppEnvironment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.GetAppEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironment`: AppEnvironment
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.GetAppEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEnvironment**](AppEnvironment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppEnvironmentAccess

> AppAccess GetAppEnvironmentAccess(ctx, id).Execute()

Get app environment access



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
	resp, r, err := apiClient.AppEnvironmentsAPI.GetAppEnvironmentAccess(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.GetAppEnvironmentAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironmentAccess`: AppAccess
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.GetAppEnvironmentAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentAccessRequest struct via the builder pattern


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


## GetAppEnvironmentByName

> AppEnvironment GetAppEnvironmentByName(ctx, appName, environmentName).OrgId(orgId).Execute()

Get app environment by name



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
	environmentName := "environmentName_example" // string | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.GetAppEnvironmentByName(context.Background(), appName, environmentName).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.GetAppEnvironmentByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironmentByName`: AppEnvironment
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.GetAppEnvironmentByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 
**environmentName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AppEnvironment**](AppEnvironment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppEnvironmentCICDSettings

> AppEnvironmentCICDSettings GetAppEnvironmentCICDSettings(ctx, id).Execute()

Get app environment CI/CD settings



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
	resp, r, err := apiClient.AppEnvironmentsAPI.GetAppEnvironmentCICDSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.GetAppEnvironmentCICDSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironmentCICDSettings`: AppEnvironmentCICDSettings
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.GetAppEnvironmentCICDSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentCICDSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEnvironmentCICDSettings**](AppEnvironmentCICDSettings.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppEnvironmentStackUpgradeChangelog

> AppEnvironmentStackUpgradeChangelog GetAppEnvironmentStackUpgradeChangelog(ctx, id).Execute()

Preview app environment stack upgrade



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
	resp, r, err := apiClient.AppEnvironmentsAPI.GetAppEnvironmentStackUpgradeChangelog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.GetAppEnvironmentStackUpgradeChangelog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironmentStackUpgradeChangelog`: AppEnvironmentStackUpgradeChangelog
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.GetAppEnvironmentStackUpgradeChangelog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentStackUpgradeChangelogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEnvironmentStackUpgradeChangelog**](AppEnvironmentStackUpgradeChangelog.md)

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
	resp, r, err := apiClient.AppEnvironmentsAPI.ListAppAccessCleanups(context.Background()).AppInstanceId(appInstanceId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.ListAppAccessCleanups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppAccessCleanups`: []AppAccessCleanup
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.ListAppAccessCleanups`: %v\n", resp)
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


## ListAppEnvironments

> []AppEnvironment ListAppEnvironments(ctx).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()

List app environments



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
	resp, r, err := apiClient.AppEnvironmentsAPI.ListAppEnvironments(context.Background()).OrgId(orgId).ProjectIds(projectIds).AppId(appId).ClusterId(clusterId).ClusterApp(clusterApp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.ListAppEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppEnvironments`: []AppEnvironment
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.ListAppEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **appId** | **int32** |  | 
 **clusterId** | **int32** |  | 
 **clusterApp** | **bool** |  | 

### Return type

[**[]AppEnvironment**](AppEnvironment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreflightAppAccess

> ValidationResult PreflightAppAccess(ctx).NewAppEnvironmentAccessInput(newAppEnvironmentAccessInput).Execute()

Preflight app environment access



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
	newAppEnvironmentAccessInput := *openapiclient.NewNewAppEnvironmentAccessInput(int32(123), "Mode_example", "Scope_example") // NewAppEnvironmentAccessInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.PreflightAppAccess(context.Background()).NewAppEnvironmentAccessInput(newAppEnvironmentAccessInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.PreflightAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreflightAppAccess`: ValidationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.PreflightAppAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreflightAppAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppEnvironmentAccessInput** | [**NewAppEnvironmentAccessInput**](NewAppEnvironmentAccessInput.md) |  | 

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


## ReconcileAppEnvironmentStack

> OperationResult ReconcileAppEnvironmentStack(ctx, id).AppEnvironmentStackReconciliationInput(appEnvironmentStackReconciliationInput).Execute()

Reconcile app environment stack



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
	appEnvironmentStackReconciliationInput := *openapiclient.NewAppEnvironmentStackReconciliationInput(false, false, false, false, false, false, false, false, false, false, false, false) // AppEnvironmentStackReconciliationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.ReconcileAppEnvironmentStack(context.Background(), id).AppEnvironmentStackReconciliationInput(appEnvironmentStackReconciliationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.ReconcileAppEnvironmentStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconcileAppEnvironmentStack`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.ReconcileAppEnvironmentStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconcileAppEnvironmentStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEnvironmentStackReconciliationInput** | [**AppEnvironmentStackReconciliationInput**](AppEnvironmentStackReconciliationInput.md) |  | 

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
	resp, r, err := apiClient.AppEnvironmentsAPI.RetryAppAccessCleanup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.RetryAppAccessCleanup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryAppAccessCleanup`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.RetryAppAccessCleanup`: %v\n", resp)
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
	resp, r, err := apiClient.AppEnvironmentsAPI.UpdateAppAccess(context.Background(), id).UpdateAppAccessInput(updateAppAccessInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.UpdateAppAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppAccess`: AppAccessOperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.UpdateAppAccess`: %v\n", resp)
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


## UpdateAppEnvironment

> AppEnvironment UpdateAppEnvironment(ctx, id).UpdateTitleRequest(updateTitleRequest).Execute()

Update app environment



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
	resp, r, err := apiClient.AppEnvironmentsAPI.UpdateAppEnvironment(context.Background(), id).UpdateTitleRequest(updateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.UpdateAppEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEnvironment`: AppEnvironment
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.UpdateAppEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTitleRequest** | [**UpdateTitleRequest**](UpdateTitleRequest.md) |  | 

### Return type

[**AppEnvironment**](AppEnvironment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppEnvironmentCICDSettings

> AppEnvironmentCICDSettings UpdateAppEnvironmentCICDSettings(ctx, id).AppEnvironmentCICDSettingsInput(appEnvironmentCICDSettingsInput).Execute()

Update app environment CI/CD settings



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
	appEnvironmentCICDSettingsInput := *openapiclient.NewAppEnvironmentCICDSettingsInput(int32(123), int32(123)) // AppEnvironmentCICDSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.UpdateAppEnvironmentCICDSettings(context.Background(), id).AppEnvironmentCICDSettingsInput(appEnvironmentCICDSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.UpdateAppEnvironmentCICDSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEnvironmentCICDSettings`: AppEnvironmentCICDSettings
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.UpdateAppEnvironmentCICDSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEnvironmentCICDSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEnvironmentCICDSettingsInput** | [**AppEnvironmentCICDSettingsInput**](AppEnvironmentCICDSettingsInput.md) |  | 

### Return type

[**AppEnvironmentCICDSettings**](AppEnvironmentCICDSettings.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppEnvironmentMaintenanceMode

> OperationResult UpdateAppEnvironmentMaintenanceMode(ctx, id).AppEnvironmentMaintenanceModeInput(appEnvironmentMaintenanceModeInput).Execute()

Update app environment maintenance mode



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
	appEnvironmentMaintenanceModeInput := *openapiclient.NewAppEnvironmentMaintenanceModeInput(false) // AppEnvironmentMaintenanceModeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.UpdateAppEnvironmentMaintenanceMode(context.Background(), id).AppEnvironmentMaintenanceModeInput(appEnvironmentMaintenanceModeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.UpdateAppEnvironmentMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEnvironmentMaintenanceMode`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.UpdateAppEnvironmentMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEnvironmentMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEnvironmentMaintenanceModeInput** | [**AppEnvironmentMaintenanceModeInput**](AppEnvironmentMaintenanceModeInput.md) |  | 

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


## UpdateAppEnvironmentSettings

> AppEnvironment UpdateAppEnvironmentSettings(ctx, id).AppEnvironmentSettingsInput(appEnvironmentSettingsInput).Execute()

Update app environment settings



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
	appEnvironmentSettingsInput := *openapiclient.NewAppEnvironmentSettingsInput() // AppEnvironmentSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.UpdateAppEnvironmentSettings(context.Background(), id).AppEnvironmentSettingsInput(appEnvironmentSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.UpdateAppEnvironmentSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEnvironmentSettings`: AppEnvironment
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.UpdateAppEnvironmentSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEnvironmentSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEnvironmentSettingsInput** | [**AppEnvironmentSettingsInput**](AppEnvironmentSettingsInput.md) |  | 

### Return type

[**AppEnvironment**](AppEnvironment.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeAppEnvironmentStack

> OperationResult UpgradeAppEnvironmentStack(ctx, id).AppEnvironmentStackUpgradeInput(appEnvironmentStackUpgradeInput).Execute()

Upgrade app environment stack



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
	appEnvironmentStackUpgradeInput := *openapiclient.NewAppEnvironmentStackUpgradeInput(false, false, false, false, false, false, false, false, false, false, false, false) // AppEnvironmentStackUpgradeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEnvironmentsAPI.UpgradeAppEnvironmentStack(context.Background(), id).AppEnvironmentStackUpgradeInput(appEnvironmentStackUpgradeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEnvironmentsAPI.UpgradeAppEnvironmentStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeAppEnvironmentStack`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppEnvironmentsAPI.UpgradeAppEnvironmentStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeAppEnvironmentStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEnvironmentStackUpgradeInput** | [**AppEnvironmentStackUpgradeInput**](AppEnvironmentStackUpgradeInput.md) |  | 

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

