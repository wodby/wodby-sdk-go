# \GitReposAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitRepoUsages**](GitReposAPI.md#GetGitRepoUsages) | **Get** /git-repos/{id}/usages | Get Git repository usages
[**UpdateGitRepoFromGit**](GitReposAPI.md#UpdateGitRepoFromGit) | **Post** /git-repos/{id}/actions/update-from-git | Update all Git repository usages



## GetGitRepoUsages

> GitRepoUsages GetGitRepoUsages(ctx, id).Execute()

Get Git repository usages



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
	resp, r, err := apiClient.GitReposAPI.GetGitRepoUsages(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitReposAPI.GetGitRepoUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitRepoUsages`: GitRepoUsages
	fmt.Fprintf(os.Stdout, "Response from `GitReposAPI.GetGitRepoUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitRepoUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitRepoUsages**](GitRepoUsages.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGitRepoFromGit

> OperationResult UpdateGitRepoFromGit(ctx, id).UpdateStackFromGitRequest(updateStackFromGitRequest).Execute()

Update all Git repository usages



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
	resp, r, err := apiClient.GitReposAPI.UpdateGitRepoFromGit(context.Background(), id).UpdateStackFromGitRequest(updateStackFromGitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitReposAPI.UpdateGitRepoFromGit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGitRepoFromGit`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `GitReposAPI.UpdateGitRepoFromGit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGitRepoFromGitRequest struct via the builder pattern


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

