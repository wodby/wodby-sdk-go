# \BackupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupsGet**](BackupsAPI.md#BackupsGet) | **Get** /backups | List backups
[**BackupsIdGet**](BackupsAPI.md#BackupsIdGet) | **Get** /backups/{id} | Get backup
[**BackupsPost**](BackupsAPI.md#BackupsPost) | **Post** /backups | Create backup



## BackupsGet

> []Backup BackupsGet(ctx).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).BackupName(backupName).Execute()

List backups

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	appInstanceId := int32(56) // int32 |  (optional)
	appServiceId := int32(56) // int32 |  (optional)
	databaseId := int32(56) // int32 |  (optional)
	databaseDbId := int32(56) // int32 |  (optional)
	backupName := "backupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.BackupsGet(context.Background()).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).BackupName(backupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.BackupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupsGet`: []Backup
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.BackupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **appServiceId** | **int32** |  | 
 **databaseId** | **int32** |  | 
 **databaseDbId** | **int32** |  | 
 **backupName** | **string** |  | 

### Return type

[**[]Backup**](Backup.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupsIdGet

> Backup BackupsIdGet(ctx, id).Execute()

Get backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.BackupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.BackupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupsIdGet`: Backup
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.BackupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Backup**](Backup.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupsPost

> OperationResult BackupsPost(ctx).NewBackupInput(newBackupInput).Execute()

Create backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	newBackupInput := *openapiclient.NewNewBackupInput(int32(123), "Bucket_example") // NewBackupInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.BackupsPost(context.Background()).NewBackupInput(newBackupInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.BackupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupsPost`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.BackupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBackupInput** | [**NewBackupInput**](NewBackupInput.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[accessTokenHeader](../README.md#accessTokenHeader), [apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

