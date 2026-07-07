# \ServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceFromManifest**](ServicesAPI.md#CreateServiceFromManifest) | **Post** /services/actions/create-from-manifest | Create service from manifest
[**GetService**](ServicesAPI.md#GetService) | **Get** /services/{id} | Get service
[**GetServiceByName**](ServicesAPI.md#GetServiceByName) | **Get** /services/by-name/{name} | Get service by name
[**GetServiceRevision**](ServicesAPI.md#GetServiceRevision) | **Get** /service-revisions/{id} | Get service revision
[**ImportServices**](ServicesAPI.md#ImportServices) | **Post** /services/actions/import | Import services from Git
[**ListServiceLinkCandidates**](ServicesAPI.md#ListServiceLinkCandidates) | **Get** /services/{name}/options/link-candidates | List service link candidates
[**ListServices**](ServicesAPI.md#ListServices) | **Get** /services | List services
[**ScaffoldServiceFromHelmChart**](ServicesAPI.md#ScaffoldServiceFromHelmChart) | **Post** /services/actions/scaffold-from-helm-chart | Scaffold service from Helm chart
[**UpdateServiceSettings**](ServicesAPI.md#UpdateServiceSettings) | **Put** /services/settings/{id} | Update service settings
[**ValidateServiceManifest**](ServicesAPI.md#ValidateServiceManifest) | **Post** /services/actions/validate-manifest | Validate service manifest



## CreateServiceFromManifest

> Service CreateServiceFromManifest(ctx).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()

Create service from manifest



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
	manifestFromYAMLInput := *openapiclient.NewManifestFromYAMLInput("ManifestYaml_example") // ManifestFromYAMLInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.CreateServiceFromManifest(context.Background()).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.CreateServiceFromManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceFromManifest`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.CreateServiceFromManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceFromManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manifestFromYAMLInput** | [**ManifestFromYAMLInput**](ManifestFromYAMLInput.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> Service GetService(ctx, id).Execute()

Get service



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
	resp, r, err := apiClient.ServicesAPI.GetService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Service**](Service.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceByName

> Service GetServiceByName(ctx, name).RevNumber(revNumber).Execute()

Get service by name



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
	revNumber := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.GetServiceByName(context.Background(), name).RevNumber(revNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetServiceByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceByName`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetServiceByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revNumber** | **int32** |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceRevision

> ServiceRevision GetServiceRevision(ctx, id).Execute()

Get service revision



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
	resp, r, err := apiClient.ServicesAPI.GetServiceRevision(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetServiceRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceRevision`: ServiceRevision
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetServiceRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceRevision**](ServiceRevision.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportServices

> OperationResult ImportServices(ctx).ImportCatalogFromGitInput(importCatalogFromGitInput).Execute()

Import services from Git



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
	importCatalogFromGitInput := *openapiclient.NewImportCatalogFromGitInput(int32(123), "RemoteGitRepoId_example", "GitRef_example", "GitRefType_example") // ImportCatalogFromGitInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ImportServices(context.Background()).ImportCatalogFromGitInput(importCatalogFromGitInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ImportServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportServices`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ImportServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importCatalogFromGitInput** | [**ImportCatalogFromGitInput**](ImportCatalogFromGitInput.md) |  | 

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


## ListServiceLinkCandidates

> []map[string]interface{} ListServiceLinkCandidates(ctx, name).Execute()

List service link candidates



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ListServiceLinkCandidates(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ListServiceLinkCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceLinkCandidates`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ListServiceLinkCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceLinkCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListServices

> ServicesResponse ListServices(ctx).OrgId(orgId).ProjectIds(projectIds).Search(search).Page(page).PageSize(pageSize).Execute()

List services



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
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ListServices(context.Background()).OrgId(orgId).ProjectIds(projectIds).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ListServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServices`: ServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ListServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **search** | **string** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**ServicesResponse**](ServicesResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaffoldServiceFromHelmChart

> HelmChartServiceScaffoldResponse ScaffoldServiceFromHelmChart(ctx).HelmChartServiceScaffoldInput(helmChartServiceScaffoldInput).Execute()

Scaffold service from Helm chart



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
	helmChartServiceScaffoldInput := *openapiclient.NewHelmChartServiceScaffoldInput(*openapiclient.NewHelmChartInput("Chart_example")) // HelmChartServiceScaffoldInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ScaffoldServiceFromHelmChart(context.Background()).HelmChartServiceScaffoldInput(helmChartServiceScaffoldInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ScaffoldServiceFromHelmChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScaffoldServiceFromHelmChart`: HelmChartServiceScaffoldResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ScaffoldServiceFromHelmChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScaffoldServiceFromHelmChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmChartServiceScaffoldInput** | [**HelmChartServiceScaffoldInput**](HelmChartServiceScaffoldInput.md) |  | 

### Return type

[**HelmChartServiceScaffoldResponse**](HelmChartServiceScaffoldResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceSettings

> Service UpdateServiceSettings(ctx, id).ServiceSettingsInput(serviceSettingsInput).Execute()

Update service settings



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
	serviceSettingsInput := *openapiclient.NewServiceSettingsInput() // ServiceSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.UpdateServiceSettings(context.Background(), id).ServiceSettingsInput(serviceSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.UpdateServiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceSettings`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.UpdateServiceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceSettingsInput** | [**ServiceSettingsInput**](ServiceSettingsInput.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateServiceManifest

> ManifestValidationResponse ValidateServiceManifest(ctx).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()

Validate service manifest



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
	manifestFromYAMLInput := *openapiclient.NewManifestFromYAMLInput("ManifestYaml_example") // ManifestFromYAMLInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ValidateServiceManifest(context.Background()).ManifestFromYAMLInput(manifestFromYAMLInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ValidateServiceManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateServiceManifest`: ManifestValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ValidateServiceManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateServiceManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manifestFromYAMLInput** | [**ManifestFromYAMLInput**](ManifestFromYAMLInput.md) |  | 

### Return type

[**ManifestValidationResponse**](ManifestValidationResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

