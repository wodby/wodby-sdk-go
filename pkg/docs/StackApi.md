# \StackApi

All URIs are relative to *https://api.wodby.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStack**](StackApi.md#GetStack) | **Get** /stacks/{id} | 
[**GetStacks**](StackApi.md#GetStacks) | **Get** /stacks | 
[**UpdateStackFromUpstream**](StackApi.md#UpdateStackFromUpstream) | **Post** /stacks/{id}/update | 
[**UpdateStacksFromUpstream**](StackApi.md#UpdateStacksFromUpstream) | **Post** /stacks/update | 


# **GetStack**
> Stack GetStack(ctx, id)


Retrieve stack

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Stack ID | 

### Return type

[**Stack**](Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStacks**
> []Stack GetStacks(ctx, optional)


Retrieve stacks

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

### Return type

[**[]Stack**](Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStackFromUpstream**
> ResponseTask UpdateStackFromUpstream(ctx, id)


Update official stack from upstream

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Stack ID | 

### Return type

[**ResponseTask**](ResponseTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStacksFromUpstream**
> ResponseTask UpdateStacksFromUpstream(ctx, data)


Update official stacks from upstream

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **data** | [**RequestStacksUpdate**](RequestStacksUpdate.md)|  | 

### Return type

[**ResponseTask**](ResponseTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

