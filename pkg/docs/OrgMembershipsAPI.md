# \OrgMembershipsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgMembership**](OrgMembershipsAPI.md#DeleteOrgMembership) | **Delete** /org-memberships/{id} | Remove organization member
[**GetOrgMembership**](OrgMembershipsAPI.md#GetOrgMembership) | **Get** /org-memberships/{id} | Get org membership
[**InviteOrgMember**](OrgMembershipsAPI.md#InviteOrgMember) | **Post** /org-memberships | Invite organization member
[**ListOrgMemberships**](OrgMembershipsAPI.md#ListOrgMemberships) | **Get** /org-memberships | List org memberships
[**UpdateOrgMembership**](OrgMembershipsAPI.md#UpdateOrgMembership) | **Put** /org-memberships/{id} | Update organization member role



## DeleteOrgMembership

> OperationResult DeleteOrgMembership(ctx, id).Execute()

Remove organization member



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
	resp, r, err := apiClient.OrgMembershipsAPI.DeleteOrgMembership(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgMembershipsAPI.DeleteOrgMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgMembership`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `OrgMembershipsAPI.DeleteOrgMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMembershipRequest struct via the builder pattern


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


## GetOrgMembership

> OrgMembership GetOrgMembership(ctx, id).Execute()

Get org membership



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
	resp, r, err := apiClient.OrgMembershipsAPI.GetOrgMembership(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgMembershipsAPI.GetOrgMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMembership`: OrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgMembershipsAPI.GetOrgMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteOrgMember

> OrgMembership InviteOrgMember(ctx).InviteOrgMemberInput(inviteOrgMemberInput).Execute()

Invite organization member



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
	inviteOrgMemberInput := *openapiclient.NewInviteOrgMemberInput("Email_example", "Role_example") // InviteOrgMemberInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgMembershipsAPI.InviteOrgMember(context.Background()).InviteOrgMemberInput(inviteOrgMemberInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgMembershipsAPI.InviteOrgMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteOrgMember`: OrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgMembershipsAPI.InviteOrgMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteOrgMemberInput** | [**InviteOrgMemberInput**](InviteOrgMemberInput.md) |  | 

### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMemberships

> []OrgMembership ListOrgMemberships(ctx).OrgId(orgId).Execute()

List org memberships



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
	resp, r, err := apiClient.OrgMembershipsAPI.ListOrgMemberships(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgMembershipsAPI.ListOrgMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMemberships`: []OrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgMembershipsAPI.ListOrgMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**[]OrgMembership**](OrgMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMembership

> OrgMembership UpdateOrgMembership(ctx, id).OrgMembershipRoleInput(orgMembershipRoleInput).Execute()

Update organization member role



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
	orgMembershipRoleInput := *openapiclient.NewOrgMembershipRoleInput("Role_example") // OrgMembershipRoleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgMembershipsAPI.UpdateOrgMembership(context.Background(), id).OrgMembershipRoleInput(orgMembershipRoleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgMembershipsAPI.UpdateOrgMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgMembership`: OrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgMembershipsAPI.UpdateOrgMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgMembershipRoleInput** | [**OrgMembershipRoleInput**](OrgMembershipRoleInput.md) |  | 

### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

