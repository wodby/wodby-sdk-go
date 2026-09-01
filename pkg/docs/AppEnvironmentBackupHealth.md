# AppEnvironmentBackupHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailingSchedulesCount** | **int32** |  | 
**LatestFailureAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppEnvironmentBackupHealth

`func NewAppEnvironmentBackupHealth(failingSchedulesCount int32, ) *AppEnvironmentBackupHealth`

NewAppEnvironmentBackupHealth instantiates a new AppEnvironmentBackupHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentBackupHealthWithDefaults

`func NewAppEnvironmentBackupHealthWithDefaults() *AppEnvironmentBackupHealth`

NewAppEnvironmentBackupHealthWithDefaults instantiates a new AppEnvironmentBackupHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailingSchedulesCount

`func (o *AppEnvironmentBackupHealth) GetFailingSchedulesCount() int32`

GetFailingSchedulesCount returns the FailingSchedulesCount field if non-nil, zero value otherwise.

### GetFailingSchedulesCountOk

`func (o *AppEnvironmentBackupHealth) GetFailingSchedulesCountOk() (*int32, bool)`

GetFailingSchedulesCountOk returns a tuple with the FailingSchedulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingSchedulesCount

`func (o *AppEnvironmentBackupHealth) SetFailingSchedulesCount(v int32)`

SetFailingSchedulesCount sets FailingSchedulesCount field to given value.


### GetLatestFailureAt

`func (o *AppEnvironmentBackupHealth) GetLatestFailureAt() time.Time`

GetLatestFailureAt returns the LatestFailureAt field if non-nil, zero value otherwise.

### GetLatestFailureAtOk

`func (o *AppEnvironmentBackupHealth) GetLatestFailureAtOk() (*time.Time, bool)`

GetLatestFailureAtOk returns a tuple with the LatestFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFailureAt

`func (o *AppEnvironmentBackupHealth) SetLatestFailureAt(v time.Time)`

SetLatestFailureAt sets LatestFailureAt field to given value.

### HasLatestFailureAt

`func (o *AppEnvironmentBackupHealth) HasLatestFailureAt() bool`

HasLatestFailureAt returns a boolean if a field has been set.

### SetLatestFailureAtNil

`func (o *AppEnvironmentBackupHealth) SetLatestFailureAtNil(b bool)`

 SetLatestFailureAtNil sets the value for LatestFailureAt to be an explicit nil

### UnsetLatestFailureAt
`func (o *AppEnvironmentBackupHealth) UnsetLatestFailureAt()`

UnsetLatestFailureAt ensures that no value is present for LatestFailureAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


