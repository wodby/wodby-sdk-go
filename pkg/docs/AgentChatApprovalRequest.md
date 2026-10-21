# AgentChatApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **int32** | The tool message that waits for approval. | 
**Approve** | **bool** |  | 

## Methods

### NewAgentChatApprovalRequest

`func NewAgentChatApprovalRequest(messageId int32, approve bool, ) *AgentChatApprovalRequest`

NewAgentChatApprovalRequest instantiates a new AgentChatApprovalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentChatApprovalRequestWithDefaults

`func NewAgentChatApprovalRequestWithDefaults() *AgentChatApprovalRequest`

NewAgentChatApprovalRequestWithDefaults instantiates a new AgentChatApprovalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *AgentChatApprovalRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AgentChatApprovalRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AgentChatApprovalRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetApprove

`func (o *AgentChatApprovalRequest) GetApprove() bool`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *AgentChatApprovalRequest) GetApproveOk() (*bool, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *AgentChatApprovalRequest) SetApprove(v bool)`

SetApprove sets Approve field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


