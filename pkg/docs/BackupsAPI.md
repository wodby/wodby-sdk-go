# \BackupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupsAPI.md#CreateBackup) | **Post** /backups | Create backup
[**CreateBackupPreset**](BackupsAPI.md#CreateBackupPreset) | **Post** /backup-presets | Create backup preset
[**DeleteBackupPreset**](BackupsAPI.md#DeleteBackupPreset) | **Delete** /backup-presets/{id} | Delete backup preset
[**GetBackup**](BackupsAPI.md#GetBackup) | **Get** /backups/{id} | Get backup
[**GetBackupPreset**](BackupsAPI.md#GetBackupPreset) | **Get** /backup-presets/{id} | Get backup preset
[**ListBackupPresetBackups**](BackupsAPI.md#ListBackupPresetBackups) | **Get** /backup-presets/{id}/backups | List backup preset backups
[**ListBackupPresets**](BackupsAPI.md#ListBackupPresets) | **Get** /backup-presets | List backup presets
[**ListBackups**](BackupsAPI.md#ListBackups) | **Get** /backups | List backups
[**UpdateBackupPreset**](BackupsAPI.md#UpdateBackupPreset) | **Put** /backup-presets/{id} | Update backup preset



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
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupPreset

> BackupPreset CreateBackupPreset(ctx).NewBackupPresetInput(newBackupPresetInput).Execute()

Create backup preset



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
	newBackupPresetInput := *openapiclient.NewNewBackupPresetInput(int32(123), "Bucket_example", false, false) // NewBackupPresetInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.CreateBackupPreset(context.Background()).NewBackupPresetInput(newBackupPresetInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.CreateBackupPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackupPreset`: BackupPreset
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.CreateBackupPreset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBackupPresetInput** | [**NewBackupPresetInput**](NewBackupPresetInput.md) |  | 

### Return type

[**BackupPreset**](BackupPreset.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupPreset

> OperationResult DeleteBackupPreset(ctx, id).Execute()

Delete backup preset



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
	resp, r, err := apiClient.BackupsAPI.DeleteBackupPreset(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.DeleteBackupPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupPreset`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.DeleteBackupPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupPresetRequest struct via the builder pattern


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
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupPreset

> BackupPreset GetBackupPreset(ctx, id).Execute()

Get backup preset



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
	resp, r, err := apiClient.BackupsAPI.GetBackupPreset(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.GetBackupPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupPreset`: BackupPreset
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.GetBackupPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupPreset**](BackupPreset.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupPresetBackups

> BackupsResponse ListBackupPresetBackups(ctx, id).Page(page).PageSize(pageSize).Execute()

List backup preset backups



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
	page := int32(56) // int32 | Page number, defaults to 1 (optional)
	pageSize := int32(56) // int32 | Page size, defaults to 30 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackupPresetBackups(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackupPresetBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupPresetBackups`: BackupsResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackupPresetBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupPresetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number, defaults to 1 | 
 **pageSize** | **int32** | Page size, defaults to 30 | 

### Return type

[**BackupsResponse**](BackupsResponse.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupPresets

> []BackupPreset ListBackupPresets(ctx).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).OrgId(orgId).BackupName(backupName).ApplicableEnvId(applicableEnvId).ApplicableBackupCategory(applicableBackupCategory).Execute()

List backup presets



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
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)
	backupName := "backupName_example" // string |  (optional)
	applicableEnvId := int32(56) // int32 | Return only presets that apply to this environment. (optional)
	applicableBackupCategory := "applicableBackupCategory_example" // string | Return only presets that apply to this backup category. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackupPresets(context.Background()).AppInstanceId(appInstanceId).AppServiceId(appServiceId).DatabaseId(databaseId).DatabaseDbId(databaseDbId).OrgId(orgId).BackupName(backupName).ApplicableEnvId(applicableEnvId).ApplicableBackupCategory(applicableBackupCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackupPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupPresets`: []BackupPreset
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackupPresets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **int32** |  | 
 **appServiceId** | **int32** |  | 
 **databaseId** | **int32** |  | 
 **databaseDbId** | **int32** |  | 
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 
 **backupName** | **string** |  | 
 **applicableEnvId** | **int32** | Return only presets that apply to this environment. | 
 **applicableBackupCategory** | **string** | Return only presets that apply to this backup category. | 

### Return type

[**[]BackupPreset**](BackupPreset.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupPreset

> BackupPreset UpdateBackupPreset(ctx, id).UpdateBackupPresetInput(updateBackupPresetInput).Execute()

Update backup preset



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
	updateBackupPresetInput := *openapiclient.NewUpdateBackupPresetInput(int32(123), "Bucket_example", false, false, false) // UpdateBackupPresetInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.UpdateBackupPreset(context.Background(), id).UpdateBackupPresetInput(updateBackupPresetInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.UpdateBackupPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBackupPreset`: BackupPreset
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.UpdateBackupPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBackupPresetInput** | [**UpdateBackupPresetInput**](UpdateBackupPresetInput.md) |  | 

### Return type

[**BackupPreset**](BackupPreset.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

