# \StackServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStackService**](StackServicesAPI.md#CreateStackService) | **Post** /stack-services | Create stack service
[**CreateStackServiceAnnotation**](StackServicesAPI.md#CreateStackServiceAnnotation) | **Post** /stack-services/{id}/annotations | Create stack service annotation
[**CreateStackServiceCronSchedule**](StackServicesAPI.md#CreateStackServiceCronSchedule) | **Post** /stack-services/{id}/cron-schedules | Create stack service cron schedule
[**CreateStackServiceEnvVar**](StackServicesAPI.md#CreateStackServiceEnvVar) | **Post** /stack-services/{id}/env-vars | Create stack service env var
[**CreateStackServiceHelmValue**](StackServicesAPI.md#CreateStackServiceHelmValue) | **Post** /stack-services/{id}/helm-values | Create stack service Helm value
[**CreateStackServiceIntegration**](StackServicesAPI.md#CreateStackServiceIntegration) | **Post** /stack-services/{id}/integrations | Create stack service integration
[**CreateStackServiceToken**](StackServicesAPI.md#CreateStackServiceToken) | **Post** /stack-services/{id}/tokens | Create stack service token
[**DeleteStackService**](StackServicesAPI.md#DeleteStackService) | **Delete** /stack-services/{id} | Delete stack service
[**DeleteStackServiceAnnotation**](StackServicesAPI.md#DeleteStackServiceAnnotation) | **Delete** /stack-service-annotations/{id} | Delete stack service annotation
[**DeleteStackServiceCronSchedule**](StackServicesAPI.md#DeleteStackServiceCronSchedule) | **Delete** /stack-service-cron-schedules/{id} | Delete stack service cron schedule
[**DeleteStackServiceEnvVar**](StackServicesAPI.md#DeleteStackServiceEnvVar) | **Delete** /stack-service-env-vars/{id} | Delete stack service env var
[**DeleteStackServiceHelmValue**](StackServicesAPI.md#DeleteStackServiceHelmValue) | **Delete** /stack-service-helm-values/{id} | Delete stack service Helm value
[**DeleteStackServiceIntegration**](StackServicesAPI.md#DeleteStackServiceIntegration) | **Delete** /stack-service-integrations/{id} | Delete stack service integration
[**DeleteStackServiceToken**](StackServicesAPI.md#DeleteStackServiceToken) | **Delete** /stack-service-tokens/{id} | Delete stack service token
[**ListStackServiceAnnotations**](StackServicesAPI.md#ListStackServiceAnnotations) | **Get** /stack-services/{id}/annotations | List stack service annotations
[**ListStackServiceConfigs**](StackServicesAPI.md#ListStackServiceConfigs) | **Get** /stack-services/{id}/configs | List stack service configs
[**ListStackServiceCronSchedules**](StackServicesAPI.md#ListStackServiceCronSchedules) | **Get** /stack-services/{id}/cron-schedules | List stack service cron schedules
[**ListStackServiceEnvVars**](StackServicesAPI.md#ListStackServiceEnvVars) | **Get** /stack-services/{id}/env-vars | List stack service env vars
[**ListStackServiceHelmValues**](StackServicesAPI.md#ListStackServiceHelmValues) | **Get** /stack-services/{id}/helm-values | List stack service Helm values
[**ListStackServiceIntegrations**](StackServicesAPI.md#ListStackServiceIntegrations) | **Get** /stack-services/{id}/integrations | List stack service integrations
[**ListStackServiceLinks**](StackServicesAPI.md#ListStackServiceLinks) | **Get** /stack-services/{id}/links | List stack service links
[**ListStackServiceTokens**](StackServicesAPI.md#ListStackServiceTokens) | **Get** /stack-services/{id}/tokens | List stack service tokens
[**ListStackServiceVolumes**](StackServicesAPI.md#ListStackServiceVolumes) | **Get** /stack-services/{id}/volumes | List stack service volumes
[**ListStackServices**](StackServicesAPI.md#ListStackServices) | **Get** /stack-services | List stack services
[**SetStackServiceConfig**](StackServicesAPI.md#SetStackServiceConfig) | **Put** /stack-services/{id}/configs/{name} | Set stack service config
[**SetStackServiceLink**](StackServicesAPI.md#SetStackServiceLink) | **Put** /stack-services/{id}/links/{name} | Set stack service link
[**SetStackServiceOptions**](StackServicesAPI.md#SetStackServiceOptions) | **Put** /stack-services/{id}/options | Update stack service options
[**SetStackServiceResources**](StackServicesAPI.md#SetStackServiceResources) | **Put** /stack-services/{id}/resources | Set stack service resources
[**SetStackServiceSetting**](StackServicesAPI.md#SetStackServiceSetting) | **Put** /stack-services/{id}/settings/{name} | Set stack service setting
[**SetStackServiceVolume**](StackServicesAPI.md#SetStackServiceVolume) | **Put** /stack-services/{id}/volumes/{name} | Set stack service volume
[**UpdateStackService**](StackServicesAPI.md#UpdateStackService) | **Put** /stack-services/{id} | Update stack service
[**UpdateStackServiceCronSchedule**](StackServicesAPI.md#UpdateStackServiceCronSchedule) | **Put** /stack-service-cron-schedules/{id} | Update stack service cron schedule
[**UpdateStackServiceEnvVar**](StackServicesAPI.md#UpdateStackServiceEnvVar) | **Put** /stack-service-env-vars/{id} | Update stack service env var
[**UpdateStackServiceHelmValue**](StackServicesAPI.md#UpdateStackServiceHelmValue) | **Put** /stack-service-helm-values/{id} | Update stack service Helm value
[**UpdateStackServiceToken**](StackServicesAPI.md#UpdateStackServiceToken) | **Put** /stack-service-tokens/{id} | Update stack service token



## CreateStackService

> StackService CreateStackService(ctx).CreateStackServiceInput(createStackServiceInput).Execute()

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
	createStackServiceInput := *openapiclient.NewCreateStackServiceInput(int32(123), int32(123), "Name_example", "Title_example", false, int32(123)) // CreateStackServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackService(context.Background()).CreateStackServiceInput(createStackServiceInput).Execute()
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
 **createStackServiceInput** | [**CreateStackServiceInput**](CreateStackServiceInput.md) |  | 

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


## CreateStackServiceAnnotation

> StackServiceAnnotation CreateStackServiceAnnotation(ctx, id).NewStackServiceAnnotationInput(newStackServiceAnnotationInput).Execute()

Create stack service annotation



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
	newStackServiceAnnotationInput := *openapiclient.NewNewStackServiceAnnotationInput("Name_example", "Value_example") // NewStackServiceAnnotationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackServiceAnnotation(context.Background(), id).NewStackServiceAnnotationInput(newStackServiceAnnotationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackServiceAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackServiceAnnotation`: StackServiceAnnotation
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackServiceAnnotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStackServiceAnnotationInput** | [**NewStackServiceAnnotationInput**](NewStackServiceAnnotationInput.md) |  | 

### Return type

[**StackServiceAnnotation**](StackServiceAnnotation.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackServiceCronSchedule

> StackServiceCronSchedule CreateStackServiceCronSchedule(ctx, id).NewStackServiceCronScheduleInput(newStackServiceCronScheduleInput).Execute()

Create stack service cron schedule



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
	newStackServiceCronScheduleInput := *openapiclient.NewNewStackServiceCronScheduleInput("Name_example", "Title_example", "Crontab_example", "Command_example") // NewStackServiceCronScheduleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackServiceCronSchedule(context.Background(), id).NewStackServiceCronScheduleInput(newStackServiceCronScheduleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackServiceCronSchedule`: StackServiceCronSchedule
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceCronScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStackServiceCronScheduleInput** | [**NewStackServiceCronScheduleInput**](NewStackServiceCronScheduleInput.md) |  | 

### Return type

[**StackServiceCronSchedule**](StackServiceCronSchedule.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackServiceEnvVar

> StackServiceEnvVar CreateStackServiceEnvVar(ctx, id).NewStackServiceEnvVarInput(newStackServiceEnvVarInput).Execute()

Create stack service env var



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
	newStackServiceEnvVarInput := *openapiclient.NewNewStackServiceEnvVarInput("Name_example", "Value_example", false) // NewStackServiceEnvVarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackServiceEnvVar(context.Background(), id).NewStackServiceEnvVarInput(newStackServiceEnvVarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackServiceEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackServiceEnvVar`: StackServiceEnvVar
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackServiceEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStackServiceEnvVarInput** | [**NewStackServiceEnvVarInput**](NewStackServiceEnvVarInput.md) |  | 

### Return type

[**StackServiceEnvVar**](StackServiceEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackServiceHelmValue

> StackServiceHelmValue CreateStackServiceHelmValue(ctx, id).NewStackServiceScopedValueInput(newStackServiceScopedValueInput).Execute()

Create stack service Helm value



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
	newStackServiceScopedValueInput := *openapiclient.NewNewStackServiceScopedValueInput("Name_example", "Value_example", false) // NewStackServiceScopedValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackServiceHelmValue(context.Background(), id).NewStackServiceScopedValueInput(newStackServiceScopedValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackServiceHelmValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackServiceHelmValue`: StackServiceHelmValue
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackServiceHelmValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceHelmValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStackServiceScopedValueInput** | [**NewStackServiceScopedValueInput**](NewStackServiceScopedValueInput.md) |  | 

### Return type

[**StackServiceHelmValue**](StackServiceHelmValue.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackServiceIntegration

> StackServiceIntegration CreateStackServiceIntegration(ctx, id).IntegrationLinkInput(integrationLinkInput).Execute()

Create stack service integration



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
	resp, r, err := apiClient.StackServicesAPI.CreateStackServiceIntegration(context.Background(), id).IntegrationLinkInput(integrationLinkInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackServiceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackServiceIntegration`: StackServiceIntegration
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationLinkInput** | [**IntegrationLinkInput**](IntegrationLinkInput.md) |  | 

### Return type

[**StackServiceIntegration**](StackServiceIntegration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackServiceToken

> StackServiceToken CreateStackServiceToken(ctx, id).NewStackServiceTokenInput(newStackServiceTokenInput).Execute()

Create stack service token



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
	newStackServiceTokenInput := *openapiclient.NewNewStackServiceTokenInput("Name_example", false) // NewStackServiceTokenInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.CreateStackServiceToken(context.Background(), id).NewStackServiceTokenInput(newStackServiceTokenInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.CreateStackServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackServiceToken`: StackServiceToken
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.CreateStackServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStackServiceTokenInput** | [**NewStackServiceTokenInput**](NewStackServiceTokenInput.md) |  | 

### Return type

[**StackServiceToken**](StackServiceToken.md)

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


## DeleteStackServiceAnnotation

> OperationResult DeleteStackServiceAnnotation(ctx, id).Execute()

Delete stack service annotation



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackServiceAnnotation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackServiceAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackServiceAnnotation`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackServiceAnnotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceAnnotationRequest struct via the builder pattern


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


## DeleteStackServiceCronSchedule

> OperationResult DeleteStackServiceCronSchedule(ctx, id).Execute()

Delete stack service cron schedule



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackServiceCronSchedule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackServiceCronSchedule`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceCronScheduleRequest struct via the builder pattern


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


## DeleteStackServiceEnvVar

> OperationResult DeleteStackServiceEnvVar(ctx, id).Execute()

Delete stack service env var



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackServiceEnvVar(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackServiceEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackServiceEnvVar`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackServiceEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceEnvVarRequest struct via the builder pattern


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


## DeleteStackServiceHelmValue

> OperationResult DeleteStackServiceHelmValue(ctx, id).Execute()

Delete stack service Helm value



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackServiceHelmValue(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackServiceHelmValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackServiceHelmValue`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackServiceHelmValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceHelmValueRequest struct via the builder pattern


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


## DeleteStackServiceIntegration

> OperationResult DeleteStackServiceIntegration(ctx, id).Execute()

Delete stack service integration



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackServiceIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackServiceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackServiceIntegration`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceIntegrationRequest struct via the builder pattern


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


## DeleteStackServiceToken

> OperationResult DeleteStackServiceToken(ctx, id).Execute()

Delete stack service token



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
	resp, r, err := apiClient.StackServicesAPI.DeleteStackServiceToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.DeleteStackServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackServiceToken`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.DeleteStackServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackServiceTokenRequest struct via the builder pattern


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


## ListStackServiceAnnotations

> []StackServiceAnnotation ListStackServiceAnnotations(ctx, id).Execute()

List stack service annotations



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceAnnotations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceAnnotations`: []StackServiceAnnotation
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceAnnotations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceAnnotation**](StackServiceAnnotation.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceConfigs

> []StackServiceConfig ListStackServiceConfigs(ctx, id).Execute()

List stack service configs



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceConfigs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceConfigs`: []StackServiceConfig
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceConfig**](StackServiceConfig.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceCronSchedules

> []StackServiceCronSchedule ListStackServiceCronSchedules(ctx, id).Execute()

List stack service cron schedules



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceCronSchedules(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceCronSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceCronSchedules`: []StackServiceCronSchedule
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceCronSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceCronSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceCronSchedule**](StackServiceCronSchedule.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceEnvVars

> []StackServiceEnvVar ListStackServiceEnvVars(ctx, id).Execute()

List stack service env vars



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceEnvVars(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceEnvVars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceEnvVars`: []StackServiceEnvVar
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceEnvVars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceEnvVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceEnvVar**](StackServiceEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceHelmValues

> []StackServiceHelmValue ListStackServiceHelmValues(ctx, id).Execute()

List stack service Helm values



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceHelmValues(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceHelmValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceHelmValues`: []StackServiceHelmValue
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceHelmValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceHelmValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceHelmValue**](StackServiceHelmValue.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceIntegrations

> []StackServiceIntegration ListStackServiceIntegrations(ctx, id).Execute()

List stack service integrations



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceIntegrations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceIntegrations`: []StackServiceIntegration
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceIntegration**](StackServiceIntegration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceLinks

> []StackServiceLink ListStackServiceLinks(ctx, id).Execute()

List stack service links



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceLinks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceLinks`: []StackServiceLink
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceLink**](StackServiceLink.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceTokens

> []StackServiceToken ListStackServiceTokens(ctx, id).Execute()

List stack service tokens



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceTokens(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceTokens`: []StackServiceToken
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceToken**](StackServiceToken.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackServiceVolumes

> []StackServiceVolume ListStackServiceVolumes(ctx, id).Execute()

List stack service volumes



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
	resp, r, err := apiClient.StackServicesAPI.ListStackServiceVolumes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.ListStackServiceVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStackServiceVolumes`: []StackServiceVolume
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.ListStackServiceVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackServiceVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackServiceVolume**](StackServiceVolume.md)

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


## SetStackServiceConfig

> OperationResult SetStackServiceConfig(ctx, id, name).StackServiceConfigInput(stackServiceConfigInput).Execute()

Set stack service config



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
	stackServiceConfigInput := *openapiclient.NewStackServiceConfigInput("Config_example") // StackServiceConfigInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.SetStackServiceConfig(context.Background(), id, name).StackServiceConfigInput(stackServiceConfigInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.SetStackServiceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStackServiceConfig`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.SetStackServiceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStackServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stackServiceConfigInput** | [**StackServiceConfigInput**](StackServiceConfigInput.md) |  | 

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


## SetStackServiceLink

> OperationResult SetStackServiceLink(ctx, id, name).StackServiceLinkInput(stackServiceLinkInput).Execute()

Set stack service link



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
	stackServiceLinkInput := *openapiclient.NewStackServiceLinkInput() // StackServiceLinkInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.SetStackServiceLink(context.Background(), id, name).StackServiceLinkInput(stackServiceLinkInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.SetStackServiceLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStackServiceLink`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.SetStackServiceLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStackServiceLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stackServiceLinkInput** | [**StackServiceLinkInput**](StackServiceLinkInput.md) |  | 

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


## SetStackServiceOptions

> OperationResult SetStackServiceOptions(ctx, id).StackServiceOptionsInput(stackServiceOptionsInput).Execute()

Update stack service options



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
	stackServiceOptionsInput := *openapiclient.NewStackServiceOptionsInput([]openapiclient.StackServiceOptionInput{*openapiclient.NewStackServiceOptionInput("Version_example", false, false)}) // StackServiceOptionsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.SetStackServiceOptions(context.Background(), id).StackServiceOptionsInput(stackServiceOptionsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.SetStackServiceOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStackServiceOptions`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.SetStackServiceOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStackServiceOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackServiceOptionsInput** | [**StackServiceOptionsInput**](StackServiceOptionsInput.md) |  | 

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


## SetStackServiceResources

> OperationResult SetStackServiceResources(ctx, id).ResourcesInput(resourcesInput).Execute()

Set stack service resources



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
	resp, r, err := apiClient.StackServicesAPI.SetStackServiceResources(context.Background(), id).ResourcesInput(resourcesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.SetStackServiceResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStackServiceResources`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.SetStackServiceResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStackServiceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcesInput** | [**ResourcesInput**](ResourcesInput.md) |  | 

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


## SetStackServiceSetting

> OperationResult SetStackServiceSetting(ctx, id, name).SetNullableStringValueInput(setNullableStringValueInput).Execute()

Set stack service setting



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
	setNullableStringValueInput := *openapiclient.NewSetNullableStringValueInput() // SetNullableStringValueInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.SetStackServiceSetting(context.Background(), id, name).SetNullableStringValueInput(setNullableStringValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.SetStackServiceSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStackServiceSetting`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.SetStackServiceSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStackServiceSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNullableStringValueInput** | [**SetNullableStringValueInput**](SetNullableStringValueInput.md) |  | 

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


## SetStackServiceVolume

> OperationResult SetStackServiceVolume(ctx, id, name).StackServiceVolumeInput(stackServiceVolumeInput).Execute()

Set stack service volume



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
	stackServiceVolumeInput := *openapiclient.NewStackServiceVolumeInput() // StackServiceVolumeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.SetStackServiceVolume(context.Background(), id, name).StackServiceVolumeInput(stackServiceVolumeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.SetStackServiceVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStackServiceVolume`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.SetStackServiceVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStackServiceVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stackServiceVolumeInput** | [**StackServiceVolumeInput**](StackServiceVolumeInput.md) |  | 

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


## UpdateStackServiceCronSchedule

> StackServiceCronSchedule UpdateStackServiceCronSchedule(ctx, id).UpdateStackServiceCronScheduleInput(updateStackServiceCronScheduleInput).Execute()

Update stack service cron schedule



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
	updateStackServiceCronScheduleInput := *openapiclient.NewUpdateStackServiceCronScheduleInput() // UpdateStackServiceCronScheduleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.UpdateStackServiceCronSchedule(context.Background(), id).UpdateStackServiceCronScheduleInput(updateStackServiceCronScheduleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.UpdateStackServiceCronSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackServiceCronSchedule`: StackServiceCronSchedule
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.UpdateStackServiceCronSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackServiceCronScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackServiceCronScheduleInput** | [**UpdateStackServiceCronScheduleInput**](UpdateStackServiceCronScheduleInput.md) |  | 

### Return type

[**StackServiceCronSchedule**](StackServiceCronSchedule.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackServiceEnvVar

> StackServiceEnvVar UpdateStackServiceEnvVar(ctx, id).UpdateStackServiceEnvVarInput(updateStackServiceEnvVarInput).Execute()

Update stack service env var



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
	updateStackServiceEnvVarInput := *openapiclient.NewUpdateStackServiceEnvVarInput("Value_example", false) // UpdateStackServiceEnvVarInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.UpdateStackServiceEnvVar(context.Background(), id).UpdateStackServiceEnvVarInput(updateStackServiceEnvVarInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.UpdateStackServiceEnvVar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackServiceEnvVar`: StackServiceEnvVar
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.UpdateStackServiceEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackServiceEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackServiceEnvVarInput** | [**UpdateStackServiceEnvVarInput**](UpdateStackServiceEnvVarInput.md) |  | 

### Return type

[**StackServiceEnvVar**](StackServiceEnvVar.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackServiceHelmValue

> StackServiceHelmValue UpdateStackServiceHelmValue(ctx, id).UpdateSecretValueInput(updateSecretValueInput).Execute()

Update stack service Helm value



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
	resp, r, err := apiClient.StackServicesAPI.UpdateStackServiceHelmValue(context.Background(), id).UpdateSecretValueInput(updateSecretValueInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.UpdateStackServiceHelmValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackServiceHelmValue`: StackServiceHelmValue
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.UpdateStackServiceHelmValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackServiceHelmValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSecretValueInput** | [**UpdateSecretValueInput**](UpdateSecretValueInput.md) |  | 

### Return type

[**StackServiceHelmValue**](StackServiceHelmValue.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackServiceToken

> StackServiceToken UpdateStackServiceToken(ctx, id).UpdateStackServiceTokenInput(updateStackServiceTokenInput).Execute()

Update stack service token



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
	updateStackServiceTokenInput := *openapiclient.NewUpdateStackServiceTokenInput(false) // UpdateStackServiceTokenInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackServicesAPI.UpdateStackServiceToken(context.Background(), id).UpdateStackServiceTokenInput(updateStackServiceTokenInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackServicesAPI.UpdateStackServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackServiceToken`: StackServiceToken
	fmt.Fprintf(os.Stdout, "Response from `StackServicesAPI.UpdateStackServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackServiceTokenInput** | [**UpdateStackServiceTokenInput**](UpdateStackServiceTokenInput.md) |  | 

### Return type

[**StackServiceToken**](StackServiceToken.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

