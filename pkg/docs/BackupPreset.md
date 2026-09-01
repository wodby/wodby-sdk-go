# BackupPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppInstanceId** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseDbId** | Pointer to **NullableInt32** |  | [optional] 
**OrgId** | Pointer to **NullableInt32** |  | [optional] 
**EnvId** | Pointer to **NullableInt32** |  | [optional] 
**EnvTypes** | **[]string** |  | 
**BackupCategory** | **string** |  | 
**BackupName** | Pointer to **NullableString** |  | [optional] 
**IntegrationId** | **int32** |  | 
**Bucket** | **string** |  | 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**Options** | [**[]BackupOption**](BackupOption.md) |  | 
**Override** | **bool** |  | 
**Auto** | **bool** |  | 
**Disabled** | **bool** |  | 
**Crontab** | Pointer to **NullableString** |  | [optional] 
**TimeWindow** | Pointer to [**AutomationTimeWindow**](AutomationTimeWindow.md) |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**NextRunAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBackupPreset

`func NewBackupPreset(id int32, envTypes []string, backupCategory string, integrationId int32, bucket string, options []BackupOption, override bool, auto bool, disabled bool, createdAt time.Time, updatedAt time.Time, ) *BackupPreset`

NewBackupPreset instantiates a new BackupPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPresetWithDefaults

`func NewBackupPresetWithDefaults() *BackupPreset`

NewBackupPresetWithDefaults instantiates a new BackupPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupPreset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupPreset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupPreset) SetId(v int32)`

SetId sets Id field to given value.


### GetAppInstanceId

`func (o *BackupPreset) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *BackupPreset) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *BackupPreset) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *BackupPreset) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### SetAppInstanceIdNil

`func (o *BackupPreset) SetAppInstanceIdNil(b bool)`

 SetAppInstanceIdNil sets the value for AppInstanceId to be an explicit nil

### UnsetAppInstanceId
`func (o *BackupPreset) UnsetAppInstanceId()`

UnsetAppInstanceId ensures that no value is present for AppInstanceId, not even an explicit nil
### GetAppServiceId

`func (o *BackupPreset) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *BackupPreset) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *BackupPreset) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *BackupPreset) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *BackupPreset) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *BackupPreset) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetDatabaseId

`func (o *BackupPreset) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *BackupPreset) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *BackupPreset) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *BackupPreset) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *BackupPreset) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *BackupPreset) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetDatabaseDbId

`func (o *BackupPreset) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *BackupPreset) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *BackupPreset) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *BackupPreset) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *BackupPreset) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *BackupPreset) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil
### GetOrgId

`func (o *BackupPreset) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BackupPreset) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BackupPreset) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BackupPreset) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *BackupPreset) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *BackupPreset) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetEnvId

`func (o *BackupPreset) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *BackupPreset) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *BackupPreset) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *BackupPreset) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### SetEnvIdNil

`func (o *BackupPreset) SetEnvIdNil(b bool)`

 SetEnvIdNil sets the value for EnvId to be an explicit nil

### UnsetEnvId
`func (o *BackupPreset) UnsetEnvId()`

UnsetEnvId ensures that no value is present for EnvId, not even an explicit nil
### GetEnvTypes

`func (o *BackupPreset) GetEnvTypes() []string`

GetEnvTypes returns the EnvTypes field if non-nil, zero value otherwise.

### GetEnvTypesOk

`func (o *BackupPreset) GetEnvTypesOk() (*[]string, bool)`

GetEnvTypesOk returns a tuple with the EnvTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvTypes

`func (o *BackupPreset) SetEnvTypes(v []string)`

SetEnvTypes sets EnvTypes field to given value.


### GetBackupCategory

`func (o *BackupPreset) GetBackupCategory() string`

GetBackupCategory returns the BackupCategory field if non-nil, zero value otherwise.

### GetBackupCategoryOk

`func (o *BackupPreset) GetBackupCategoryOk() (*string, bool)`

GetBackupCategoryOk returns a tuple with the BackupCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCategory

`func (o *BackupPreset) SetBackupCategory(v string)`

SetBackupCategory sets BackupCategory field to given value.


### GetBackupName

`func (o *BackupPreset) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *BackupPreset) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *BackupPreset) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *BackupPreset) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *BackupPreset) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *BackupPreset) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetIntegrationId

`func (o *BackupPreset) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *BackupPreset) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *BackupPreset) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetBucket

`func (o *BackupPreset) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *BackupPreset) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *BackupPreset) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetStorageClass

`func (o *BackupPreset) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *BackupPreset) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *BackupPreset) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *BackupPreset) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *BackupPreset) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *BackupPreset) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetOptions

`func (o *BackupPreset) GetOptions() []BackupOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BackupPreset) GetOptionsOk() (*[]BackupOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BackupPreset) SetOptions(v []BackupOption)`

SetOptions sets Options field to given value.


### GetOverride

`func (o *BackupPreset) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *BackupPreset) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *BackupPreset) SetOverride(v bool)`

SetOverride sets Override field to given value.


### GetAuto

`func (o *BackupPreset) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *BackupPreset) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *BackupPreset) SetAuto(v bool)`

SetAuto sets Auto field to given value.


### GetDisabled

`func (o *BackupPreset) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *BackupPreset) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *BackupPreset) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetCrontab

`func (o *BackupPreset) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *BackupPreset) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *BackupPreset) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *BackupPreset) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### SetCrontabNil

`func (o *BackupPreset) SetCrontabNil(b bool)`

 SetCrontabNil sets the value for Crontab to be an explicit nil

### UnsetCrontab
`func (o *BackupPreset) UnsetCrontab()`

UnsetCrontab ensures that no value is present for Crontab, not even an explicit nil
### GetTimeWindow

`func (o *BackupPreset) GetTimeWindow() AutomationTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *BackupPreset) GetTimeWindowOk() (*AutomationTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *BackupPreset) SetTimeWindow(v AutomationTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *BackupPreset) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetDuration

`func (o *BackupPreset) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BackupPreset) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BackupPreset) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BackupPreset) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *BackupPreset) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *BackupPreset) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetNextRunAt

`func (o *BackupPreset) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *BackupPreset) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *BackupPreset) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *BackupPreset) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *BackupPreset) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *BackupPreset) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetCreatedAt

`func (o *BackupPreset) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupPreset) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupPreset) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BackupPreset) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BackupPreset) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BackupPreset) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


