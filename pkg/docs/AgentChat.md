# AgentChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**OrgId** | **int32** |  | 
**Title** | **string** |  | 
**Status** | **string** | idle takes a message; running is the agent working; waiting is a call that waits for approval or an answer. | 
**Error** | **string** | Why the last run stopped without an answer, or empty. | 
**ApprovalMode** | [**AgentChatApprovalMode**](AgentChatApprovalMode.md) |  | 
**ContextTokens** | **int64** | Tokens in the chat&#39;s context after the newest reply. | 
**ContextLimit** | **int64** | Context size at which the agent summarizes the chat to continue it. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAgentChat

`func NewAgentChat(id int32, orgId int32, title string, status string, error_ string, approvalMode AgentChatApprovalMode, contextTokens int64, contextLimit int64, createdAt time.Time, updatedAt time.Time, ) *AgentChat`

NewAgentChat instantiates a new AgentChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentChatWithDefaults

`func NewAgentChatWithDefaults() *AgentChat`

NewAgentChatWithDefaults instantiates a new AgentChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentChat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentChat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentChat) SetId(v int32)`

SetId sets Id field to given value.


### GetOrgId

`func (o *AgentChat) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AgentChat) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AgentChat) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetTitle

`func (o *AgentChat) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentChat) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentChat) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *AgentChat) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentChat) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentChat) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *AgentChat) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentChat) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentChat) SetError(v string)`

SetError sets Error field to given value.


### GetApprovalMode

`func (o *AgentChat) GetApprovalMode() AgentChatApprovalMode`

GetApprovalMode returns the ApprovalMode field if non-nil, zero value otherwise.

### GetApprovalModeOk

`func (o *AgentChat) GetApprovalModeOk() (*AgentChatApprovalMode, bool)`

GetApprovalModeOk returns a tuple with the ApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalMode

`func (o *AgentChat) SetApprovalMode(v AgentChatApprovalMode)`

SetApprovalMode sets ApprovalMode field to given value.


### GetContextTokens

`func (o *AgentChat) GetContextTokens() int64`

GetContextTokens returns the ContextTokens field if non-nil, zero value otherwise.

### GetContextTokensOk

`func (o *AgentChat) GetContextTokensOk() (*int64, bool)`

GetContextTokensOk returns a tuple with the ContextTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextTokens

`func (o *AgentChat) SetContextTokens(v int64)`

SetContextTokens sets ContextTokens field to given value.


### GetContextLimit

`func (o *AgentChat) GetContextLimit() int64`

GetContextLimit returns the ContextLimit field if non-nil, zero value otherwise.

### GetContextLimitOk

`func (o *AgentChat) GetContextLimitOk() (*int64, bool)`

GetContextLimitOk returns a tuple with the ContextLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextLimit

`func (o *AgentChat) SetContextLimit(v int64)`

SetContextLimit sets ContextLimit field to given value.


### GetCreatedAt

`func (o *AgentChat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentChat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentChat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgentChat) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentChat) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentChat) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


