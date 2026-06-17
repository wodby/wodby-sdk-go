# TaskJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Status** | **string** |  | 
**After** | Pointer to **[]string** |  | [optional] 
**Timeout** | **int32** |  | 
**IsSystem** | **bool** |  | 
**StatusTitle** | **string** |  | 
**Steps** | [**[]TaskStep**](TaskStep.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTaskJob

`func NewTaskJob(id int32, name string, title string, status string, timeout int32, isSystem bool, statusTitle string, steps []TaskStep, createdAt time.Time, updatedAt time.Time, ) *TaskJob`

NewTaskJob instantiates a new TaskJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskJobWithDefaults

`func NewTaskJobWithDefaults() *TaskJob`

NewTaskJobWithDefaults instantiates a new TaskJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskJob) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskJob) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TaskJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskJob) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *TaskJob) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskJob) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskJob) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *TaskJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAfter

`func (o *TaskJob) GetAfter() []string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *TaskJob) GetAfterOk() (*[]string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *TaskJob) SetAfter(v []string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *TaskJob) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetTimeout

`func (o *TaskJob) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TaskJob) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TaskJob) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetIsSystem

`func (o *TaskJob) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *TaskJob) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *TaskJob) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetStatusTitle

`func (o *TaskJob) GetStatusTitle() string`

GetStatusTitle returns the StatusTitle field if non-nil, zero value otherwise.

### GetStatusTitleOk

`func (o *TaskJob) GetStatusTitleOk() (*string, bool)`

GetStatusTitleOk returns a tuple with the StatusTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTitle

`func (o *TaskJob) SetStatusTitle(v string)`

SetStatusTitle sets StatusTitle field to given value.


### GetSteps

`func (o *TaskJob) GetSteps() []TaskStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TaskJob) GetStepsOk() (*[]TaskStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TaskJob) SetSteps(v []TaskStep)`

SetSteps sets Steps field to given value.


### GetCreatedAt

`func (o *TaskJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaskJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartedAt

`func (o *TaskJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TaskJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TaskJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TaskJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *TaskJob) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *TaskJob) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *TaskJob) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *TaskJob) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *TaskJob) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *TaskJob) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *TaskJob) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *TaskJob) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


