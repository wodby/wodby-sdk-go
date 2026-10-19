# AppEnvironmentMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppEnvironmentId** | **int32** |  | 
**SourceClusterId** | **int32** |  | 
**TargetClusterId** | **int32** |  | 
**TaskId** | **int32** | Task that runs the migration. | 
**DeploymentId** | Pointer to **NullableInt32** | Deployment that brings the environment up on the target cluster. | [optional] 
**Status** | **string** |  | 
**Phase** | **string** |  | 
**Components** | [**[]AppEnvironmentMigrationComponent**](AppEnvironmentMigrationComponent.md) | Services whose data the migration copies with a backup on the source and an import on the target. | 
**FailureMessage** | **string** |  | 
**RetainUntil** | Pointer to **NullableTime** | When the retained source is removed automatically. | [optional] 
**CutoverAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAppEnvironmentMigration

`func NewAppEnvironmentMigration(id int32, appEnvironmentId int32, sourceClusterId int32, targetClusterId int32, taskId int32, status string, phase string, components []AppEnvironmentMigrationComponent, failureMessage string, createdAt time.Time, updatedAt time.Time, ) *AppEnvironmentMigration`

NewAppEnvironmentMigration instantiates a new AppEnvironmentMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentMigrationWithDefaults

`func NewAppEnvironmentMigrationWithDefaults() *AppEnvironmentMigration`

NewAppEnvironmentMigrationWithDefaults instantiates a new AppEnvironmentMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppEnvironmentMigration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEnvironmentMigration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEnvironmentMigration) SetId(v int32)`

SetId sets Id field to given value.


### GetAppEnvironmentId

`func (o *AppEnvironmentMigration) GetAppEnvironmentId() int32`

GetAppEnvironmentId returns the AppEnvironmentId field if non-nil, zero value otherwise.

### GetAppEnvironmentIdOk

`func (o *AppEnvironmentMigration) GetAppEnvironmentIdOk() (*int32, bool)`

GetAppEnvironmentIdOk returns a tuple with the AppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnvironmentId

`func (o *AppEnvironmentMigration) SetAppEnvironmentId(v int32)`

SetAppEnvironmentId sets AppEnvironmentId field to given value.


### GetSourceClusterId

`func (o *AppEnvironmentMigration) GetSourceClusterId() int32`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *AppEnvironmentMigration) GetSourceClusterIdOk() (*int32, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *AppEnvironmentMigration) SetSourceClusterId(v int32)`

SetSourceClusterId sets SourceClusterId field to given value.


### GetTargetClusterId

`func (o *AppEnvironmentMigration) GetTargetClusterId() int32`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *AppEnvironmentMigration) GetTargetClusterIdOk() (*int32, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *AppEnvironmentMigration) SetTargetClusterId(v int32)`

SetTargetClusterId sets TargetClusterId field to given value.


### GetTaskId

`func (o *AppEnvironmentMigration) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AppEnvironmentMigration) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AppEnvironmentMigration) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.


### GetDeploymentId

`func (o *AppEnvironmentMigration) GetDeploymentId() int32`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AppEnvironmentMigration) GetDeploymentIdOk() (*int32, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AppEnvironmentMigration) SetDeploymentId(v int32)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *AppEnvironmentMigration) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *AppEnvironmentMigration) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *AppEnvironmentMigration) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetStatus

`func (o *AppEnvironmentMigration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppEnvironmentMigration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppEnvironmentMigration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPhase

`func (o *AppEnvironmentMigration) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *AppEnvironmentMigration) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *AppEnvironmentMigration) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetComponents

`func (o *AppEnvironmentMigration) GetComponents() []AppEnvironmentMigrationComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *AppEnvironmentMigration) GetComponentsOk() (*[]AppEnvironmentMigrationComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *AppEnvironmentMigration) SetComponents(v []AppEnvironmentMigrationComponent)`

SetComponents sets Components field to given value.


### GetFailureMessage

`func (o *AppEnvironmentMigration) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *AppEnvironmentMigration) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *AppEnvironmentMigration) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.


### GetRetainUntil

`func (o *AppEnvironmentMigration) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *AppEnvironmentMigration) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *AppEnvironmentMigration) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.

### HasRetainUntil

`func (o *AppEnvironmentMigration) HasRetainUntil() bool`

HasRetainUntil returns a boolean if a field has been set.

### SetRetainUntilNil

`func (o *AppEnvironmentMigration) SetRetainUntilNil(b bool)`

 SetRetainUntilNil sets the value for RetainUntil to be an explicit nil

### UnsetRetainUntil
`func (o *AppEnvironmentMigration) UnsetRetainUntil()`

UnsetRetainUntil ensures that no value is present for RetainUntil, not even an explicit nil
### GetCutoverAt

`func (o *AppEnvironmentMigration) GetCutoverAt() time.Time`

GetCutoverAt returns the CutoverAt field if non-nil, zero value otherwise.

### GetCutoverAtOk

`func (o *AppEnvironmentMigration) GetCutoverAtOk() (*time.Time, bool)`

GetCutoverAtOk returns a tuple with the CutoverAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoverAt

`func (o *AppEnvironmentMigration) SetCutoverAt(v time.Time)`

SetCutoverAt sets CutoverAt field to given value.

### HasCutoverAt

`func (o *AppEnvironmentMigration) HasCutoverAt() bool`

HasCutoverAt returns a boolean if a field has been set.

### SetCutoverAtNil

`func (o *AppEnvironmentMigration) SetCutoverAtNil(b bool)`

 SetCutoverAtNil sets the value for CutoverAt to be an explicit nil

### UnsetCutoverAt
`func (o *AppEnvironmentMigration) UnsetCutoverAt()`

UnsetCutoverAt ensures that no value is present for CutoverAt, not even an explicit nil
### GetEndedAt

`func (o *AppEnvironmentMigration) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AppEnvironmentMigration) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AppEnvironmentMigration) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AppEnvironmentMigration) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *AppEnvironmentMigration) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *AppEnvironmentMigration) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetCreatedAt

`func (o *AppEnvironmentMigration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppEnvironmentMigration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppEnvironmentMigration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppEnvironmentMigration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppEnvironmentMigration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppEnvironmentMigration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


