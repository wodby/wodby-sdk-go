# AgentChatDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**AgentChat**](AgentChat.md) |  | 
**Messages** | [**[]AgentChatMessage**](AgentChatMessage.md) |  | 

## Methods

### NewAgentChatDetail

`func NewAgentChatDetail(chat AgentChat, messages []AgentChatMessage, ) *AgentChatDetail`

NewAgentChatDetail instantiates a new AgentChatDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentChatDetailWithDefaults

`func NewAgentChatDetailWithDefaults() *AgentChatDetail`

NewAgentChatDetailWithDefaults instantiates a new AgentChatDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *AgentChatDetail) GetChat() AgentChat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *AgentChatDetail) GetChatOk() (*AgentChat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *AgentChatDetail) SetChat(v AgentChat)`

SetChat sets Chat field to given value.


### GetMessages

`func (o *AgentChatDetail) GetMessages() []AgentChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AgentChatDetail) GetMessagesOk() (*[]AgentChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AgentChatDetail) SetMessages(v []AgentChatMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


