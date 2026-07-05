# \AppServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppServiceAnnotation**](AppServicesAPI.md#CreateAppServiceAnnotation) | **Post** /app-services/{id}/annotations | Create app service annotation
[**CreateAppServiceCronSchedule**](AppServicesAPI.md#CreateAppServiceCronSchedule) | **Post** /app-services/{id}/cron-schedules | Create app service cron schedule
[**CreateAppServiceEnvVar**](AppServicesAPI.md#CreateAppServiceEnvVar) | **Post** /app-services/{id}/env-vars | Create app service env var
[**CreateAppServiceHelmValue**](AppServicesAPI.md#CreateAppServiceHelmValue) | **Post** /app-services/{id}/helm-values | Create app service Helm value
[**CreateAppServiceIntegration**](AppServicesAPI.md#CreateAppServiceIntegration) | **Post** /app-services/{id}/integrations | Create app service integration
[**CreateAppServiceLogStream**](AppServicesAPI.md#CreateAppServiceLogStream) | **Post** /app-services/{id}/log-streams | Create app service log stream
[**CreateAppServiceToken**](AppServicesAPI.md#CreateAppServiceToken) | **Post** /app-services/{id}/tokens | Create app service token
[**DeleteAppServiceAnnotation**](AppServicesAPI.md#DeleteAppServiceAnnotation) | **Delete** /app-service-annotations/{id} | Delete app service annotation
[**DeleteAppServiceCronSchedule**](AppServicesAPI.md#DeleteAppServiceCronSchedule) | **Delete** /app-service-cron-schedules/{id} | Delete app service cron schedule
[**DeleteAppServiceEnvVar**](AppServicesAPI.md#DeleteAppServiceEnvVar) | **Delete** /app-service-env-vars/{id} | Delete app service env var
[**DeleteAppServiceHelmValue**](AppServicesAPI.md#DeleteAppServiceHelmValue) | **Delete** /app-service-helm-values/{id} | Delete app service Helm value
[**DeleteAppServiceIntegration**](AppServicesAPI.md#DeleteAppServiceIntegration) | **Delete** /app-service-integrations/{id} | Delete app service integration
[**DeleteAppServiceToken**](AppServicesAPI.md#DeleteAppServiceToken) | **Delete** /app-service-tokens/{id} | Delete app service token
[**GetAppService**](AppServicesAPI.md#GetAppService) | **Get** /app-services/{id} | Get app service
[**GetAppServiceCronJob**](AppServicesAPI.md#GetAppServiceCronJob) | **Get** /app-service-cron-jobs/{id} | Get app service cron job
[**KeepLogStreamAlive**](AppServicesAPI.md#KeepLogStreamAlive) | **Post** /log-streams/{id}/keep-alive | Keep log stream alive
[**ListAppServiceAnnotations**](AppServicesAPI.md#ListAppServiceAnnotations) | **Get** /app-services/{id}/annotations | List app service annotations
[**ListAppServiceConfigs**](AppServicesAPI.md#ListAppServiceConfigs) | **Get** /app-services/{id}/configs | List app service configs
[**ListAppServiceContainers**](AppServicesAPI.md#ListAppServiceContainers) | **Get** /app-services/{id}/containers | List app service containers
[**ListAppServiceCronJobs**](AppServicesAPI.md#ListAppServiceCronJobs) | **Get** /app-service-cron-jobs | List app service cron jobs
[**ListAppServiceCronSchedules**](AppServicesAPI.md#ListAppServiceCronSchedules) | **Get** /app-services/{id}/cron-schedules | List app service cron schedules
[**ListAppServiceEnvVars**](AppServicesAPI.md#ListAppServiceEnvVars) | **Get** /app-services/{id}/env-vars | List app service env vars
[**ListAppServiceHelmValues**](AppServicesAPI.md#ListAppServiceHelmValues) | **Get** /app-services/{id}/helm-values | List app service Helm values
[**ListAppServiceIntegrations**](AppServicesAPI.md#ListAppServiceIntegrations) | **Get** /app-services/{id}/integrations | List app service integrations
[**ListAppServiceLinks**](AppServicesAPI.md#ListAppServiceLinks) | **Get** /app-services/{id}/links | List app service links
[**ListAppServiceSettings**](AppServicesAPI.md#ListAppServiceSettings) | **Get** /app-services/{id}/settings | List app service settings
[**ListAppServiceTokens**](AppServicesAPI.md#ListAppServiceTokens) | **Get** /app-services/{id}/tokens | List app service tokens
[**ListAppServices**](AppServicesAPI.md#ListAppServices) | **Get** /app-services | List app services
[**RunAppServiceAction**](AppServicesAPI.md#RunAppServiceAction) | **Post** /app-services/{id}/actions/{name} | Run app service action
[**RunAppServiceCronSchedule**](AppServicesAPI.md#RunAppServiceCronSchedule) | **Post** /app-service-cron-schedules/{id}/run | Run app service cron schedule
[**SetAppServiceConfig**](AppServicesAPI.md#SetAppServiceConfig) | **Put** /app-services/{id}/configs/{name} | Set app service config
[**SetAppServiceLink**](AppServicesAPI.md#SetAppServiceLink) | **Put** /app-services/{id}/links/{name} | Set app service link
[**SetAppServiceResources**](AppServicesAPI.md#SetAppServiceResources) | **Put** /app-services/{id}/resources | Set app service resources
[**SetAppServiceSetting**](AppServicesAPI.md#SetAppServiceSetting) | **Put** /app-services/{id}/settings/{name} | Set app service setting
[**StartLogStream**](AppServicesAPI.md#StartLogStream) | **Post** /log-streams/{id}/start | Start log stream
[**StopLogStream**](AppServicesAPI.md#StopLogStream) | **Post** /log-streams/{id}/stop | Stop log stream
[**UpdateAppService**](AppServicesAPI.md#UpdateAppService) | **Put** /app-services/{id} | Update app service
[**UpdateAppServiceCronSchedule**](AppServicesAPI.md#UpdateAppServiceCronSchedule) | **Put** /app-service-cron-schedules/{id} | Update app service cron schedule
[**UpdateAppServiceDatabase**](AppServicesAPI.md#UpdateAppServiceDatabase) | **Put** /app-services/{id}/database | Update app service database references
[**UpdateAppServiceEnvVar**](AppServicesAPI.md#UpdateAppServiceEnvVar) | **Put** /app-service-env-vars/{id} | Update app service env var
[**UpdateAppServiceHelmValue**](AppServicesAPI.md#UpdateAppServiceHelmValue) | **Put** /app-service-helm-values/{id} | Update app service Helm value
[**UpdateAppServiceToken**](AppServicesAPI.md#UpdateAppServiceToken) | **Put** /app-service-tokens/{id} | Update app service token



## CreateAppServiceAnnotation

> AppServiceAnnotation CreateAppServiceAnnotation(ctx, id).NewAnnotationInput(newAnnotationInput).Execute()

Create app service annotation



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
	newAnnotationInput := *openapiclient.NewNewAnnotationInput("Name_example", "Value_example") // NewAnnotationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceAnnotation(context.Background(), id).NewAnnotationInput(newAnnotationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceAnnotation`: AppServiceAnnotation
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceAnnotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAnnotationInput** | [**NewAnnotationInput**](NewAnnotationInput.md) |  | 

### Return type

[**AppServiceAnnotation**](AppServiceAnnotation.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceCronSchedule

> AppServiceCronSchedule CreateAppServiceCronSchedule(ctx, id).NewAppServiceCronScheduleInput(newAppServiceCronScheduleInput).Execute()

Create app service cron schedule



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
	newAppServiceCronScheduleInput := *openapiclient.NewNewAppServiceCronScheduleInput("Title_example", "Crontab_example", "Command_example") // NewAppServiceCronScheduleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceCronSchedule(context.Background(), id).NewAppServiceCronScheduleInput(newAppServiceCronScheduleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceCronSchedule`: AppServiceCronSchedule
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceCronScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAppServiceCronScheduleInput** | [**NewAppServiceCronScheduleInput**](NewAppServiceCronScheduleInput.md) |  | 

### Return type

[**AppServiceCronSchedule**](AppServiceCronSchedule.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceEnvVar

> AppServiceEnvVar CreateAppServiceEnvVar(ctx, id).NewAppServiceEnvVarInput(newAppServiceEnvVarInput).Execute()

Create app service env var



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
	newAppServiceEnvVarInput := *openapiclient.NewNewAppServiceEnvVarInput("Name_example", "Value_example", false) // NewAppServiceEnvVarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceEnvVar(context.Background(), id).NewAppServiceEnvVarInput(newAppServiceEnvVarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceEnvVar`: AppServiceEnvVar
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAppServiceEnvVarInput** | [**NewAppServiceEnvVarInput**](NewAppServiceEnvVarInput.md) |  | 

### Return type

[**AppServiceEnvVar**](AppServiceEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceHelmValue

> AppServiceHelmValue CreateAppServiceHelmValue(ctx, id).NamedSecretValueInput(namedSecretValueInput).Execute()

Create app service Helm value



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
	namedSecretValueInput := *openapiclient.NewNamedSecretValueInput("Name_example", "Value_example", false) // NamedSecretValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceHelmValue(context.Background(), id).NamedSecretValueInput(namedSecretValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceHelmValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceHelmValue`: AppServiceHelmValue
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceHelmValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceHelmValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namedSecretValueInput** | [**NamedSecretValueInput**](NamedSecretValueInput.md) |  | 

### Return type

[**AppServiceHelmValue**](AppServiceHelmValue.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceIntegration

> AppServiceIntegration CreateAppServiceIntegration(ctx, id).IntegrationLinkInput(integrationLinkInput).Execute()

Create app service integration



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
	integrationLinkInput := *openapiclient.NewIntegrationLinkInput("Name_example", int32(123)) // IntegrationLinkInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceIntegration(context.Background(), id).IntegrationLinkInput(integrationLinkInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceIntegration`: AppServiceIntegration
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationLinkInput** | [**IntegrationLinkInput**](IntegrationLinkInput.md) |  | 

### Return type

[**AppServiceIntegration**](AppServiceIntegration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceLogStream

> LogStream CreateAppServiceLogStream(ctx, id).NewAppServiceLogStreamInput(newAppServiceLogStreamInput).Execute()

Create app service log stream



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
	newAppServiceLogStreamInput := *openapiclient.NewNewAppServiceLogStreamInput() // NewAppServiceLogStreamInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceLogStream(context.Background(), id).NewAppServiceLogStreamInput(newAppServiceLogStreamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceLogStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceLogStream`: LogStream
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceLogStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAppServiceLogStreamInput** | [**NewAppServiceLogStreamInput**](NewAppServiceLogStreamInput.md) |  | 

### Return type

[**LogStream**](LogStream.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceToken

> AppServiceToken CreateAppServiceToken(ctx, id).NamedSecretValueInput(namedSecretValueInput).Execute()

Create app service token



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
	namedSecretValueInput := *openapiclient.NewNamedSecretValueInput("Name_example", "Value_example", false) // NamedSecretValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.CreateAppServiceToken(context.Background(), id).NamedSecretValueInput(namedSecretValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.CreateAppServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceToken`: AppServiceToken
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.CreateAppServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namedSecretValueInput** | [**NamedSecretValueInput**](NamedSecretValueInput.md) |  | 

### Return type

[**AppServiceToken**](AppServiceToken.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppServiceAnnotation

> OperationResult DeleteAppServiceAnnotation(ctx, id).Execute()

Delete app service annotation



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
	resp, r, err := apiClient.AppServicesAPI.DeleteAppServiceAnnotation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.DeleteAppServiceAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceAnnotation`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.DeleteAppServiceAnnotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceAnnotationRequest struct via the builder pattern


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


## DeleteAppServiceCronSchedule

> OperationResult DeleteAppServiceCronSchedule(ctx, id).Execute()

Delete app service cron schedule



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
	resp, r, err := apiClient.AppServicesAPI.DeleteAppServiceCronSchedule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.DeleteAppServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceCronSchedule`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.DeleteAppServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceCronScheduleRequest struct via the builder pattern


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


## DeleteAppServiceEnvVar

> OperationResult DeleteAppServiceEnvVar(ctx, id).Execute()

Delete app service env var



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
	resp, r, err := apiClient.AppServicesAPI.DeleteAppServiceEnvVar(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.DeleteAppServiceEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceEnvVar`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.DeleteAppServiceEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceEnvVarRequest struct via the builder pattern


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


## DeleteAppServiceHelmValue

> OperationResult DeleteAppServiceHelmValue(ctx, id).Execute()

Delete app service Helm value



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
	resp, r, err := apiClient.AppServicesAPI.DeleteAppServiceHelmValue(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.DeleteAppServiceHelmValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceHelmValue`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.DeleteAppServiceHelmValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceHelmValueRequest struct via the builder pattern


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


## DeleteAppServiceIntegration

> OperationResult DeleteAppServiceIntegration(ctx, id).Execute()

Delete app service integration



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
	resp, r, err := apiClient.AppServicesAPI.DeleteAppServiceIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.DeleteAppServiceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceIntegration`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.DeleteAppServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceIntegrationRequest struct via the builder pattern


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


## DeleteAppServiceToken

> OperationResult DeleteAppServiceToken(ctx, id).Execute()

Delete app service token



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
	resp, r, err := apiClient.AppServicesAPI.DeleteAppServiceToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.DeleteAppServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceToken`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.DeleteAppServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceTokenRequest struct via the builder pattern


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


## GetAppService

> AppService GetAppService(ctx, id).Execute()

Get app service



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
	resp, r, err := apiClient.AppServicesAPI.GetAppService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.GetAppService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppService`: AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.GetAppService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppServiceCronJob

> AppServiceCronJob GetAppServiceCronJob(ctx, id).Execute()

Get app service cron job



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
	resp, r, err := apiClient.AppServicesAPI.GetAppServiceCronJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.GetAppServiceCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppServiceCronJob`: AppServiceCronJob
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.GetAppServiceCronJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppServiceCronJob**](AppServiceCronJob.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeepLogStreamAlive

> OperationResult KeepLogStreamAlive(ctx, id).Execute()

Keep log stream alive



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
	resp, r, err := apiClient.AppServicesAPI.KeepLogStreamAlive(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.KeepLogStreamAlive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeepLogStreamAlive`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.KeepLogStreamAlive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeepLogStreamAliveRequest struct via the builder pattern


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


## ListAppServiceAnnotations

> []AppServiceAnnotation ListAppServiceAnnotations(ctx, id).Execute()

List app service annotations



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceAnnotations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceAnnotations`: []AppServiceAnnotation
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceAnnotations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceAnnotation**](AppServiceAnnotation.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceConfigs

> []AppServiceConfig ListAppServiceConfigs(ctx, id).Execute()

List app service configs



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceConfigs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceConfigs`: []AppServiceConfig
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceConfig**](AppServiceConfig.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceContainers

> []AppServiceContainer ListAppServiceContainers(ctx, id).Execute()

List app service containers



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceContainers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceContainers`: []AppServiceContainer
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceContainer**](AppServiceContainer.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceCronJobs

> AppServiceCronJobsResponse ListAppServiceCronJobs(ctx).AppInstanceId(appInstanceId).AppServiceId(appServiceId).ScheduleId(scheduleId).Page(page).PageSize(pageSize).Execute()

List app service cron jobs



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
	scheduleId := int32(56) // int32 |  (optional)
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceCronJobs(context.Background()).AppInstanceId(appInstanceId).AppServiceId(appServiceId).ScheduleId(scheduleId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceCronJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceCronJobs`: AppServiceCronJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceCronJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceCronJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **appServiceId** | **int32** |  | 
 **scheduleId** | **int32** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**AppServiceCronJobsResponse**](AppServiceCronJobsResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceCronSchedules

> []AppServiceCronSchedule ListAppServiceCronSchedules(ctx, id).Execute()

List app service cron schedules



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceCronSchedules(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceCronSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceCronSchedules`: []AppServiceCronSchedule
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceCronSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceCronSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceCronSchedule**](AppServiceCronSchedule.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceEnvVars

> []AppServiceEnvVar ListAppServiceEnvVars(ctx, id).Execute()

List app service env vars



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceEnvVars(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceEnvVars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceEnvVars`: []AppServiceEnvVar
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceEnvVars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceEnvVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceEnvVar**](AppServiceEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceHelmValues

> []AppServiceHelmValue ListAppServiceHelmValues(ctx, id).Execute()

List app service Helm values



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceHelmValues(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceHelmValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceHelmValues`: []AppServiceHelmValue
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceHelmValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceHelmValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceHelmValue**](AppServiceHelmValue.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceIntegrations

> []AppServiceIntegration ListAppServiceIntegrations(ctx, id).Execute()

List app service integrations



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceIntegrations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceIntegrations`: []AppServiceIntegration
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceIntegration**](AppServiceIntegration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceLinks

> []AppServiceLink ListAppServiceLinks(ctx, id).Execute()

List app service links



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceLinks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceLinks`: []AppServiceLink
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceLink**](AppServiceLink.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceSettings

> []AppServiceSetting ListAppServiceSettings(ctx, id).Execute()

List app service settings



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceSettings`: []AppServiceSetting
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceSetting**](AppServiceSetting.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceTokens

> []AppServiceToken ListAppServiceTokens(ctx, id).Execute()

List app service tokens



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
	resp, r, err := apiClient.AppServicesAPI.ListAppServiceTokens(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServiceTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceTokens`: []AppServiceToken
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServiceTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppServiceToken**](AppServiceToken.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServices

> []AppService ListAppServices(ctx).AppInstanceId(appInstanceId).Execute()

List app services



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
	appInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.ListAppServices(context.Background()).AppInstanceId(appInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.ListAppServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServices`: []AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.ListAppServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 

### Return type

[**[]AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAppServiceAction

> OperationResult RunAppServiceAction(ctx, id, name).Execute()

Run app service action



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.RunAppServiceAction(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.RunAppServiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunAppServiceAction`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.RunAppServiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunAppServiceActionRequest struct via the builder pattern


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


## RunAppServiceCronSchedule

> Task RunAppServiceCronSchedule(ctx, id).Execute()

Run app service cron schedule



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
	resp, r, err := apiClient.AppServicesAPI.RunAppServiceCronSchedule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.RunAppServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunAppServiceCronSchedule`: Task
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.RunAppServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunAppServiceCronScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAppServiceConfig

> OperationResult SetAppServiceConfig(ctx, id, name).ConfigOverrideInput(configOverrideInput).Execute()

Set app service config



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
	name := "name_example" // string | 
	configOverrideInput := *openapiclient.NewConfigOverrideInput() // ConfigOverrideInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.SetAppServiceConfig(context.Background(), id, name).ConfigOverrideInput(configOverrideInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.SetAppServiceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAppServiceConfig`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.SetAppServiceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAppServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configOverrideInput** | [**ConfigOverrideInput**](ConfigOverrideInput.md) |  | 

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


## SetAppServiceLink

> OperationResult SetAppServiceLink(ctx, id, name).AppServiceLinkInput(appServiceLinkInput).Execute()

Set app service link



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
	name := "name_example" // string | 
	appServiceLinkInput := *openapiclient.NewAppServiceLinkInput() // AppServiceLinkInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.SetAppServiceLink(context.Background(), id, name).AppServiceLinkInput(appServiceLinkInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.SetAppServiceLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAppServiceLink`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.SetAppServiceLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAppServiceLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appServiceLinkInput** | [**AppServiceLinkInput**](AppServiceLinkInput.md) |  | 

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


## SetAppServiceResources

> OperationResult SetAppServiceResources(ctx, id).ResourcesInput(resourcesInput).Execute()

Set app service resources



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
	resourcesInput := *openapiclient.NewResourcesInput() // ResourcesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.SetAppServiceResources(context.Background(), id).ResourcesInput(resourcesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.SetAppServiceResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAppServiceResources`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.SetAppServiceResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAppServiceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcesInput** | [**ResourcesInput**](ResourcesInput.md) |  | 

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


## SetAppServiceSetting

> AppServiceSetting SetAppServiceSetting(ctx, id, name).SetStringValueInput(setStringValueInput).Execute()

Set app service setting



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
	name := "name_example" // string | 
	setStringValueInput := *openapiclient.NewSetStringValueInput("Value_example") // SetStringValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.SetAppServiceSetting(context.Background(), id, name).SetStringValueInput(setStringValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.SetAppServiceSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAppServiceSetting`: AppServiceSetting
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.SetAppServiceSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAppServiceSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setStringValueInput** | [**SetStringValueInput**](SetStringValueInput.md) |  | 

### Return type

[**AppServiceSetting**](AppServiceSetting.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartLogStream

> OperationResult StartLogStream(ctx, id).Execute()

Start log stream



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
	resp, r, err := apiClient.AppServicesAPI.StartLogStream(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.StartLogStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartLogStream`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.StartLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartLogStreamRequest struct via the builder pattern


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


## StopLogStream

> OperationResult StopLogStream(ctx, id).Execute()

Stop log stream



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
	resp, r, err := apiClient.AppServicesAPI.StopLogStream(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.StopLogStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopLogStream`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.StopLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopLogStreamRequest struct via the builder pattern


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


## UpdateAppService

> AppService UpdateAppService(ctx, id).AppServiceInput(appServiceInput).Execute()

Update app service



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
	appServiceInput := *openapiclient.NewAppServiceInput() // AppServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppService(context.Background(), id).AppServiceInput(appServiceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppService`: AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appServiceInput** | [**AppServiceInput**](AppServiceInput.md) |  | 

### Return type

[**AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceCronSchedule

> AppServiceCronSchedule UpdateAppServiceCronSchedule(ctx, id).UpdateAppServiceCronScheduleInput(updateAppServiceCronScheduleInput).Execute()

Update app service cron schedule



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
	updateAppServiceCronScheduleInput := *openapiclient.NewUpdateAppServiceCronScheduleInput() // UpdateAppServiceCronScheduleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppServiceCronSchedule(context.Background(), id).UpdateAppServiceCronScheduleInput(updateAppServiceCronScheduleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceCronSchedule`: AppServiceCronSchedule
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceCronScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppServiceCronScheduleInput** | [**UpdateAppServiceCronScheduleInput**](UpdateAppServiceCronScheduleInput.md) |  | 

### Return type

[**AppServiceCronSchedule**](AppServiceCronSchedule.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceDatabase

> AppService UpdateAppServiceDatabase(ctx, id).UpdateAppServiceDatabaseInput(updateAppServiceDatabaseInput).Execute()

Update app service database references



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
	updateAppServiceDatabaseInput := *openapiclient.NewUpdateAppServiceDatabaseInput() // UpdateAppServiceDatabaseInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppServiceDatabase(context.Background(), id).UpdateAppServiceDatabaseInput(updateAppServiceDatabaseInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppServiceDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceDatabase`: AppService
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppServiceDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppServiceDatabaseInput** | [**UpdateAppServiceDatabaseInput**](UpdateAppServiceDatabaseInput.md) |  | 

### Return type

[**AppService**](AppService.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceEnvVar

> AppServiceEnvVar UpdateAppServiceEnvVar(ctx, id).UpdateAppServiceEnvVarInput(updateAppServiceEnvVarInput).Execute()

Update app service env var



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
	updateAppServiceEnvVarInput := *openapiclient.NewUpdateAppServiceEnvVarInput(false) // UpdateAppServiceEnvVarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppServiceEnvVar(context.Background(), id).UpdateAppServiceEnvVarInput(updateAppServiceEnvVarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppServiceEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceEnvVar`: AppServiceEnvVar
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppServiceEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppServiceEnvVarInput** | [**UpdateAppServiceEnvVarInput**](UpdateAppServiceEnvVarInput.md) |  | 

### Return type

[**AppServiceEnvVar**](AppServiceEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceHelmValue

> AppServiceHelmValue UpdateAppServiceHelmValue(ctx, id).UpdateSecretValueInput(updateSecretValueInput).Execute()

Update app service Helm value



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
	updateSecretValueInput := *openapiclient.NewUpdateSecretValueInput("Value_example", false) // UpdateSecretValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppServiceHelmValue(context.Background(), id).UpdateSecretValueInput(updateSecretValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppServiceHelmValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceHelmValue`: AppServiceHelmValue
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppServiceHelmValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceHelmValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSecretValueInput** | [**UpdateSecretValueInput**](UpdateSecretValueInput.md) |  | 

### Return type

[**AppServiceHelmValue**](AppServiceHelmValue.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceToken

> AppServiceToken UpdateAppServiceToken(ctx, id).UpdateSecretValueInput(updateSecretValueInput).Execute()

Update app service token



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
	updateSecretValueInput := *openapiclient.NewUpdateSecretValueInput("Value_example", false) // UpdateSecretValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppServicesAPI.UpdateAppServiceToken(context.Background(), id).UpdateSecretValueInput(updateSecretValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppServicesAPI.UpdateAppServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceToken`: AppServiceToken
	fmt.Fprintf(os.Stdout, "Response from `AppServicesAPI.UpdateAppServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSecretValueInput** | [**UpdateSecretValueInput**](UpdateSecretValueInput.md) |  | 

### Return type

[**AppServiceToken**](AppServiceToken.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

