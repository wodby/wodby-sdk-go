# AppPreviewEnvironmentsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**BaseAppEnvironmentId** | Pointer to **NullableInt32** |  | [optional] 
**EnvironmentType** | **string** |  | 
**Imports** | [**[]AppPreviewImport**](AppPreviewImport.md) |  | 
**BuildDrafts** | **bool** |  | 
**RunCronJobs** | **bool** | Run the base environment&#39;s cron jobs in every preview. When false, the default, cron jobs are copied disabled. Previews never copy backup presets and run every service with one replica. | 
**MaxConcurrent** | **int32** |  | 
**IdleTtlDays** | **int32** |  | 

## Methods

### NewAppPreviewEnvironmentsSettings

`func NewAppPreviewEnvironmentsSettings(enabled bool, environmentType string, imports []AppPreviewImport, buildDrafts bool, runCronJobs bool, maxConcurrent int32, idleTtlDays int32, ) *AppPreviewEnvironmentsSettings`

NewAppPreviewEnvironmentsSettings instantiates a new AppPreviewEnvironmentsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewEnvironmentsSettingsWithDefaults

`func NewAppPreviewEnvironmentsSettingsWithDefaults() *AppPreviewEnvironmentsSettings`

NewAppPreviewEnvironmentsSettingsWithDefaults instantiates a new AppPreviewEnvironmentsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppPreviewEnvironmentsSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppPreviewEnvironmentsSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppPreviewEnvironmentsSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseAppEnvironmentId

`func (o *AppPreviewEnvironmentsSettings) GetBaseAppEnvironmentId() int32`

GetBaseAppEnvironmentId returns the BaseAppEnvironmentId field if non-nil, zero value otherwise.

### GetBaseAppEnvironmentIdOk

`func (o *AppPreviewEnvironmentsSettings) GetBaseAppEnvironmentIdOk() (*int32, bool)`

GetBaseAppEnvironmentIdOk returns a tuple with the BaseAppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAppEnvironmentId

`func (o *AppPreviewEnvironmentsSettings) SetBaseAppEnvironmentId(v int32)`

SetBaseAppEnvironmentId sets BaseAppEnvironmentId field to given value.

### HasBaseAppEnvironmentId

`func (o *AppPreviewEnvironmentsSettings) HasBaseAppEnvironmentId() bool`

HasBaseAppEnvironmentId returns a boolean if a field has been set.

### SetBaseAppEnvironmentIdNil

`func (o *AppPreviewEnvironmentsSettings) SetBaseAppEnvironmentIdNil(b bool)`

 SetBaseAppEnvironmentIdNil sets the value for BaseAppEnvironmentId to be an explicit nil

### UnsetBaseAppEnvironmentId
`func (o *AppPreviewEnvironmentsSettings) UnsetBaseAppEnvironmentId()`

UnsetBaseAppEnvironmentId ensures that no value is present for BaseAppEnvironmentId, not even an explicit nil
### GetEnvironmentType

`func (o *AppPreviewEnvironmentsSettings) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *AppPreviewEnvironmentsSettings) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *AppPreviewEnvironmentsSettings) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetImports

`func (o *AppPreviewEnvironmentsSettings) GetImports() []AppPreviewImport`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *AppPreviewEnvironmentsSettings) GetImportsOk() (*[]AppPreviewImport, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *AppPreviewEnvironmentsSettings) SetImports(v []AppPreviewImport)`

SetImports sets Imports field to given value.


### GetBuildDrafts

`func (o *AppPreviewEnvironmentsSettings) GetBuildDrafts() bool`

GetBuildDrafts returns the BuildDrafts field if non-nil, zero value otherwise.

### GetBuildDraftsOk

`func (o *AppPreviewEnvironmentsSettings) GetBuildDraftsOk() (*bool, bool)`

GetBuildDraftsOk returns a tuple with the BuildDrafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDrafts

`func (o *AppPreviewEnvironmentsSettings) SetBuildDrafts(v bool)`

SetBuildDrafts sets BuildDrafts field to given value.


### GetRunCronJobs

`func (o *AppPreviewEnvironmentsSettings) GetRunCronJobs() bool`

GetRunCronJobs returns the RunCronJobs field if non-nil, zero value otherwise.

### GetRunCronJobsOk

`func (o *AppPreviewEnvironmentsSettings) GetRunCronJobsOk() (*bool, bool)`

GetRunCronJobsOk returns a tuple with the RunCronJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCronJobs

`func (o *AppPreviewEnvironmentsSettings) SetRunCronJobs(v bool)`

SetRunCronJobs sets RunCronJobs field to given value.


### GetMaxConcurrent

`func (o *AppPreviewEnvironmentsSettings) GetMaxConcurrent() int32`

GetMaxConcurrent returns the MaxConcurrent field if non-nil, zero value otherwise.

### GetMaxConcurrentOk

`func (o *AppPreviewEnvironmentsSettings) GetMaxConcurrentOk() (*int32, bool)`

GetMaxConcurrentOk returns a tuple with the MaxConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrent

`func (o *AppPreviewEnvironmentsSettings) SetMaxConcurrent(v int32)`

SetMaxConcurrent sets MaxConcurrent field to given value.


### GetIdleTtlDays

`func (o *AppPreviewEnvironmentsSettings) GetIdleTtlDays() int32`

GetIdleTtlDays returns the IdleTtlDays field if non-nil, zero value otherwise.

### GetIdleTtlDaysOk

`func (o *AppPreviewEnvironmentsSettings) GetIdleTtlDaysOk() (*int32, bool)`

GetIdleTtlDaysOk returns a tuple with the IdleTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTtlDays

`func (o *AppPreviewEnvironmentsSettings) SetIdleTtlDays(v int32)`

SetIdleTtlDays sets IdleTtlDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


