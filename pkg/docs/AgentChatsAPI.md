# \AgentChatsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerAgentChatApproval**](AgentChatsAPI.md#AnswerAgentChatApproval) | **Post** /agent-chats/{id}/approvals | Approve or decline agent change
[**AnswerAgentChatQuestion**](AgentChatsAPI.md#AnswerAgentChatQuestion) | **Post** /agent-chats/{id}/answers | Answer agent question
[**DeleteAgentChat**](AgentChatsAPI.md#DeleteAgentChat) | **Delete** /agent-chats/{id} | Delete agent chat
[**GetAgentChat**](AgentChatsAPI.md#GetAgentChat) | **Get** /agent-chats/{id} | Get agent chat
[**ListAgentChats**](AgentChatsAPI.md#ListAgentChats) | **Get** /agent-chats | List agent chats
[**SendAgentChatMessage**](AgentChatsAPI.md#SendAgentChatMessage) | **Post** /agent-chats/{id}/messages | Send agent chat message
[**StartAgentChat**](AgentChatsAPI.md#StartAgentChat) | **Post** /agent-chats | Start agent chat
[**StopAgentChat**](AgentChatsAPI.md#StopAgentChat) | **Post** /agent-chats/{id}/stop | Stop agent chat
[**UpdateAgentChatApprovalMode**](AgentChatsAPI.md#UpdateAgentChatApprovalMode) | **Put** /agent-chats/{id}/approval-mode | Set agent chat approval mode



## AnswerAgentChatApproval

> AgentChat AnswerAgentChatApproval(ctx, id).AgentChatApprovalRequest(agentChatApprovalRequest).OrgId(orgId).Execute()

Approve or decline agent change



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
	agentChatApprovalRequest := *openapiclient.NewAgentChatApprovalRequest(int32(123), false) // AgentChatApprovalRequest | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.AnswerAgentChatApproval(context.Background(), id).AgentChatApprovalRequest(agentChatApprovalRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.AnswerAgentChatApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnswerAgentChatApproval`: AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.AnswerAgentChatApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerAgentChatApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentChatApprovalRequest** | [**AgentChatApprovalRequest**](AgentChatApprovalRequest.md) |  | 
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnswerAgentChatQuestion

> AgentChat AnswerAgentChatQuestion(ctx, id).AgentChatAnswerRequest(agentChatAnswerRequest).OrgId(orgId).Execute()

Answer agent question



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
	agentChatAnswerRequest := *openapiclient.NewAgentChatAnswerRequest(int32(123), []string{"Answers_example"}) // AgentChatAnswerRequest | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.AnswerAgentChatQuestion(context.Background(), id).AgentChatAnswerRequest(agentChatAnswerRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.AnswerAgentChatQuestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnswerAgentChatQuestion`: AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.AnswerAgentChatQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerAgentChatQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentChatAnswerRequest** | [**AgentChatAnswerRequest**](AgentChatAnswerRequest.md) |  | 
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentChat

> OperationResult DeleteAgentChat(ctx, id).OrgId(orgId).Execute()

Delete agent chat



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
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.DeleteAgentChat(context.Background(), id).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.DeleteAgentChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgentChat`: OperationResult
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.DeleteAgentChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

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


## GetAgentChat

> AgentChatDetail GetAgentChat(ctx, id).OrgId(orgId).Execute()

Get agent chat



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
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.GetAgentChat(context.Background(), id).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.GetAgentChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentChat`: AgentChatDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.GetAgentChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AgentChatDetail**](AgentChatDetail.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentChats

> []AgentChat ListAgentChats(ctx).OrgId(orgId).Execute()

List agent chats



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
	resp, r, err := apiClient.AgentChatsAPI.ListAgentChats(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.ListAgentChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentChats`: []AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.ListAgentChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAgentChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**[]AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAgentChatMessage

> AgentChat SendAgentChatMessage(ctx, id).AgentChatMessageRequest(agentChatMessageRequest).OrgId(orgId).Execute()

Send agent chat message



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
	agentChatMessageRequest := *openapiclient.NewAgentChatMessageRequest("Message_example") // AgentChatMessageRequest | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.SendAgentChatMessage(context.Background(), id).AgentChatMessageRequest(agentChatMessageRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.SendAgentChatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendAgentChatMessage`: AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.SendAgentChatMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendAgentChatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentChatMessageRequest** | [**AgentChatMessageRequest**](AgentChatMessageRequest.md) |  | 
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartAgentChat

> AgentChat StartAgentChat(ctx).StartAgentChatRequest(startAgentChatRequest).Execute()

Start agent chat



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
	startAgentChatRequest := *openapiclient.NewStartAgentChatRequest("Message_example") // StartAgentChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.StartAgentChat(context.Background()).StartAgentChatRequest(startAgentChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.StartAgentChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartAgentChat`: AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.StartAgentChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartAgentChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAgentChatRequest** | [**StartAgentChatRequest**](StartAgentChatRequest.md) |  | 

### Return type

[**AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopAgentChat

> AgentChat StopAgentChat(ctx, id).OrgId(orgId).Execute()

Stop agent chat



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
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.StopAgentChat(context.Background(), id).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.StopAgentChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopAgentChat`: AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.StopAgentChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAgentChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentChatApprovalMode

> AgentChat UpdateAgentChatApprovalMode(ctx, id).AgentChatApprovalModeRequest(agentChatApprovalModeRequest).OrgId(orgId).Execute()

Set agent chat approval mode



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
	agentChatApprovalModeRequest := *openapiclient.NewAgentChatApprovalModeRequest(openapiclient.AgentChatApprovalMode("ask")) // AgentChatApprovalModeRequest | 
	orgId := int32(56) // int32 | Optional for API-key requests; defaults to the API key's organization. If provided, it must match the key's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentChatsAPI.UpdateAgentChatApprovalMode(context.Background(), id).AgentChatApprovalModeRequest(agentChatApprovalModeRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentChatsAPI.UpdateAgentChatApprovalMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentChatApprovalMode`: AgentChat
	fmt.Fprintf(os.Stdout, "Response from `AgentChatsAPI.UpdateAgentChatApprovalMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentChatApprovalModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentChatApprovalModeRequest** | [**AgentChatApprovalModeRequest**](AgentChatApprovalModeRequest.md) |  | 
 **orgId** | **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. If provided, it must match the key&#39;s organization. | 

### Return type

[**AgentChat**](AgentChat.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

