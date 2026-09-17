# \BackupApi

All URIs are relative to *https://api.wodby.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupApi.md#CreateBackup) | **Post** /backups | Create manual backup
[**DeleteBackup**](BackupApi.md#DeleteBackup) | **Delete** /backups/{id} | Delete backup
[**GetBackup**](BackupApi.md#GetBackup) | **Get** /backups/{id} | Retrieve backup
[**GetBackups**](BackupApi.md#GetBackups) | **Get** /backups | Retrieve backups by instance


# **CreateBackup**
> ResponseTask CreateBackup(ctx, data)
Create manual backup

Create a manual instance backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **data** | [**RequestBackupCreate**](RequestBackupCreate.md)|  | 

### Return type

[**ResponseTask**](ResponseTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackup**
> ResponseTask DeleteBackup(ctx, id)
Delete backup

Delete backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Backup ID | 

### Return type

[**ResponseTask**](ResponseTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackup**
> Backup GetBackup(ctx, id)
Retrieve backup

Retrieve backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Backup ID | 

### Return type

[**Backup**](Backup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackups**
> []Backup GetBackups(ctx, optional)
Retrieve backups by instance

Retrieve backups by instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | [**string**](.md)| Organization ID | 
 **instanceId** | [**string**](.md)| Instance ID | 
 **serverId** | [**string**](.md)| Server ID | 
 **type_** | **string**| Backup type | 
 **status** | **string**| Backup status | 
 **days** | **int32**| Get backups for N days | 

### Return type

[**[]Backup**](Backup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

