# \OrgSsoAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgSSOProvider**](OrgSsoAPI.md#GetOrgSSOProvider) | **Get** /org-sso-providers/{id} | Get organization SSO provider
[**ListOrgSSOProviders**](OrgSsoAPI.md#ListOrgSSOProviders) | **Get** /org-sso-providers | List organization SSO providers



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

