# AppPreviewEnvironmentsSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**BaseAppEnvironmentId** | Pointer to **int32** |  | [optional] 
**EnvironmentType** | Pointer to **string** |  | [optional] 
**Imports** | Pointer to [**[]AppPreviewImport**](AppPreviewImport.md) |  | [optional] 
**BuildDrafts** | Pointer to **bool** |  | [optional] 
**RunCronJobs** | Pointer to **bool** | Omitted keeps the stored value; apps that never set it do not run cron jobs in previews. | [optional] 
**MaxConcurrent** | Pointer to **int32** |  | [optional] 
**IdleTtlDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewAppPreviewEnvironmentsSettingsInput

`func NewAppPreviewEnvironmentsSettingsInput(enabled bool, ) *AppPreviewEnvironmentsSettingsInput`

NewAppPreviewEnvironmentsSettingsInput instantiates a new AppPreviewEnvironmentsSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewEnvironmentsSettingsInputWithDefaults

`func NewAppPreviewEnvironmentsSettingsInputWithDefaults() *AppPreviewEnvironmentsSettingsInput`

NewAppPreviewEnvironmentsSettingsInputWithDefaults instantiates a new AppPreviewEnvironmentsSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppPreviewEnvironmentsSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppPreviewEnvironmentsSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseAppEnvironmentId

`func (o *AppPreviewEnvironmentsSettingsInput) GetBaseAppEnvironmentId() int32`

GetBaseAppEnvironmentId returns the BaseAppEnvironmentId field if non-nil, zero value otherwise.

### GetBaseAppEnvironmentIdOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetBaseAppEnvironmentIdOk() (*int32, bool)`

GetBaseAppEnvironmentIdOk returns a tuple with the BaseAppEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAppEnvironmentId

`func (o *AppPreviewEnvironmentsSettingsInput) SetBaseAppEnvironmentId(v int32)`

SetBaseAppEnvironmentId sets BaseAppEnvironmentId field to given value.

### HasBaseAppEnvironmentId

`func (o *AppPreviewEnvironmentsSettingsInput) HasBaseAppEnvironmentId() bool`

HasBaseAppEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *AppPreviewEnvironmentsSettingsInput) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *AppPreviewEnvironmentsSettingsInput) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *AppPreviewEnvironmentsSettingsInput) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetImports

`func (o *AppPreviewEnvironmentsSettingsInput) GetImports() []AppPreviewImport`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetImportsOk() (*[]AppPreviewImport, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *AppPreviewEnvironmentsSettingsInput) SetImports(v []AppPreviewImport)`

SetImports sets Imports field to given value.

### HasImports

`func (o *AppPreviewEnvironmentsSettingsInput) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetBuildDrafts

`func (o *AppPreviewEnvironmentsSettingsInput) GetBuildDrafts() bool`

GetBuildDrafts returns the BuildDrafts field if non-nil, zero value otherwise.

### GetBuildDraftsOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetBuildDraftsOk() (*bool, bool)`

GetBuildDraftsOk returns a tuple with the BuildDrafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDrafts

`func (o *AppPreviewEnvironmentsSettingsInput) SetBuildDrafts(v bool)`

SetBuildDrafts sets BuildDrafts field to given value.

### HasBuildDrafts

`func (o *AppPreviewEnvironmentsSettingsInput) HasBuildDrafts() bool`

HasBuildDrafts returns a boolean if a field has been set.

### GetRunCronJobs

`func (o *AppPreviewEnvironmentsSettingsInput) GetRunCronJobs() bool`

GetRunCronJobs returns the RunCronJobs field if non-nil, zero value otherwise.

### GetRunCronJobsOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetRunCronJobsOk() (*bool, bool)`

GetRunCronJobsOk returns a tuple with the RunCronJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCronJobs

`func (o *AppPreviewEnvironmentsSettingsInput) SetRunCronJobs(v bool)`

SetRunCronJobs sets RunCronJobs field to given value.

### HasRunCronJobs

`func (o *AppPreviewEnvironmentsSettingsInput) HasRunCronJobs() bool`

HasRunCronJobs returns a boolean if a field has been set.

### GetMaxConcurrent

`func (o *AppPreviewEnvironmentsSettingsInput) GetMaxConcurrent() int32`

GetMaxConcurrent returns the MaxConcurrent field if non-nil, zero value otherwise.

### GetMaxConcurrentOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetMaxConcurrentOk() (*int32, bool)`

GetMaxConcurrentOk returns a tuple with the MaxConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrent

`func (o *AppPreviewEnvironmentsSettingsInput) SetMaxConcurrent(v int32)`

SetMaxConcurrent sets MaxConcurrent field to given value.

### HasMaxConcurrent

`func (o *AppPreviewEnvironmentsSettingsInput) HasMaxConcurrent() bool`

HasMaxConcurrent returns a boolean if a field has been set.

### GetIdleTtlDays

`func (o *AppPreviewEnvironmentsSettingsInput) GetIdleTtlDays() int32`

GetIdleTtlDays returns the IdleTtlDays field if non-nil, zero value otherwise.

### GetIdleTtlDaysOk

`func (o *AppPreviewEnvironmentsSettingsInput) GetIdleTtlDaysOk() (*int32, bool)`

GetIdleTtlDaysOk returns a tuple with the IdleTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTtlDays

`func (o *AppPreviewEnvironmentsSettingsInput) SetIdleTtlDays(v int32)`

SetIdleTtlDays sets IdleTtlDays field to given value.

### HasIdleTtlDays

`func (o *AppPreviewEnvironmentsSettingsInput) HasIdleTtlDays() bool`

HasIdleTtlDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


