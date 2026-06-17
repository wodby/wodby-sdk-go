# NewAppServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Disabled** | **bool** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 
**BuildSource** | Pointer to [**BuildSourceInput**](BuildSourceInput.md) |  | [optional] 
**Imports** | Pointer to [**[]ImportInput**](ImportInput.md) |  | [optional] 
**Volumes** | Pointer to [**[]VolumeSizeInput**](VolumeSizeInput.md) |  | [optional] 
**Integrations** | Pointer to [**[]AppServiceIntegrationInput**](AppServiceIntegrationInput.md) |  | [optional] 
**Settings** | Pointer to [**[]AppServiceSettingInput**](AppServiceSettingInput.md) |  | [optional] 
**Database** | Pointer to [**AppServiceDatabaseInput**](AppServiceDatabaseInput.md) |  | [optional] 
**Resources** | Pointer to [**ResourcesInput**](ResourcesInput.md) |  | [optional] 
**Scalability** | Pointer to [**ScalabilityInput**](ScalabilityInput.md) |  | [optional] 

## Methods

### NewNewAppServiceInput

`func NewNewAppServiceInput(id int32, disabled bool, ) *NewAppServiceInput`

NewNewAppServiceInput instantiates a new NewAppServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppServiceInputWithDefaults

`func NewNewAppServiceInputWithDefaults() *NewAppServiceInput`

NewNewAppServiceInputWithDefaults instantiates a new NewAppServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewAppServiceInput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewAppServiceInput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewAppServiceInput) SetId(v int32)`

SetId sets Id field to given value.


### GetDisabled

`func (o *NewAppServiceInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NewAppServiceInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NewAppServiceInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetVersion

`func (o *NewAppServiceInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewAppServiceInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewAppServiceInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NewAppServiceInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NewAppServiceInput) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NewAppServiceInput) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildSource

`func (o *NewAppServiceInput) GetBuildSource() BuildSourceInput`

GetBuildSource returns the BuildSource field if non-nil, zero value otherwise.

### GetBuildSourceOk

`func (o *NewAppServiceInput) GetBuildSourceOk() (*BuildSourceInput, bool)`

GetBuildSourceOk returns a tuple with the BuildSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSource

`func (o *NewAppServiceInput) SetBuildSource(v BuildSourceInput)`

SetBuildSource sets BuildSource field to given value.

### HasBuildSource

`func (o *NewAppServiceInput) HasBuildSource() bool`

HasBuildSource returns a boolean if a field has been set.

### GetImports

`func (o *NewAppServiceInput) GetImports() []ImportInput`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *NewAppServiceInput) GetImportsOk() (*[]ImportInput, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *NewAppServiceInput) SetImports(v []ImportInput)`

SetImports sets Imports field to given value.

### HasImports

`func (o *NewAppServiceInput) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetVolumes

`func (o *NewAppServiceInput) GetVolumes() []VolumeSizeInput`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *NewAppServiceInput) GetVolumesOk() (*[]VolumeSizeInput, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *NewAppServiceInput) SetVolumes(v []VolumeSizeInput)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *NewAppServiceInput) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetIntegrations

`func (o *NewAppServiceInput) GetIntegrations() []AppServiceIntegrationInput`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *NewAppServiceInput) GetIntegrationsOk() (*[]AppServiceIntegrationInput, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *NewAppServiceInput) SetIntegrations(v []AppServiceIntegrationInput)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *NewAppServiceInput) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetSettings

`func (o *NewAppServiceInput) GetSettings() []AppServiceSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewAppServiceInput) GetSettingsOk() (*[]AppServiceSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewAppServiceInput) SetSettings(v []AppServiceSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewAppServiceInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetDatabase

`func (o *NewAppServiceInput) GetDatabase() AppServiceDatabaseInput`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *NewAppServiceInput) GetDatabaseOk() (*AppServiceDatabaseInput, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *NewAppServiceInput) SetDatabase(v AppServiceDatabaseInput)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *NewAppServiceInput) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetResources

`func (o *NewAppServiceInput) GetResources() ResourcesInput`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NewAppServiceInput) GetResourcesOk() (*ResourcesInput, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NewAppServiceInput) SetResources(v ResourcesInput)`

SetResources sets Resources field to given value.

### HasResources

`func (o *NewAppServiceInput) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetScalability

`func (o *NewAppServiceInput) GetScalability() ScalabilityInput`

GetScalability returns the Scalability field if non-nil, zero value otherwise.

### GetScalabilityOk

`func (o *NewAppServiceInput) GetScalabilityOk() (*ScalabilityInput, bool)`

GetScalabilityOk returns a tuple with the Scalability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalability

`func (o *NewAppServiceInput) SetScalability(v ScalabilityInput)`

SetScalability sets Scalability field to given value.

### HasScalability

`func (o *NewAppServiceInput) HasScalability() bool`

HasScalability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


