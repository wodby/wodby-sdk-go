# \OrgSsoAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSSOProvider**](OrgSsoAPI.md#CreateOrgSSOProvider) | **Post** /org-sso-providers | Create organization SSO provider
[**DeleteOrgSSOProvider**](OrgSsoAPI.md#DeleteOrgSSOProvider) | **Delete** /org-sso-providers/{id} | Delete organization SSO provider
[**GetOrgSSOProvider**](OrgSsoAPI.md#GetOrgSSOProvider) | **Get** /org-sso-providers/{id} | Get organization SSO provider
[**ListOrgSSOProviders**](OrgSsoAPI.md#ListOrgSSOProviders) | **Get** /org-sso-providers | List organization SSO providers
[**UpdateOrgSSOProvider**](OrgSsoAPI.md#UpdateOrgSSOProvider) | **Put** /org-sso-providers/{id} | Update organization SSO provider
[**VerifyOrgSSODomain**](OrgSsoAPI.md#VerifyOrgSSODomain) | **Post** /org-sso-providers/{id}/actions/verify-domain | Verify SSO domain



## CreateOrgSSOProvider

> OrgSSOProvider CreateOrgSSOProvider(ctx).OrgSSOProviderInput(orgSSOProviderInput).Execute()

Create organization SSO provider



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
	orgSSOProviderInput := *openapiclient.NewOrgSSOProviderInput("Name_example", false, false, []string{"Domains_example"}) // OrgSSOProviderInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSsoAPI.CreateOrgSSOProvider(context.Background()).OrgSSOProviderInput(orgSSOProviderInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSsoAPI.CreateOrgSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSSOProvider`: OrgSSOProvider
	fmt.Fprintf(os.Stdout, "Response from `OrgSsoAPI.CreateOrgSSOProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSSOProviderInput** | [**OrgSSOProviderInput**](OrgSSOProviderInput.md) |  | 

### Return type

[**OrgSSOProvider**](OrgSSOProvider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSSOProvider

> OperationResult DeleteOrgSSOProvider(ctx, id).Execute()

Delete organization SSO provider



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
	resp, r, err := apiClient.OrgSsoAPI.DeleteOrgSSOProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSsoAPI.DeleteOrgSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgSSOProvider`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `OrgSsoAPI.DeleteOrgSSOProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSSOProviderRequest struct via the builder pattern


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


## GetOrgSSOProvider

> OrgSSOProvider GetOrgSSOProvider(ctx, id).Execute()

Get organization SSO provider



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
	resp, r, err := apiClient.OrgSsoAPI.GetOrgSSOProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSsoAPI.GetOrgSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSSOProvider`: OrgSSOProvider
	fmt.Fprintf(os.Stdout, "Response from `OrgSsoAPI.GetOrgSSOProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgSSOProvider**](OrgSSOProvider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSSOProviders

> []OrgSSOProvider ListOrgSSOProviders(ctx).OrgId(orgId).Execute()

List organization SSO providers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSsoAPI.ListOrgSSOProviders(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSsoAPI.ListOrgSSOProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSSOProviders`: []OrgSSOProvider
	fmt.Fprintf(os.Stdout, "Response from `OrgSsoAPI.ListOrgSSOProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSSOProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**[]OrgSSOProvider**](OrgSSOProvider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSSOProvider

> OrgSSOProvider UpdateOrgSSOProvider(ctx, id).OrgSSOProviderInput(orgSSOProviderInput).Execute()

Update organization SSO provider



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
	orgSSOProviderInput := *openapiclient.NewOrgSSOProviderInput("Name_example", false, false, []string{"Domains_example"}) // OrgSSOProviderInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSsoAPI.UpdateOrgSSOProvider(context.Background(), id).OrgSSOProviderInput(orgSSOProviderInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSsoAPI.UpdateOrgSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSSOProvider`: OrgSSOProvider
	fmt.Fprintf(os.Stdout, "Response from `OrgSsoAPI.UpdateOrgSSOProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgSSOProviderInput** | [**OrgSSOProviderInput**](OrgSSOProviderInput.md) |  | 

### Return type

[**OrgSSOProvider**](OrgSSOProvider.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyOrgSSODomain

> OrgSSODomain VerifyOrgSSODomain(ctx, id).VerifyOrgSSODomainInput(verifyOrgSSODomainInput).Execute()

Verify SSO domain



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
	verifyOrgSSODomainInput := *openapiclient.NewVerifyOrgSSODomainInput("Domain_example") // VerifyOrgSSODomainInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSsoAPI.VerifyOrgSSODomain(context.Background(), id).VerifyOrgSSODomainInput(verifyOrgSSODomainInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSsoAPI.VerifyOrgSSODomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyOrgSSODomain`: OrgSSODomain
	fmt.Fprintf(os.Stdout, "Response from `OrgSsoAPI.VerifyOrgSSODomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyOrgSSODomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyOrgSSODomainInput** | [**VerifyOrgSSODomainInput**](VerifyOrgSSODomainInput.md) |  | 

### Return type

[**OrgSSODomain**](OrgSSODomain.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

