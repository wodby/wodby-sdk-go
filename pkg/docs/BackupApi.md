# \BackupApi

All URIs are relative to *https://api.wodby.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackup**](BackupApi.md#GetBackup) | **Get** /backups/{id} | 
[**GetBackups**](BackupApi.md#GetBackups) | **Get** /backups | 


# **GetBackup**
> Backup GetBackup(ctx, id)


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

