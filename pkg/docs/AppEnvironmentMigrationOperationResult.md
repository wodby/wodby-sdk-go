# AppEnvironmentMigrationOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**TaskId** | Pointer to **NullableInt32** | Task that runs the requested step. | [optional] 
**Migration** | [**AppEnvironmentMigration**](AppEnvironmentMigration.md) |  | 

## Methods

### NewAppEnvironmentMigrationOperationResult

`func NewAppEnvironmentMigrationOperationResult(success bool, migration AppEnvironmentMigration, ) *AppEnvironmentMigrationOperationResult`

NewAppEnvironmentMigrationOperationResult instantiates a new AppEnvironmentMigrationOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentMigrationOperationResultWithDefaults

`func NewAppEnvironmentMigrationOperationResultWithDefaults() *AppEnvironmentMigrationOperationResult`

NewAppEnvironmentMigrationOperationResultWithDefaults instantiates a new AppEnvironmentMigrationOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AppEnvironmentMigrationOperationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AppEnvironmentMigrationOperationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AppEnvironmentMigrationOperationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTaskId

`func (o *AppEnvironmentMigrationOperationResult) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AppEnvironmentMigrationOperationResult) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AppEnvironmentMigrationOperationResult) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *AppEnvironmentMigrationOperationResult) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *AppEnvironmentMigrationOperationResult) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *AppEnvironmentMigrationOperationResult) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetMigration

`func (o *AppEnvironmentMigrationOperationResult) GetMigration() AppEnvironmentMigration`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *AppEnvironmentMigrationOperationResult) GetMigrationOk() (*AppEnvironmentMigration, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *AppEnvironmentMigrationOperationResult) SetMigration(v AppEnvironmentMigration)`

SetMigration sets Migration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


