# ImportFromInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseDbId** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceId** | Pointer to **NullableInt32** |  | [optional] 
**BackupName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImportFromInput

`func NewImportFromInput() *ImportFromInput`

NewImportFromInput instantiates a new ImportFromInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFromInputWithDefaults

`func NewImportFromInputWithDefaults() *ImportFromInput`

NewImportFromInputWithDefaults instantiates a new ImportFromInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseDbId

`func (o *ImportFromInput) GetDatabaseDbId() int32`

GetDatabaseDbId returns the DatabaseDbId field if non-nil, zero value otherwise.

### GetDatabaseDbIdOk

`func (o *ImportFromInput) GetDatabaseDbIdOk() (*int32, bool)`

GetDatabaseDbIdOk returns a tuple with the DatabaseDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDbId

`func (o *ImportFromInput) SetDatabaseDbId(v int32)`

SetDatabaseDbId sets DatabaseDbId field to given value.

### HasDatabaseDbId

`func (o *ImportFromInput) HasDatabaseDbId() bool`

HasDatabaseDbId returns a boolean if a field has been set.

### SetDatabaseDbIdNil

`func (o *ImportFromInput) SetDatabaseDbIdNil(b bool)`

 SetDatabaseDbIdNil sets the value for DatabaseDbId to be an explicit nil

### UnsetDatabaseDbId
`func (o *ImportFromInput) UnsetDatabaseDbId()`

UnsetDatabaseDbId ensures that no value is present for DatabaseDbId, not even an explicit nil
### GetAppServiceId

`func (o *ImportFromInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *ImportFromInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *ImportFromInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.

### HasAppServiceId

`func (o *ImportFromInput) HasAppServiceId() bool`

HasAppServiceId returns a boolean if a field has been set.

### SetAppServiceIdNil

`func (o *ImportFromInput) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *ImportFromInput) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetBackupName

`func (o *ImportFromInput) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *ImportFromInput) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *ImportFromInput) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *ImportFromInput) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *ImportFromInput) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *ImportFromInput) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


