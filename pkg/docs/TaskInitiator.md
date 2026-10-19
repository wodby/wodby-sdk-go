# TaskInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**UserId** | **int32** |  | 
**ApiKeyId** | Pointer to **NullableInt32** |  | [optional] 
**AgentConnectionId** | Pointer to **NullableInt32** |  | [optional] 
**OauthClientId** | **string** |  | 
**OauthClientName** | **string** |  | 
**ClientName** | **string** |  | 
**ClientVersion** | **string** |  | 
**ToolName** | **string** |  | 
**RequestId** | **string** |  | 

## Methods

### NewTaskInitiator

`func NewTaskInitiator(source string, userId int32, oauthClientId string, oauthClientName string, clientName string, clientVersion string, toolName string, requestId string, ) *TaskInitiator`

NewTaskInitiator instantiates a new TaskInitiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInitiatorWithDefaults

`func NewTaskInitiatorWithDefaults() *TaskInitiator`

NewTaskInitiatorWithDefaults instantiates a new TaskInitiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TaskInitiator) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TaskInitiator) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TaskInitiator) SetSource(v string)`

SetSource sets Source field to given value.


### GetUserId

`func (o *TaskInitiator) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TaskInitiator) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TaskInitiator) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetApiKeyId

`func (o *TaskInitiator) GetApiKeyId() int32`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *TaskInitiator) GetApiKeyIdOk() (*int32, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *TaskInitiator) SetApiKeyId(v int32)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *TaskInitiator) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *TaskInitiator) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *TaskInitiator) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetAgentConnectionId

`func (o *TaskInitiator) GetAgentConnectionId() int32`

GetAgentConnectionId returns the AgentConnectionId field if non-nil, zero value otherwise.

### GetAgentConnectionIdOk

`func (o *TaskInitiator) GetAgentConnectionIdOk() (*int32, bool)`

GetAgentConnectionIdOk returns a tuple with the AgentConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionId

`func (o *TaskInitiator) SetAgentConnectionId(v int32)`

SetAgentConnectionId sets AgentConnectionId field to given value.

### HasAgentConnectionId

`func (o *TaskInitiator) HasAgentConnectionId() bool`

HasAgentConnectionId returns a boolean if a field has been set.

### SetAgentConnectionIdNil

`func (o *TaskInitiator) SetAgentConnectionIdNil(b bool)`

 SetAgentConnectionIdNil sets the value for AgentConnectionId to be an explicit nil

### UnsetAgentConnectionId
`func (o *TaskInitiator) UnsetAgentConnectionId()`

UnsetAgentConnectionId ensures that no value is present for AgentConnectionId, not even an explicit nil
### GetOauthClientId

`func (o *TaskInitiator) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *TaskInitiator) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *TaskInitiator) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### GetOauthClientName

`func (o *TaskInitiator) GetOauthClientName() string`

GetOauthClientName returns the OauthClientName field if non-nil, zero value otherwise.

### GetOauthClientNameOk

`func (o *TaskInitiator) GetOauthClientNameOk() (*string, bool)`

GetOauthClientNameOk returns a tuple with the OauthClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientName

`func (o *TaskInitiator) SetOauthClientName(v string)`

SetOauthClientName sets OauthClientName field to given value.


### GetClientName

`func (o *TaskInitiator) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *TaskInitiator) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *TaskInitiator) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetClientVersion

`func (o *TaskInitiator) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *TaskInitiator) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *TaskInitiator) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.


### GetToolName

`func (o *TaskInitiator) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *TaskInitiator) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *TaskInitiator) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetRequestId

`func (o *TaskInitiator) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TaskInitiator) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TaskInitiator) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


