# NewAppServiceOverrideInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Stack service ID in the selected stack revision. | 
**Disabled** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**BuildSource** | Pointer to [**BuildSourceInput**](BuildSourceInput.md) |  | [optional] 
**Settings** | Pointer to [**[]AppServiceSettingInput**](AppServiceSettingInput.md) | Service settings to set. A setting that is not named keeps its default. Use it to supply a required setting that has no default. | [optional] 
**Database** | Pointer to [**AppServiceDatabaseInput**](AppServiceDatabaseInput.md) | The database server, and optionally the database, of a database service. Required to enable an external database service, such as a managed PostgreSQL server. | [optional] 

## Methods

### NewNewAppServiceOverrideInput

`func NewNewAppServiceOverrideInput(id int32, ) *NewAppServiceOverrideInput`

NewNewAppServiceOverrideInput instantiates a new NewAppServiceOverrideInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppServiceOverrideInputWithDefaults

`func NewNewAppServiceOverrideInputWithDefaults() *NewAppServiceOverrideInput`

NewNewAppServiceOverrideInputWithDefaults instantiates a new NewAppServiceOverrideInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewAppServiceOverrideInput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewAppServiceOverrideInput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewAppServiceOverrideInput) SetId(v int32)`

SetId sets Id field to given value.


### GetDisabled

`func (o *NewAppServiceOverrideInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NewAppServiceOverrideInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NewAppServiceOverrideInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NewAppServiceOverrideInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetVersion

`func (o *NewAppServiceOverrideInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewAppServiceOverrideInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewAppServiceOverrideInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NewAppServiceOverrideInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuildSource

`func (o *NewAppServiceOverrideInput) GetBuildSource() BuildSourceInput`

GetBuildSource returns the BuildSource field if non-nil, zero value otherwise.

### GetBuildSourceOk

`func (o *NewAppServiceOverrideInput) GetBuildSourceOk() (*BuildSourceInput, bool)`

GetBuildSourceOk returns a tuple with the BuildSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSource

`func (o *NewAppServiceOverrideInput) SetBuildSource(v BuildSourceInput)`

SetBuildSource sets BuildSource field to given value.

### HasBuildSource

`func (o *NewAppServiceOverrideInput) HasBuildSource() bool`

HasBuildSource returns a boolean if a field has been set.

### GetSettings

`func (o *NewAppServiceOverrideInput) GetSettings() []AppServiceSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppServiceOverrideInput) GetSettingsOk() (*[]AppServiceSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppServiceOverrideInput) SetSettings(v []AppServiceSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppServiceOverrideInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetDatabase

`func (o *NewAppServiceOverrideInput) GetDatabase() AppServiceDatabaseInput`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *NewAppServiceOverrideInput) GetDatabaseOk() (*AppServiceDatabaseInput, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *NewAppServiceOverrideInput) SetDatabase(v AppServiceDatabaseInput)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *NewAppServiceOverrideInput) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


