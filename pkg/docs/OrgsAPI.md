# \OrgsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrg**](OrgsAPI.md#GetOrg) | **Get** /orgs/{id} | Get org
[**GetOrgSubscription**](OrgsAPI.md#GetOrgSubscription) | **Get** /orgs/{id}/subscription | Get org subscription
[**ListOrgs**](OrgsAPI.md#ListOrgs) | **Get** /orgs | List orgs
[**PreflightAppServiceCapacity**](OrgsAPI.md#PreflightAppServiceCapacity) | **Post** /orgs/{id}/actions/preflight-app-service-capacity | Preflight app-service capacity
[**UpdateOrg**](OrgsAPI.md#UpdateOrg) | **Put** /orgs/{id} | Update org



## GetOrg

> Org GetOrg(ctx, id).Execute()

Get org



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
	resp, r, err := apiClient.OrgsAPI.GetOrg(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrg`: Org
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Org**](Org.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSubscription

> OrgSubscriptionDetails GetOrgSubscription(ctx, id).Execute()

Get org subscription



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
	resp, r, err := apiClient.OrgsAPI.GetOrgSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSubscription`: OrgSubscriptionDetails
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgSubscriptionDetails**](OrgSubscriptionDetails.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgs

> []Org ListOrgs(ctx).Execute()

List orgs



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.ListOrgs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.ListOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgs`: []Org
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.ListOrgs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsRequest struct via the builder pattern


### Return type

[**[]Org**](Org.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreflightAppServiceCapacity

> AppServiceCapacityPreflight PreflightAppServiceCapacity(ctx, id).AppServiceCapacityPreflightRequest(appServiceCapacityPreflightRequest).Execute()

Preflight app-service capacity



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
	appServiceCapacityPreflightRequest := *openapiclient.NewAppServiceCapacityPreflightRequest(int32(123)) // AppServiceCapacityPreflightRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.PreflightAppServiceCapacity(context.Background(), id).AppServiceCapacityPreflightRequest(appServiceCapacityPreflightRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.PreflightAppServiceCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreflightAppServiceCapacity`: AppServiceCapacityPreflight
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.PreflightAppServiceCapacity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreflightAppServiceCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appServiceCapacityPreflightRequest** | [**AppServiceCapacityPreflightRequest**](AppServiceCapacityPreflightRequest.md) |  | 

### Return type

[**AppServiceCapacityPreflight**](AppServiceCapacityPreflight.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> Org UpdateOrg(ctx, id).UpdateOrgRequest(updateOrgRequest).Execute()

Update org



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
	updateOrgRequest := *openapiclient.NewUpdateOrgRequest("Title_example") // UpdateOrgRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.UpdateOrg(context.Background(), id).UpdateOrgRequest(updateOrgRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.UpdateOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrg`: Org
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrgRequest** | [**UpdateOrgRequest**](UpdateOrgRequest.md) |  | 

### Return type

[**Org**](Org.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

