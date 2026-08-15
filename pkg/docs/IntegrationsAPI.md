# \IntegrationsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureIntegration**](IntegrationsAPI.md#ConfigureIntegration) | **Put** /integrations/configuration/{id} | Configure integration
[**CreateIntegration**](IntegrationsAPI.md#CreateIntegration) | **Post** /integrations | Create integration
[**DeleteIntegration**](IntegrationsAPI.md#DeleteIntegration) | **Delete** /integrations/{id} | Delete integration
[**GetAppAccessProviderOptions**](IntegrationsAPI.md#GetAppAccessProviderOptions) | **Get** /integrations/{id}/options/app-access | Get app-access provider options
[**GetIntegration**](IntegrationsAPI.md#GetIntegration) | **Get** /integrations/{id} | Get integration
[**GetIntegrationByName**](IntegrationsAPI.md#GetIntegrationByName) | **Get** /integrations/by-name/{name} | Get integration by name
[**GetIntegrationKubeSettings**](IntegrationsAPI.md#GetIntegrationKubeSettings) | **Get** /integrations/{id}/options/kube-settings | Get Kubernetes settings
[**GetIntegrationRemoteGitRepoFilePresence**](IntegrationsAPI.md#GetIntegrationRemoteGitRepoFilePresence) | **Get** /integrations/{id}/options/remote-git-repo-file | Check a remote Git repository file
[**ListIntegrationKubeMachineTypes**](IntegrationsAPI.md#ListIntegrationKubeMachineTypes) | **Get** /integrations/{id}/options/kube-machine-types | List Kubernetes machine types
[**ListIntegrationKubeRegions**](IntegrationsAPI.md#ListIntegrationKubeRegions) | **Get** /integrations/{id}/options/kube-regions | List Kubernetes regions
[**ListIntegrationKubeVersions**](IntegrationsAPI.md#ListIntegrationKubeVersions) | **Get** /integrations/{id}/options/kube-versions | List Kubernetes versions
[**ListIntegrationKubeZones**](IntegrationsAPI.md#ListIntegrationKubeZones) | **Get** /integrations/{id}/options/kube-zones | List Kubernetes zones
[**ListIntegrationRemoteGitRepoBranches**](IntegrationsAPI.md#ListIntegrationRemoteGitRepoBranches) | **Get** /integrations/{id}/options/remote-git-repo-branches | List remote Git repository branches
[**ListIntegrationRemoteGitRepoTags**](IntegrationsAPI.md#ListIntegrationRemoteGitRepoTags) | **Get** /integrations/{id}/options/remote-git-repo-tags | List remote Git repository tags
[**ListIntegrationRemoteGitRepos**](IntegrationsAPI.md#ListIntegrationRemoteGitRepos) | **Get** /integrations/{id}/options/remote-git-repos | List remote Git repositories
[**ListIntegrationScopes**](IntegrationsAPI.md#ListIntegrationScopes) | **Get** /integrations/{id}/options/scopes | List integration scopes
[**ListIntegrationStorageBuckets**](IntegrationsAPI.md#ListIntegrationStorageBuckets) | **Get** /integrations/{id}/options/storage-buckets | List storage buckets
[**ListIntegrationStorageClasses**](IntegrationsAPI.md#ListIntegrationStorageClasses) | **Get** /integrations/{id}/options/storage-classes | List storage classes
[**ListIntegrations**](IntegrationsAPI.md#ListIntegrations) | **Get** /integrations | List integrations
[**ResolveIntegration**](IntegrationsAPI.md#ResolveIntegration) | **Post** /integrations/actions/resolve | Resolve or create integration
[**TestIntegrationPermissions**](IntegrationsAPI.md#TestIntegrationPermissions) | **Post** /integrations/{id}/actions/test-permissions | Test integration permissions
[**UpdateIntegration**](IntegrationsAPI.md#UpdateIntegration) | **Put** /integrations/{id} | Update integration
[**ValidateAppAccessHostname**](IntegrationsAPI.md#ValidateAppAccessHostname) | **Post** /integrations/{id}/actions/validate-app-access-hostname | Validate an app-access hostname



## ConfigureIntegration

> IntegrationConfigurationResult ConfigureIntegration(ctx, id).UpdateIntegrationInput(updateIntegrationInput).Execute()

Configure integration



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
	updateIntegrationInput := *openapiclient.NewUpdateIntegrationInput("Title_example", "Name_example", []string{"Kinds_example"}) // UpdateIntegrationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ConfigureIntegration(context.Background(), id).UpdateIntegrationInput(updateIntegrationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ConfigureIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureIntegration`: IntegrationConfigurationResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ConfigureIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIntegrationInput** | [**UpdateIntegrationInput**](UpdateIntegrationInput.md) |  | 

### Return type

[**IntegrationConfigurationResult**](IntegrationConfigurationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIntegration

> Integration CreateIntegration(ctx).NewIntegrationInput(newIntegrationInput).Execute()

Create integration



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
	newIntegrationInput := *openapiclient.NewNewIntegrationInput(int32(123), "Name_example", "Title_example", []string{"Kinds_example"}) // NewIntegrationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CreateIntegration(context.Background()).NewIntegrationInput(newIntegrationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CreateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIntegration`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CreateIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newIntegrationInput** | [**NewIntegrationInput**](NewIntegrationInput.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegration

> OperationResult DeleteIntegration(ctx, id).Execute()

Delete integration



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
	resp, r, err := apiClient.IntegrationsAPI.DeleteIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.DeleteIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntegration`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.DeleteIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationRequest struct via the builder pattern


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


## GetAppAccessProviderOptions

> AppAccessProviderOptions GetAppAccessProviderOptions(ctx, id).Execute()

Get app-access provider options



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
	resp, r, err := apiClient.IntegrationsAPI.GetAppAccessProviderOptions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetAppAccessProviderOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppAccessProviderOptions`: AppAccessProviderOptions
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetAppAccessProviderOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppAccessProviderOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppAccessProviderOptions**](AppAccessProviderOptions.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegration

> Integration GetIntegration(ctx, id).Execute()

Get integration



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegration`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationByName

> Integration GetIntegrationByName(ctx, name).OrgId(orgId).Execute()

Get integration by name



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationByName(context.Background(), name).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationByName`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**Integration**](Integration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationKubeSettings

> map[string]interface{} GetIntegrationKubeSettings(ctx, id).Execute()

Get Kubernetes settings



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationKubeSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationKubeSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationKubeSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationKubeSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationKubeSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIntegrationRemoteGitRepoFilePresence

> RemoteGitRepoFilePresence GetIntegrationRemoteGitRepoFilePresence(ctx, id).RemoteGitRepoId(remoteGitRepoId).Path(path).Ref(ref).Execute()

Check a remote Git repository file



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
	remoteGitRepoId := "remoteGitRepoId_example" // string | 
	path := "path_example" // string | 
	ref := "ref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationRemoteGitRepoFilePresence(context.Background(), id).RemoteGitRepoId(remoteGitRepoId).Path(path).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationRemoteGitRepoFilePresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationRemoteGitRepoFilePresence`: RemoteGitRepoFilePresence
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationRemoteGitRepoFilePresence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationRemoteGitRepoFilePresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteGitRepoId** | **string** |  | 
 **path** | **string** |  | 
 **ref** | **string** |  | 

### Return type

[**RemoteGitRepoFilePresence**](RemoteGitRepoFilePresence.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKubeMachineTypes

> []map[string]interface{} ListIntegrationKubeMachineTypes(ctx, id).Location(location).Execute()

List Kubernetes machine types



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
	location := "location_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationKubeMachineTypes(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationKubeMachineTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKubeMachineTypes`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationKubeMachineTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKubeMachineTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | **string** |  | 

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


## ListIntegrationKubeRegions

> []map[string]interface{} ListIntegrationKubeRegions(ctx, id).Execute()

List Kubernetes regions



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationKubeRegions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationKubeRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKubeRegions`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationKubeRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKubeRegionsRequest struct via the builder pattern


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


## ListIntegrationKubeVersions

> []KubeVersion ListIntegrationKubeVersions(ctx, id).Location(location).Execute()

List Kubernetes versions



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
	location := "location_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationKubeVersions(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationKubeVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKubeVersions`: []KubeVersion
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationKubeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKubeVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | **string** |  | 

### Return type

[**[]KubeVersion**](KubeVersion.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKubeZones

> []map[string]interface{} ListIntegrationKubeZones(ctx, id).Execute()

List Kubernetes zones



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationKubeZones(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationKubeZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKubeZones`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationKubeZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKubeZonesRequest struct via the builder pattern


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


## ListIntegrationRemoteGitRepoBranches

> []string ListIntegrationRemoteGitRepoBranches(ctx, id).RemoteGitRepoId(remoteGitRepoId).Execute()

List remote Git repository branches



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
	remoteGitRepoId := "remoteGitRepoId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationRemoteGitRepoBranches(context.Background(), id).RemoteGitRepoId(remoteGitRepoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationRemoteGitRepoBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationRemoteGitRepoBranches`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationRemoteGitRepoBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationRemoteGitRepoBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteGitRepoId** | **string** |  | 

### Return type

**[]string**

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationRemoteGitRepoTags

> []string ListIntegrationRemoteGitRepoTags(ctx, id).RemoteGitRepoId(remoteGitRepoId).Execute()

List remote Git repository tags



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
	remoteGitRepoId := "remoteGitRepoId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationRemoteGitRepoTags(context.Background(), id).RemoteGitRepoId(remoteGitRepoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationRemoteGitRepoTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationRemoteGitRepoTags`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationRemoteGitRepoTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationRemoteGitRepoTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteGitRepoId** | **string** |  | 

### Return type

**[]string**

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationRemoteGitRepos

> []RemoteGitRepo ListIntegrationRemoteGitRepos(ctx, id).Execute()

List remote Git repositories



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationRemoteGitRepos(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationRemoteGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationRemoteGitRepos`: []RemoteGitRepo
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationRemoteGitRepos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationRemoteGitReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RemoteGitRepo**](RemoteGitRepo.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationScopes

> []IntegrationScope ListIntegrationScopes(ctx, id).Execute()

List integration scopes



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationScopes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationScopes`: []IntegrationScope
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IntegrationScope**](IntegrationScope.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationStorageBuckets

> []string ListIntegrationStorageBuckets(ctx, id).Execute()

List storage buckets



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationStorageBuckets(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationStorageBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationStorageBuckets`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationStorageBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationStorageBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationStorageClasses

> []string ListIntegrationStorageClasses(ctx, id).Execute()

List storage classes



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrationStorageClasses(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrationStorageClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationStorageClasses`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrationStorageClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationStorageClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrations

> []Integration ListIntegrations(ctx).OrgId(orgId).ProjectIds(projectIds).Labels(labels).Execute()

List integrations



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
	labels := "labels_example" // string | Comma-separated labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrations(context.Background()).OrgId(orgId).ProjectIds(projectIds).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrations`: []Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **projectIds** | **string** | Comma-separated project ids | 
 **labels** | **string** | Comma-separated labels | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveIntegration

> ResolveIntegrationResult ResolveIntegration(ctx).NewIntegrationInput(newIntegrationInput).Execute()

Resolve or create integration



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
	newIntegrationInput := *openapiclient.NewNewIntegrationInput(int32(123), "Name_example", "Title_example", []string{"Kinds_example"}) // NewIntegrationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ResolveIntegration(context.Background()).NewIntegrationInput(newIntegrationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ResolveIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveIntegration`: ResolveIntegrationResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ResolveIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newIntegrationInput** | [**NewIntegrationInput**](NewIntegrationInput.md) |  | 

### Return type

[**ResolveIntegrationResult**](ResolveIntegrationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIntegrationPermissions

> OperationResult TestIntegrationPermissions(ctx, id).Execute()

Test integration permissions



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
	resp, r, err := apiClient.IntegrationsAPI.TestIntegrationPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.TestIntegrationPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestIntegrationPermissions`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.TestIntegrationPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestIntegrationPermissionsRequest struct via the builder pattern


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


## UpdateIntegration

> Integration UpdateIntegration(ctx, id).UpdateIntegrationInput(updateIntegrationInput).Execute()

Update integration



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
	updateIntegrationInput := *openapiclient.NewUpdateIntegrationInput("Title_example", "Name_example", []string{"Kinds_example"}) // UpdateIntegrationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UpdateIntegration(context.Background(), id).UpdateIntegrationInput(updateIntegrationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UpdateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIntegration`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UpdateIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIntegrationInput** | [**UpdateIntegrationInput**](UpdateIntegrationInput.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAppAccessHostname

> ValidationResult ValidateAppAccessHostname(ctx, id).ValidateAppAccessHostnameInput(validateAppAccessHostnameInput).Execute()

Validate an app-access hostname



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
	validateAppAccessHostnameInput := *openapiclient.NewValidateAppAccessHostnameInput("Host_example") // ValidateAppAccessHostnameInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ValidateAppAccessHostname(context.Background(), id).ValidateAppAccessHostnameInput(validateAppAccessHostnameInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ValidateAppAccessHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAppAccessHostname`: ValidationResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ValidateAppAccessHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAppAccessHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validateAppAccessHostnameInput** | [**ValidateAppAccessHostnameInput**](ValidateAppAccessHostnameInput.md) |  | 

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

