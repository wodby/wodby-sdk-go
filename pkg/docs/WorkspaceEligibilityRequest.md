# WorkspaceEligibilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackRevId** | **int32** |  | 
**ClusterId** | Pointer to **NullableInt32** |  | [optional] 
**DisabledServiceIds** | **[]int32** |  | 

## Methods

### NewWorkspaceEligibilityRequest

`func NewWorkspaceEligibilityRequest(stackRevId int32, disabledServiceIds []int32, ) *WorkspaceEligibilityRequest`

NewWorkspaceEligibilityRequest instantiates a new WorkspaceEligibilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceEligibilityRequestWithDefaults

`func NewWorkspaceEligibilityRequestWithDefaults() *WorkspaceEligibilityRequest`

NewWorkspaceEligibilityRequestWithDefaults instantiates a new WorkspaceEligibilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackRevId

`func (o *WorkspaceEligibilityRequest) GetStackRevId() int32`

GetStackRevId returns the StackRevId field if non-nil, zero value otherwise.

### GetStackRevIdOk

`func (o *WorkspaceEligibilityRequest) GetStackRevIdOk() (*int32, bool)`

GetStackRevIdOk returns a tuple with the StackRevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevId

`func (o *WorkspaceEligibilityRequest) SetStackRevId(v int32)`

SetStackRevId sets StackRevId field to given value.


### GetClusterId

`func (o *WorkspaceEligibilityRequest) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *WorkspaceEligibilityRequest) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *WorkspaceEligibilityRequest) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *WorkspaceEligibilityRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *WorkspaceEligibilityRequest) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *WorkspaceEligibilityRequest) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetDisabledServiceIds

`func (o *WorkspaceEligibilityRequest) GetDisabledServiceIds() []int32`

GetDisabledServiceIds returns the DisabledServiceIds field if non-nil, zero value otherwise.

### GetDisabledServiceIdsOk

`func (o *WorkspaceEligibilityRequest) GetDisabledServiceIdsOk() (*[]int32, bool)`

GetDisabledServiceIdsOk returns a tuple with the DisabledServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledServiceIds

`func (o *WorkspaceEligibilityRequest) SetDisabledServiceIds(v []int32)`

SetDisabledServiceIds sets DisabledServiceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


