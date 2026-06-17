# ImportInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportName** | Pointer to **NullableString** |  | [optional] 
**Source** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**BackupID** | Pointer to **NullableInt32** |  | [optional] 
**From** | Pointer to [**ImportFromInput**](ImportFromInput.md) |  | [optional] 

## Methods

### NewImportInput

`func NewImportInput(source string, ) *ImportInput`

NewImportInput instantiates a new ImportInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportInputWithDefaults

`func NewImportInputWithDefaults() *ImportInput`

NewImportInputWithDefaults instantiates a new ImportInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportName

`func (o *ImportInput) GetImportName() string`

GetImportName returns the ImportName field if non-nil, zero value otherwise.

### GetImportNameOk

`func (o *ImportInput) GetImportNameOk() (*string, bool)`

GetImportNameOk returns a tuple with the ImportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportName

`func (o *ImportInput) SetImportName(v string)`

SetImportName sets ImportName field to given value.

### HasImportName

`func (o *ImportInput) HasImportName() bool`

HasImportName returns a boolean if a field has been set.

### SetImportNameNil

`func (o *ImportInput) SetImportNameNil(b bool)`

 SetImportNameNil sets the value for ImportName to be an explicit nil

### UnsetImportName
`func (o *ImportInput) UnsetImportName()`

UnsetImportName ensures that no value is present for ImportName, not even an explicit nil
### GetSource

`func (o *ImportInput) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImportInput) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImportInput) SetSource(v string)`

SetSource sets Source field to given value.


### GetUrl

`func (o *ImportInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImportInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImportInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImportInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ImportInput) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ImportInput) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetBackupID

`func (o *ImportInput) GetBackupID() int32`

GetBackupID returns the BackupID field if non-nil, zero value otherwise.

### GetBackupIDOk

`func (o *ImportInput) GetBackupIDOk() (*int32, bool)`

GetBackupIDOk returns a tuple with the BackupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupID

`func (o *ImportInput) SetBackupID(v int32)`

SetBackupID sets BackupID field to given value.

### HasBackupID

`func (o *ImportInput) HasBackupID() bool`

HasBackupID returns a boolean if a field has been set.

### SetBackupIDNil

`func (o *ImportInput) SetBackupIDNil(b bool)`

 SetBackupIDNil sets the value for BackupID to be an explicit nil

### UnsetBackupID
`func (o *ImportInput) UnsetBackupID()`

UnsetBackupID ensures that no value is present for BackupID, not even an explicit nil
### GetFrom

`func (o *ImportInput) GetFrom() ImportFromInput`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ImportInput) GetFromOk() (*ImportFromInput, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ImportInput) SetFrom(v ImportFromInput)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ImportInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


