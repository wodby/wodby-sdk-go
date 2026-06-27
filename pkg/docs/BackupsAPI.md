# \BackupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupsAPI.md#CreateBackup) | **Post** /backups | Create backup
[**GetBackup**](BackupsAPI.md#GetBackup) | **Get** /backups/{id} | Get backup
[**ListBackups**](BackupsAPI.md#ListBackups) | **Get** /backups | List backups



## CreateBackup

> OperationResult CreateBackup(ctx).NewBackupInput(newBackupInput).Execute()

Create backup



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
	newBackupInput := *openapiclient.NewNewBackupInput(int32(123), "Bucket_example") // NewBackupInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.CreateBackup(context.Background()).NewBackupInput(newBackupInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackup`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.CreateBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBackupInput** | [**NewBackupInput**](NewBackupInput.md) |  | 

### Return type

[**OperationResult**](OperationResult.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackup

> Backup GetBackup(ctx, id).Execute()

Get backup



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
	resp, r, err := apiClient.BackupsAPI.GetBackup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.GetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackup`: Backup
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.GetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Backup**](Backup.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> []Backup ListBackups(ctx).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).BackupName(backupName).Execute()

List backups



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
	appInstanceId := int32(56) // int32 |  (optional)
	appServiceId := int32(56) // int32 |  (optional)
	databaseId := int32(56) // int32 |  (optional)
	databaseDbId := int32(56) // int32 |  (optional)
	backupName := "backupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackups(context.Background()).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).BackupName(backupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: []Backup
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


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

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

