# AgentChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Role** | **string** | A tool message is a call&#39;s result, or a call that waits; a summary replaced earlier messages for the model. | 
**Content** | **string** |  | 
**Tool** | Pointer to **string** | For a tool message, the Wodby tool the call is for. ask_user is a question from the agent. | [optional] 
**ToolArguments** | Pointer to **map[string]interface{}** | For a tool message, the call&#39;s arguments. For ask_user, the question, its options (label, description, recommended) and whether several may be chosen. The recommended options come first and their labels end with (Recommended). | [optional] 
**Approval** | **NullableString** | Set on a call that changes something or asks a question. | 
**ToolDestructive** | **bool** | True for a call that deletes, overwrites, upgrades or interrupts something, or runs a command in a container. | 
**ToolTargets** | [**[]AgentChatToolTarget**](AgentChatToolTarget.md) | For a call that waits for approval, what its ID arguments refer to, as far as the user may see them, each with the argument it names. | 
**IsError** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewAgentChatMessage

`func NewAgentChatMessage(id int32, role string, content string, approval NullableString, toolDestructive bool, toolTargets []AgentChatToolTarget, isError bool, createdAt time.Time, ) *AgentChatMessage`

NewAgentChatMessage instantiates a new AgentChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentChatMessageWithDefaults

`func NewAgentChatMessageWithDefaults() *AgentChatMessage`

NewAgentChatMessageWithDefaults instantiates a new AgentChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentChatMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentChatMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentChatMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetRole

`func (o *AgentChatMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AgentChatMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AgentChatMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *AgentChatMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AgentChatMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AgentChatMessage) SetContent(v string)`

SetContent sets Content field to given value.


### GetTool

`func (o *AgentChatMessage) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *AgentChatMessage) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *AgentChatMessage) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *AgentChatMessage) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetToolArguments

`func (o *AgentChatMessage) GetToolArguments() map[string]interface{}`

GetToolArguments returns the ToolArguments field if non-nil, zero value otherwise.

### GetToolArgumentsOk

`func (o *AgentChatMessage) GetToolArgumentsOk() (*map[string]interface{}, bool)`

GetToolArgumentsOk returns a tuple with the ToolArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolArguments

`func (o *AgentChatMessage) SetToolArguments(v map[string]interface{})`

SetToolArguments sets ToolArguments field to given value.

### HasToolArguments

`func (o *AgentChatMessage) HasToolArguments() bool`

HasToolArguments returns a boolean if a field has been set.

### GetApproval

`func (o *AgentChatMessage) GetApproval() string`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *AgentChatMessage) GetApprovalOk() (*string, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *AgentChatMessage) SetApproval(v string)`

SetApproval sets Approval field to given value.


### SetApprovalNil

`func (o *AgentChatMessage) SetApprovalNil(b bool)`

 SetApprovalNil sets the value for Approval to be an explicit nil

### UnsetApproval
`func (o *AgentChatMessage) UnsetApproval()`

UnsetApproval ensures that no value is present for Approval, not even an explicit nil
### GetToolDestructive

`func (o *AgentChatMessage) GetToolDestructive() bool`

GetToolDestructive returns the ToolDestructive field if non-nil, zero value otherwise.

### GetToolDestructiveOk

`func (o *AgentChatMessage) GetToolDestructiveOk() (*bool, bool)`

GetToolDestructiveOk returns a tuple with the ToolDestructive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolDestructive

`func (o *AgentChatMessage) SetToolDestructive(v bool)`

SetToolDestructive sets ToolDestructive field to given value.


### GetToolTargets

`func (o *AgentChatMessage) GetToolTargets() []AgentChatToolTarget`

GetToolTargets returns the ToolTargets field if non-nil, zero value otherwise.

### GetToolTargetsOk

`func (o *AgentChatMessage) GetToolTargetsOk() (*[]AgentChatToolTarget, bool)`

GetToolTargetsOk returns a tuple with the ToolTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTargets

`func (o *AgentChatMessage) SetToolTargets(v []AgentChatToolTarget)`

SetToolTargets sets ToolTargets field to given value.


### GetIsError

`func (o *AgentChatMessage) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *AgentChatMessage) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *AgentChatMessage) SetIsError(v bool)`

SetIsError sets IsError field to given value.


### GetCreatedAt

`func (o *AgentChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


