# \ServerApi

All URIs are relative to *https://api.wodby.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServer**](ServerApi.md#CreateServer) | **Post** /servers | Create server
[**GetServer**](ServerApi.md#GetServer) | **Get** /servers/{id} | Retrieve server
[**GetServerIntegrationOptions**](ServerApi.md#GetServerIntegrationOptions) | **Get** /servers/integrations/{id}/options | Retrieve server integration options
[**GetServerIntegrations**](ServerApi.md#GetServerIntegrations) | **Get** /servers/integrations | Retrieve server integrations
[**GetServers**](ServerApi.md#GetServers) | **Get** /servers | Retrieve servers


# **CreateServer**
> ResponseTask CreateServer(ctx, data)
Create server

Create a server with a hosting integration, such as a connected DigitalOcean team. The integration's provider creates the machine and Wodby installs its infrastructure on it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **data** | [**RequestServerCreate**](RequestServerCreate.md)|  | 

### Return type

[**ResponseTask**](ResponseTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServer**
> Server GetServer(ctx, id)
Retrieve server

Retrieve server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Server ID | 

### Return type

[**Server**](Server.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerIntegrationOptions**
> []ServerIntegrationRegion GetServerIntegrationOptions(ctx, id)
Retrieve server integration options

Retrieve the regions and server sizes a hosting integration offers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Integration ID | 

### Return type

[**[]ServerIntegrationRegion**](ServerIntegrationRegion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerIntegrations**
> []ServerIntegration GetServerIntegrations(ctx, orgId)
Retrieve server integrations

Retrieve an organization's integrations that can create servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orgId** | [**string**](.md)| Organization ID | 

### Return type

[**[]ServerIntegration**](ServerIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServers**
> []Server GetServers(ctx, optional)
Retrieve servers

Retrieve servers

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
 **name** | **string**| Server name | 

### Return type

[**[]Server**](Server.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

