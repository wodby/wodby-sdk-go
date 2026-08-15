# AppAccessOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | [**AppAccess**](AppAccess.md) |  | 
**TaskId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppAccessOperationResult

`func NewAppAccessOperationResult(access AppAccess, ) *AppAccessOperationResult`

NewAppAccessOperationResult instantiates a new AppAccessOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessOperationResultWithDefaults

`func NewAppAccessOperationResultWithDefaults() *AppAccessOperationResult`

NewAppAccessOperationResultWithDefaults instantiates a new AppAccessOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AppAccessOperationResult) GetAccess() AppAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AppAccessOperationResult) GetAccessOk() (*AppAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AppAccessOperationResult) SetAccess(v AppAccess)`

SetAccess sets Access field to given value.


### GetTaskId

`func (o *AppAccessOperationResult) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AppAccessOperationResult) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AppAccessOperationResult) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *AppAccessOperationResult) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *AppAccessOperationResult) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *AppAccessOperationResult) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


