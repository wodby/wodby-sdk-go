# ImportFromInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseDBID** | Pointer to **NullableInt32** |  | [optional] 
**AppServiceID** | Pointer to **NullableInt32** |  | [optional] 
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

### GetDatabaseDBID

`func (o *ImportFromInput) GetDatabaseDBID() int32`

GetDatabaseDBID returns the DatabaseDBID field if non-nil, zero value otherwise.

### GetDatabaseDBIDOk

`func (o *ImportFromInput) GetDatabaseDBIDOk() (*int32, bool)`

GetDatabaseDBIDOk returns a tuple with the DatabaseDBID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDBID

`func (o *ImportFromInput) SetDatabaseDBID(v int32)`

SetDatabaseDBID sets DatabaseDBID field to given value.

### HasDatabaseDBID

`func (o *ImportFromInput) HasDatabaseDBID() bool`

HasDatabaseDBID returns a boolean if a field has been set.

### SetDatabaseDBIDNil

`func (o *ImportFromInput) SetDatabaseDBIDNil(b bool)`

 SetDatabaseDBIDNil sets the value for DatabaseDBID to be an explicit nil

### UnsetDatabaseDBID
`func (o *ImportFromInput) UnsetDatabaseDBID()`

UnsetDatabaseDBID ensures that no value is present for DatabaseDBID, not even an explicit nil
### GetAppServiceID

`func (o *ImportFromInput) GetAppServiceID() int32`

GetAppServiceID returns the AppServiceID field if non-nil, zero value otherwise.

### GetAppServiceIDOk

`func (o *ImportFromInput) GetAppServiceIDOk() (*int32, bool)`

GetAppServiceIDOk returns a tuple with the AppServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceID

`func (o *ImportFromInput) SetAppServiceID(v int32)`

SetAppServiceID sets AppServiceID field to given value.

### HasAppServiceID

`func (o *ImportFromInput) HasAppServiceID() bool`

HasAppServiceID returns a boolean if a field has been set.

### SetAppServiceIDNil

`func (o *ImportFromInput) SetAppServiceIDNil(b bool)`

 SetAppServiceIDNil sets the value for AppServiceID to be an explicit nil

### UnsetAppServiceID
`func (o *ImportFromInput) UnsetAppServiceID()`

UnsetAppServiceID ensures that no value is present for AppServiceID, not even an explicit nil
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


