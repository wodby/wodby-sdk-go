# NewBackupPresetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseId** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseDbId** | Pointer to **NullableInt32** |  | [optional] 
**OrgId** | Pointer to **NullableInt32** | Optional for API-key requests; defaults to the API key&#39;s organization when no more specific target is provided. | [optional] 
**EnvId** | Pointer to **NullableInt32** |  | [optional] 
**BackupName** | Pointer to **NullableString** |  | [optional] 
**IntegrationId** | **int32** |  | 
**Bucket** | **string** |  | 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**Disabled** | **bool** |  | 
**Override** | **bool** |  | 
**Auto** | Pointer to **NullableBool** |  | [optional] 
**Crontab** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNewBackupPresetInput

`func NewNewBackupPresetInput(integrationId int32, bucket string, disabled bool, override bool, ) *NewBackupPresetInput`

NewNewBackupPresetInput instantiates a new NewBackupPresetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBackupPresetInputWithDefaults

`func NewNewBackupPresetInputWithDefaults() *NewBackupPresetInput`

NewNewBackupPresetInputWithDefaults instantiates a new NewBackupPresetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *NewBackupPresetInput) GetAppInstanceId() int32`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *NewBackupPresetInput) GetAppInstanceIdOk() (*int32, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *NewBackupPresetInput) SetAppInstanceId(v int32)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *NewBackupPresetInput) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### SetAppInstanceIdNil

`func (o *NewBackupPresetInput) SetAppInstanceIdNil(b bool)`

 SetAppInstanceIdNil sets the value for AppInstanceId to be an explicit nil

### UnsetAppInstanceId
`func (o *NewBackupPresetInput) UnsetAppInstanceId()`

UnsetAppInstanceId ensures that no value is present for AppInstanceId, not even an explicit nil
### GetAppServiceId

`func (o *NewBackupPresetInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *NewBackupPresetInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *NewBackupPresetInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *NewBackupPresetInput) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *NewBackupPresetInput) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *NewBackupPresetInput) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetDatabaseId

`func (o *NewBackupPresetInput) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *NewBackupPresetInput) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *NewBackupPresetInput) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *NewBackupPresetInput) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *NewBackupPresetInput) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *NewBackupPresetInput) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetDatabaseDbId

`func (o *NewBackupPresetInput) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *NewBackupPresetInput) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *NewBackupPresetInput) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *NewBackupPresetInput) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *NewBackupPresetInput) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *NewBackupPresetInput) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil
### GetOrgId

`func (o *NewBackupPresetInput) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NewBackupPresetInput) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NewBackupPresetInput) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NewBackupPresetInput) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *NewBackupPresetInput) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *NewBackupPresetInput) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetEnvId

`func (o *NewBackupPresetInput) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *NewBackupPresetInput) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *NewBackupPresetInput) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *NewBackupPresetInput) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### SetEnvIdNil

`func (o *NewBackupPresetInput) SetEnvIdNil(b bool)`

 SetEnvIdNil sets the value for EnvId to be an explicit nil

### UnsetEnvId
`func (o *NewBackupPresetInput) UnsetEnvId()`

UnsetEnvId ensures that no value is present for EnvId, not even an explicit nil
### GetBackupName

`func (o *NewBackupPresetInput) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *NewBackupPresetInput) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *NewBackupPresetInput) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *NewBackupPresetInput) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *NewBackupPresetInput) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *NewBackupPresetInput) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetIntegrationId

`func (o *NewBackupPresetInput) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewBackupPresetInput) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewBackupPresetInput) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.


### GetBucket

`func (o *NewBackupPresetInput) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *NewBackupPresetInput) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *NewBackupPresetInput) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetStorageClass

`func (o *NewBackupPresetInput) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *NewBackupPresetInput) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *NewBackupPresetInput) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *NewBackupPresetInput) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *NewBackupPresetInput) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *NewBackupPresetInput) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetDisabled

`func (o *NewBackupPresetInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NewBackupPresetInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NewBackupPresetInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOverride

`func (o *NewBackupPresetInput) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *NewBackupPresetInput) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *NewBackupPresetInput) SetOverride(v bool)`

SetOverride sets Override field to given value.


### GetAuto

`func (o *NewBackupPresetInput) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *NewBackupPresetInput) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *NewBackupPresetInput) SetAuto(v bool)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *NewBackupPresetInput) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### SetAutoNil

`func (o *NewBackupPresetInput) SetAutoNil(b bool)`

 SetAutoNil sets the value for Auto to be an explicit nil

### UnsetAuto
`func (o *NewBackupPresetInput) UnsetAuto()`

UnsetAuto ensures that no value is present for Auto, not even an explicit nil
### GetCrontab

`func (o *NewBackupPresetInput) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *NewBackupPresetInput) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *NewBackupPresetInput) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *NewBackupPresetInput) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### SetCrontabNil

`func (o *NewBackupPresetInput) SetCrontabNil(b bool)`

 SetCrontabNil sets the value for Crontab to be an explicit nil

### UnsetCrontab
`func (o *NewBackupPresetInput) UnsetCrontab()`

UnsetCrontab ensures that no value is present for Crontab, not even an explicit nil
### GetDuration

`func (o *NewBackupPresetInput) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NewBackupPresetInput) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NewBackupPresetInput) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NewBackupPresetInput) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *NewBackupPresetInput) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *NewBackupPresetInput) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


