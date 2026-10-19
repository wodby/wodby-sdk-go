# \ProjectMembershipsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectMemberships**](ProjectMembershipsAPI.md#CreateProjectMemberships) | **Post** /project-memberships | Add project members
[**DeleteProjectMembership**](ProjectMembershipsAPI.md#DeleteProjectMembership) | **Delete** /project-memberships/{id} | Remove project membership
[**ListProjectMemberships**](ProjectMembershipsAPI.md#ListProjectMemberships) | **Get** /project-memberships | List project memberships
[**UpdateProjectMembership**](ProjectMembershipsAPI.md#UpdateProjectMembership) | **Put** /project-memberships/{id} | Update project membership



## CreateProjectMemberships

> []ProjectMembership CreateProjectMemberships(ctx).NewProjectMembershipsInput(newProjectMembershipsInput).Execute()

Add project members



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
	newProjectMembershipsInput := *openapiclient.NewNewProjectMembershipsInput(int32(123), "Role_example") // NewProjectMembershipsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectMembershipsAPI.CreateProjectMemberships(context.Background()).NewProjectMembershipsInput(newProjectMembershipsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectMembershipsAPI.CreateProjectMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectMemberships`: []ProjectMembership
	fmt.Fprintf(os.Stdout, "Response from `ProjectMembershipsAPI.CreateProjectMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newProjectMembershipsInput** | [**NewProjectMembershipsInput**](NewProjectMembershipsInput.md) |  | 

### Return type

[**[]ProjectMembership**](ProjectMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectMembership

> OperationResult DeleteProjectMembership(ctx, id).Execute()

Remove project membership



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
	resp, r, err := apiClient.ProjectMembershipsAPI.DeleteProjectMembership(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectMembershipsAPI.DeleteProjectMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectMembership`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectMembershipsAPI.DeleteProjectMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectMembershipRequest struct via the builder pattern


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


## ListProjectMemberships

> []ProjectMembership ListProjectMemberships(ctx).ProjectId(projectId).Execute()

List project memberships



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
	projectId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectMembershipsAPI.ListProjectMemberships(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectMembershipsAPI.ListProjectMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectMemberships`: []ProjectMembership
	fmt.Fprintf(os.Stdout, "Response from `ProjectMembershipsAPI.ListProjectMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProjectMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 

### Return type

[**[]ProjectMembership**](ProjectMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectMembership

> ProjectMembership UpdateProjectMembership(ctx, id).ProjectMembershipRoleInput(projectMembershipRoleInput).Execute()

Update project membership



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
	projectMembershipRoleInput := *openapiclient.NewProjectMembershipRoleInput("Role_example") // ProjectMembershipRoleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectMembershipsAPI.UpdateProjectMembership(context.Background(), id).ProjectMembershipRoleInput(projectMembershipRoleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectMembershipsAPI.UpdateProjectMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectMembership`: ProjectMembership
	fmt.Fprintf(os.Stdout, "Response from `ProjectMembershipsAPI.UpdateProjectMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectMembershipRoleInput** | [**ProjectMembershipRoleInput**](ProjectMembershipRoleInput.md) |  | 

### Return type

[**ProjectMembership**](ProjectMembership.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

