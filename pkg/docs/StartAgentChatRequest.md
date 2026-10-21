# StartAgentChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **int32** | Optional for API-key requests; defaults to the API key&#39;s organization. | [optional] 
**Message** | **string** |  | 
**ApprovalMode** | Pointer to [**AgentChatApprovalMode**](AgentChatApprovalMode.md) | Optional; without it, the chat asks before every change. | [optional] 

## Methods

### NewStartAgentChatRequest

`func NewStartAgentChatRequest(message string, ) *StartAgentChatRequest`

NewStartAgentChatRequest instantiates a new StartAgentChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartAgentChatRequestWithDefaults

`func NewStartAgentChatRequestWithDefaults() *StartAgentChatRequest`

NewStartAgentChatRequestWithDefaults instantiates a new StartAgentChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *StartAgentChatRequest) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *StartAgentChatRequest) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *StartAgentChatRequest) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *StartAgentChatRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetMessage

`func (o *StartAgentChatRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StartAgentChatRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StartAgentChatRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetApprovalMode

`func (o *StartAgentChatRequest) GetApprovalMode() AgentChatApprovalMode`

GetApprovalMode returns the ApprovalMode field if non-nil, zero value otherwise.

### GetApprovalModeOk

`func (o *StartAgentChatRequest) GetApprovalModeOk() (*AgentChatApprovalMode, bool)`

GetApprovalModeOk returns a tuple with the ApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalMode

`func (o *StartAgentChatRequest) SetApprovalMode(v AgentChatApprovalMode)`

SetApprovalMode sets ApprovalMode field to given value.

### HasApprovalMode

`func (o *StartAgentChatRequest) HasApprovalMode() bool`

HasApprovalMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


