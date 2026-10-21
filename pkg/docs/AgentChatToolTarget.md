# AgentChatToolTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | What the argument refers to, such as Environment, Project or CI integration. | 
**Name** | **string** |  | 
**Argument** | **string** | The call&#39;s argument this names, such as project_id. | 

## Methods

### NewAgentChatToolTarget

`func NewAgentChatToolTarget(kind string, name string, argument string, ) *AgentChatToolTarget`

NewAgentChatToolTarget instantiates a new AgentChatToolTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentChatToolTargetWithDefaults

`func NewAgentChatToolTargetWithDefaults() *AgentChatToolTarget`

NewAgentChatToolTargetWithDefaults instantiates a new AgentChatToolTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AgentChatToolTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentChatToolTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentChatToolTarget) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *AgentChatToolTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentChatToolTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentChatToolTarget) SetName(v string)`

SetName sets Name field to given value.


### GetArgument

`func (o *AgentChatToolTarget) GetArgument() string`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *AgentChatToolTarget) GetArgumentOk() (*string, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *AgentChatToolTarget) SetArgument(v string)`

SetArgument sets Argument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


