# \ProvidersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProviderFromManifest**](ProvidersAPI.md#CreateProviderFromManifest) | **Post** /providers/actions/create-from-manifest | Create provider from manifest
[**CreateVariableProvider**](ProvidersAPI.md#CreateVariableProvider) | **Post** /providers/actions/create-variable | Create variable provider
[**GetProvider**](ProvidersAPI.md#GetProvider) | **Get** /providers/{id} | Get provider
[**GetProviderByName**](ProvidersAPI.md#GetProviderByName) | **Get** /providers/by-name/{name} | Get provider by name
[**GetProviderRevision**](ProvidersAPI.md#GetProviderRevision) | **Get** /provider-revisions/{id} | Get provider revision
[**ImportProviders**](ProvidersAPI.md#ImportProviders) | **Post** /providers/actions/import | Import providers from Git
[**ListProviders**](ProvidersAPI.md#ListProviders) | **Get** /providers | List providers
[**UpdateProviderFromGit**](ProvidersAPI.md#UpdateProviderFromGit) | **Post** /providers/{id}/actions/update-from-git | Update provider from Git
[**UpdateProviderFromManifest**](ProvidersAPI.md#UpdateProviderFromManifest) | **Post** /providers/{id}/actions/update-from-manifest | Update provider from manifest
[**UpdateProviderSettings**](ProvidersAPI.md#UpdateProviderSettings) | **Put** /providers/settings/{id} | Update provider settings



## CreateProviderFromManifest

> Provider CreateProviderFromManifest(ctx).ProviderManifestInput(providerManifestInput).Execute()

Create provider from manifest



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
	providerManifestInput := *openapiclient.NewProviderManifestInput("ManifestYaml_example") // ProviderManifestInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.CreateProviderFromManifest(context.Background()).ProviderManifestInput(providerManifestInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.CreateProviderFromManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProviderFromManifest`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.CreateProviderFromManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderFromManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerManifestInput** | [**ProviderManifestInput**](ProviderManifestInput.md) |  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVariableProvider

> Provider CreateVariableProvider(ctx).NewVariableProviderInput(newVariableProviderInput).Execute()

Create variable provider



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
	newVariableProviderInput := *openapiclient.NewNewVariableProviderInput("Name_example", "Title_example", []openapiclient.NewVariableProviderFieldInput{*openapiclient.NewNewVariableProviderFieldInput("Name_example", "Label_example", "Variable_example", false, false)}) // NewVariableProviderInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.CreateVariableProvider(context.Background()).NewVariableProviderInput(newVariableProviderInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.CreateVariableProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVariableProvider`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.CreateVariableProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newVariableProviderInput** | [**NewVariableProviderInput**](NewVariableProviderInput.md) |  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvider

> Provider GetProvider(ctx, id).Execute()

Get provider



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
	resp, r, err := apiClient.ProvidersAPI.GetProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderByName

> Provider GetProviderByName(ctx, name).Execute()

Get provider by name



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
	resp, r, err := apiClient.ProvidersAPI.GetProviderByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProviderByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderByName`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProviderByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderRevision

> ProviderRevision GetProviderRevision(ctx, id).Execute()

Get provider revision



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
	resp, r, err := apiClient.ProvidersAPI.GetProviderRevision(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProviderRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderRevision`: ProviderRevision
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProviderRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderRevision**](ProviderRevision.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportProviders

> OperationResult ImportProviders(ctx).ImportCatalogFromGitInput(importCatalogFromGitInput).Execute()

Import providers from Git



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
	resp, r, err := apiClient.ProvidersAPI.ImportProviders(context.Background()).ImportCatalogFromGitInput(importCatalogFromGitInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ImportProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportProviders`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ImportProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportProvidersRequest struct via the builder pattern


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


## ListProviders

> ProvidersResponse ListProviders(ctx).OrgId(orgId).ProjectIds(projectIds).ExcludePublic(excludePublic).Search(search).Page(page).PageSize(pageSize).Execute()

List providers



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
	excludePublic := true // bool |  (optional)
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ListProviders(context.Background()).OrgId(orgId).ProjectIds(projectIds).ExcludePublic(excludePublic).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: ProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **excludePublic** | **bool** |  | 
 **search** | **string** |  | 
 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**ProvidersResponse**](ProvidersResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProviderFromGit

> OperationResult UpdateProviderFromGit(ctx, id).UpdateStackFromGitRequest(updateStackFromGitRequest).Execute()

Update provider from Git



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
	updateStackFromGitRequest := *openapiclient.NewUpdateStackFromGitRequest("GitRef_example", "GitRefType_example") // UpdateStackFromGitRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.UpdateProviderFromGit(context.Background(), id).UpdateStackFromGitRequest(updateStackFromGitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.UpdateProviderFromGit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProviderFromGit`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.UpdateProviderFromGit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderFromGitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStackFromGitRequest** | [**UpdateStackFromGitRequest**](UpdateStackFromGitRequest.md) |  | 

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


## UpdateProviderFromManifest

> OperationResult UpdateProviderFromManifest(ctx, id).ProviderManifestUpdateInput(providerManifestUpdateInput).Execute()

Update provider from manifest



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
	providerManifestUpdateInput := *openapiclient.NewProviderManifestUpdateInput("ManifestYaml_example") // ProviderManifestUpdateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.UpdateProviderFromManifest(context.Background(), id).ProviderManifestUpdateInput(providerManifestUpdateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.UpdateProviderFromManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProviderFromManifest`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.UpdateProviderFromManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderFromManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerManifestUpdateInput** | [**ProviderManifestUpdateInput**](ProviderManifestUpdateInput.md) |  | 

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


## UpdateProviderSettings

> Provider UpdateProviderSettings(ctx, id).ProviderSettingsInput(providerSettingsInput).Execute()

Update provider settings



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
	providerSettingsInput := *openapiclient.NewProviderSettingsInput(*openapiclient.NewGitAutoUpdateSettingsInput(false, false, false, false, false, false)) // ProviderSettingsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.UpdateProviderSettings(context.Background(), id).ProviderSettingsInput(providerSettingsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.UpdateProviderSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProviderSettings`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.UpdateProviderSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerSettingsInput** | [**ProviderSettingsInput**](ProviderSettingsInput.md) |  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

