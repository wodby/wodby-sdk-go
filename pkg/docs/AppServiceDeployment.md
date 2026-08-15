# AppServiceDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**JobName** | **string** |  | 
**Status** | **string** |  | 
**AppServiceId** | **int32** |  | 
**AppServiceBuildId** | Pointer to **NullableInt32** |  | [optional] 
**PreviousAppServiceBuildId** | Pointer to **NullableInt32** |  | [optional] 
**BuildSelectionKind** | **string** |  | 
**SkipPostDeployment** | **bool** |  | 
**Force** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppServiceDeployment

`func NewAppServiceDeployment(id int32, jobName string, status string, appServiceId int32, buildSelectionKind string, skipPostDeployment bool, force bool, createdAt time.Time, updatedAt time.Time, ) *AppServiceDeployment`

NewAppServiceDeployment instantiates a new AppServiceDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceDeploymentWithDefaults

`func NewAppServiceDeploymentWithDefaults() *AppServiceDeployment`

NewAppServiceDeploymentWithDefaults instantiates a new AppServiceDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceDeployment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceDeployment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceDeployment) SetId(v int32)`

SetId sets Id field to given value.


### GetJobName

`func (o *AppServiceDeployment) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *AppServiceDeployment) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *AppServiceDeployment) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetStatus

`func (o *AppServiceDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppServiceDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppServiceDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAppServiceId

`func (o *AppServiceDeployment) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceDeployment) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceDeployment) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetAppServiceBuildId

`func (o *AppServiceDeployment) GetAppServiceBuildId() int32`

GetAppServiceBuildId returns the AppServiceBuildId field if non-nil, zero value otherwise.

### GetAppServiceBuildIdOk

`func (o *AppServiceDeployment) GetAppServiceBuildIdOk() (*int32, bool)`

GetAppServiceBuildIdOk returns a tuple with the AppServiceBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceBuildId

`func (o *AppServiceDeployment) SetAppServiceBuildId(v int32)`

SetAppServiceBuildId sets AppServiceBuildId field to given value.

### HasAppServiceBuildId

`func (o *AppServiceDeployment) HasAppServiceBuildId() bool`

HasAppServiceBuildId returns a boolean if a field has been set.

### SetAppServiceBuildIdNil

`func (o *AppServiceDeployment) SetAppServiceBuildIdNil(b bool)`

 SetAppServiceBuildIdNil sets the value for AppServiceBuildId to be an explicit nil

### UnsetAppServiceBuildId
`func (o *AppServiceDeployment) UnsetAppServiceBuildId()`

UnsetAppServiceBuildId ensures that no value is present for AppServiceBuildId, not even an explicit nil
### GetPreviousAppServiceBuildId

`func (o *AppServiceDeployment) GetPreviousAppServiceBuildId() int32`

GetPreviousAppServiceBuildId returns the PreviousAppServiceBuildId field if non-nil, zero value otherwise.

### GetPreviousAppServiceBuildIdOk

`func (o *AppServiceDeployment) GetPreviousAppServiceBuildIdOk() (*int32, bool)`

GetPreviousAppServiceBuildIdOk returns a tuple with the PreviousAppServiceBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAppServiceBuildId

`func (o *AppServiceDeployment) SetPreviousAppServiceBuildId(v int32)`

SetPreviousAppServiceBuildId sets PreviousAppServiceBuildId field to given value.

### HasPreviousAppServiceBuildId

`func (o *AppServiceDeployment) HasPreviousAppServiceBuildId() bool`

HasPreviousAppServiceBuildId returns a boolean if a field has been set.

### SetPreviousAppServiceBuildIdNil

`func (o *AppServiceDeployment) SetPreviousAppServiceBuildIdNil(b bool)`

 SetPreviousAppServiceBuildIdNil sets the value for PreviousAppServiceBuildId to be an explicit nil

### UnsetPreviousAppServiceBuildId
`func (o *AppServiceDeployment) UnsetPreviousAppServiceBuildId()`

UnsetPreviousAppServiceBuildId ensures that no value is present for PreviousAppServiceBuildId, not even an explicit nil
### GetBuildSelectionKind

`func (o *AppServiceDeployment) GetBuildSelectionKind() string`

GetBuildSelectionKind returns the BuildSelectionKind field if non-nil, zero value otherwise.

### GetBuildSelectionKindOk

`func (o *AppServiceDeployment) GetBuildSelectionKindOk() (*string, bool)`

GetBuildSelectionKindOk returns a tuple with the BuildSelectionKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSelectionKind

`func (o *AppServiceDeployment) SetBuildSelectionKind(v string)`

SetBuildSelectionKind sets BuildSelectionKind field to given value.


### GetSkipPostDeployment

`func (o *AppServiceDeployment) GetSkipPostDeployment() bool`

GetSkipPostDeployment returns the SkipPostDeployment field if non-nil, zero value otherwise.

### GetSkipPostDeploymentOk

`func (o *AppServiceDeployment) GetSkipPostDeploymentOk() (*bool, bool)`

GetSkipPostDeploymentOk returns a tuple with the SkipPostDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPostDeployment

`func (o *AppServiceDeployment) SetSkipPostDeployment(v bool)`

SetSkipPostDeployment sets SkipPostDeployment field to given value.


### GetForce

`func (o *AppServiceDeployment) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AppServiceDeployment) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AppServiceDeployment) SetForce(v bool)`

SetForce sets Force field to given value.


### GetCreatedAt

`func (o *AppServiceDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppServiceDeployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppServiceDeployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppServiceDeployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartedAt

`func (o *AppServiceDeployment) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AppServiceDeployment) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AppServiceDeployment) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AppServiceDeployment) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *AppServiceDeployment) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *AppServiceDeployment) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *AppServiceDeployment) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AppServiceDeployment) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AppServiceDeployment) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AppServiceDeployment) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *AppServiceDeployment) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *AppServiceDeployment) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


