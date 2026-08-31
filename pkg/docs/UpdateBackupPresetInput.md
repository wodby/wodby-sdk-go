# UpdateBackupPresetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **int32** | Use 0 for Wodby Blob Storage. Enabling the preset requires a paid subscription. | 
**Bucket** | **string** | Must be empty for Wodby Blob Storage. | 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]BackupOption**](BackupOption.md) |  | [optional] 
**Disabled** | **bool** |  | 
**Override** | **bool** |  | 
**Auto** | **bool** |  | 
**Crontab** | Pointer to **NullableString** |  | [optional] 
**TimeWindow** | Pointer to [**AutomationTimeWindowInput**](AutomationTimeWindowInput.md) |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateBackupPresetInput

`func NewUpdateBackupPresetInput(integrationId int32, bucket string, disabled bool, override bool, auto bool, ) *UpdateBackupPresetInput`

NewUpdateBackupPresetInput instantiates a new UpdateBackupPresetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackupPresetInputWithDefaults

`func NewUpdateBackupPresetInputWithDefaults() *UpdateBackupPresetInput`

NewUpdateBackupPresetInputWithDefaults instantiates a new UpdateBackupPresetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *UpdateBackupPresetInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *UpdateBackupPresetInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *UpdateBackupPresetInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetBucket

`func (o *UpdateBackupPresetInput) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *UpdateBackupPresetInput) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *UpdateBackupPresetInput) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetStorageClass

`func (o *UpdateBackupPresetInput) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *UpdateBackupPresetInput) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *UpdateBackupPresetInput) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *UpdateBackupPresetInput) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *UpdateBackupPresetInput) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *UpdateBackupPresetInput) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetOptions

`func (o *UpdateBackupPresetInput) GetOptions() []BackupOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateBackupPresetInput) GetOptionsOk() (*[]BackupOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateBackupPresetInput) SetOptions(v []BackupOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateBackupPresetInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateBackupPresetInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateBackupPresetInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDisabled

`func (o *UpdateBackupPresetInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateBackupPresetInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateBackupPresetInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOverride

`func (o *UpdateBackupPresetInput) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *UpdateBackupPresetInput) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *UpdateBackupPresetInput) SetOverride(v bool)`

SetOverride sets Override field to given value.


### GetAuto

`func (o *UpdateBackupPresetInput) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *UpdateBackupPresetInput) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *UpdateBackupPresetInput) SetAuto(v bool)`

SetAuto sets Auto field to given value.


### GetCrontab

`func (o *UpdateBackupPresetInput) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *UpdateBackupPresetInput) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *UpdateBackupPresetInput) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *UpdateBackupPresetInput) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### SetCrontabNil

`func (o *UpdateBackupPresetInput) SetCrontabNil(b bool)`

 SetCrontabNil sets the value for Crontab to be an explicit nil

### UnsetCrontab
`func (o *UpdateBackupPresetInput) UnsetCrontab()`

UnsetCrontab ensures that no value is present for Crontab, not even an explicit nil
### GetTimeWindow

`func (o *UpdateBackupPresetInput) GetTimeWindow() AutomationTimeWindowInput`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *UpdateBackupPresetInput) GetTimeWindowOk() (*AutomationTimeWindowInput, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *UpdateBackupPresetInput) SetTimeWindow(v AutomationTimeWindowInput)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *UpdateBackupPresetInput) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetDuration

`func (o *UpdateBackupPresetInput) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateBackupPresetInput) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateBackupPresetInput) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UpdateBackupPresetInput) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *UpdateBackupPresetInput) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *UpdateBackupPresetInput) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


