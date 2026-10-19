# AppEnvironmentMigrationComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Name** | **string** |  | 
**BackupId** | **int32** |  | 
**BackupStatus** | Pointer to **NullableString** |  | [optional] 
**ImportId** | **int32** |  | 
**ImportStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppEnvironmentMigrationComponent

`func NewAppEnvironmentMigrationComponent(id int32, appServiceId int32, name string, backupId int32, importId int32, ) *AppEnvironmentMigrationComponent`

NewAppEnvironmentMigrationComponent instantiates a new AppEnvironmentMigrationComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentMigrationComponentWithDefaults

`func NewAppEnvironmentMigrationComponentWithDefaults() *AppEnvironmentMigrationComponent`

NewAppEnvironmentMigrationComponentWithDefaults instantiates a new AppEnvironmentMigrationComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppEnvironmentMigrationComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEnvironmentMigrationComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEnvironmentMigrationComponent) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppEnvironmentMigrationComponent) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppEnvironmentMigrationComponent) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppEnvironmentMigrationComponent) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetName

`func (o *AppEnvironmentMigrationComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppEnvironmentMigrationComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppEnvironmentMigrationComponent) SetName(v string)`

SetName sets Name field to given value.


### GetBackupId

`func (o *AppEnvironmentMigrationComponent) GetBackupId() int32`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *AppEnvironmentMigrationComponent) GetBackupIdOk() (*int32, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *AppEnvironmentMigrationComponent) SetBackupId(v int32)`

SetBackupId sets BackupId field to given value.


### GetBackupStatus

`func (o *AppEnvironmentMigrationComponent) GetBackupStatus() string`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *AppEnvironmentMigrationComponent) GetBackupStatusOk() (*string, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *AppEnvironmentMigrationComponent) SetBackupStatus(v string)`

SetBackupStatus sets BackupStatus field to given value.

### HasBackupStatus

`func (o *AppEnvironmentMigrationComponent) HasBackupStatus() bool`

HasBackupStatus returns a boolean if a field has been set.

### SetBackupStatusNil

`func (o *AppEnvironmentMigrationComponent) SetBackupStatusNil(b bool)`

 SetBackupStatusNil sets the value for BackupStatus to be an explicit nil

### UnsetBackupStatus
`func (o *AppEnvironmentMigrationComponent) UnsetBackupStatus()`

UnsetBackupStatus ensures that no value is present for BackupStatus, not even an explicit nil
### GetImportId

`func (o *AppEnvironmentMigrationComponent) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *AppEnvironmentMigrationComponent) GetImportIdOk() (*int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *AppEnvironmentMigrationComponent) SetImportId(v int32)`

SetImportId sets ImportId field to given value.


### GetImportStatus

`func (o *AppEnvironmentMigrationComponent) GetImportStatus() string`

GetImportStatus returns the ImportStatus field if non-nil, zero value otherwise.

### GetImportStatusOk

`func (o *AppEnvironmentMigrationComponent) GetImportStatusOk() (*string, bool)`

GetImportStatusOk returns a tuple with the ImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStatus

`func (o *AppEnvironmentMigrationComponent) SetImportStatus(v string)`

SetImportStatus sets ImportStatus field to given value.

### HasImportStatus

`func (o *AppEnvironmentMigrationComponent) HasImportStatus() bool`

HasImportStatus returns a boolean if a field has been set.

### SetImportStatusNil

`func (o *AppEnvironmentMigrationComponent) SetImportStatusNil(b bool)`

 SetImportStatusNil sets the value for ImportStatus to be an explicit nil

### UnsetImportStatus
`func (o *AppEnvironmentMigrationComponent) UnsetImportStatus()`

UnsetImportStatus ensures that no value is present for ImportStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


