# \AppAuthsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppAuth**](AppAuthsAPI.md#CreateAppAuth) | **Post** /app-auths | Create app auth
[**DeleteAppAuth**](AppAuthsAPI.md#DeleteAppAuth) | **Delete** /app-auths/{id} | Delete app auth
[**ListAppAuths**](AppAuthsAPI.md#ListAppAuths) | **Get** /app-auths | List app auths
[**UpdateAppAuth**](AppAuthsAPI.md#UpdateAppAuth) | **Put** /app-auths/{id} | Update app auth



## CreateAppAuth

> AppAuth CreateAppAuth(ctx).NewAppAuthInput(newAppAuthInput).Execute()

Create app auth



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
	newAppAuthInput := *openapiclient.NewNewAppAuthInput(int32(123), "Login_example", "Password_example", "Realm_example") // NewAppAuthInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAuthsAPI.CreateAppAuth(context.Background()).NewAppAuthInput(newAppAuthInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAuthsAPI.CreateAppAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppAuth`: AppAuth
	fmt.Fprintf(os.Stdout, "Response from `AppAuthsAPI.CreateAppAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAppAuthInput** | [**NewAppAuthInput**](NewAppAuthInput.md) |  | 

### Return type

[**AppAuth**](AppAuth.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppAuth

> OperationResult DeleteAppAuth(ctx, id).Execute()

Delete app auth



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
	resp, r, err := apiClient.AppAuthsAPI.DeleteAppAuth(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAuthsAPI.DeleteAppAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppAuth`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AppAuthsAPI.DeleteAppAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppAuthRequest struct via the builder pattern


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


## ListAppAuths

> []AppAuth ListAppAuths(ctx).AppInstanceId(appInstanceId).Execute()

List app auths



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
	resp, r, err := apiClient.AppAuthsAPI.ListAppAuths(context.Background()).AppInstanceId(appInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAuthsAPI.ListAppAuths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppAuths`: []AppAuth
	fmt.Fprintf(os.Stdout, "Response from `AppAuthsAPI.ListAppAuths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppAuthsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 

### Return type

[**[]AppAuth**](AppAuth.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppAuth

> AppAuth UpdateAppAuth(ctx, id).UpdateAppAuthInput(updateAppAuthInput).Execute()

Update app auth



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
	updateAppAuthInput := *openapiclient.NewUpdateAppAuthInput("Login_example", "Realm_example") // UpdateAppAuthInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAuthsAPI.UpdateAppAuth(context.Background(), id).UpdateAppAuthInput(updateAppAuthInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAuthsAPI.UpdateAppAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppAuth`: AppAuth
	fmt.Fprintf(os.Stdout, "Response from `AppAuthsAPI.UpdateAppAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppAuthInput** | [**UpdateAppAuthInput**](UpdateAppAuthInput.md) |  | 

### Return type

[**AppAuth**](AppAuth.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

