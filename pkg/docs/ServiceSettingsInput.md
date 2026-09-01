# ServiceSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitAutoUpdate** | Pointer to [**GitAutoUpdateSettingsInput**](GitAutoUpdateSettingsInput.md) |  | [optional] 
**AutoBaseRevisionUpdate** | Pointer to [**ServiceAutoBaseRevisionUpdateSettingsInput**](ServiceAutoBaseRevisionUpdateSettingsInput.md) |  | [optional] 

## Methods

### NewServiceSettingsInput

`func NewServiceSettingsInput() *ServiceSettingsInput`

NewServiceSettingsInput instantiates a new ServiceSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSettingsInputWithDefaults

`func NewServiceSettingsInputWithDefaults() *ServiceSettingsInput`

NewServiceSettingsInputWithDefaults instantiates a new ServiceSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitAutoUpdate

`func (o *ServiceSettingsInput) GetGitAutoUpdate() GitAutoUpdateSettingsInput`

GetGitAutoUpdate returns the GitAutoUpdate field if non-nil, zero value otherwise.

### GetGitAutoUpdateOk

`func (o *ServiceSettingsInput) GetGitAutoUpdateOk() (*GitAutoUpdateSettingsInput, bool)`

GetGitAutoUpdateOk returns a tuple with the GitAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitAutoUpdate

`func (o *ServiceSettingsInput) SetGitAutoUpdate(v GitAutoUpdateSettingsInput)`

SetGitAutoUpdate sets GitAutoUpdate field to given value.

### HasGitAutoUpdate

`func (o *ServiceSettingsInput) HasGitAutoUpdate() bool`

HasGitAutoUpdate returns a boolean if a field has been set.

### GetAutoBaseRevisionUpdate

`func (o *ServiceSettingsInput) GetAutoBaseRevisionUpdate() ServiceAutoBaseRevisionUpdateSettingsInput`

GetAutoBaseRevisionUpdate returns the AutoBaseRevisionUpdate field if non-nil, zero value otherwise.

### GetAutoBaseRevisionUpdateOk

`func (o *ServiceSettingsInput) GetAutoBaseRevisionUpdateOk() (*ServiceAutoBaseRevisionUpdateSettingsInput, bool)`

GetAutoBaseRevisionUpdateOk returns a tuple with the AutoBaseRevisionUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBaseRevisionUpdate

`func (o *ServiceSettingsInput) SetAutoBaseRevisionUpdate(v ServiceAutoBaseRevisionUpdateSettingsInput)`

SetAutoBaseRevisionUpdate sets AutoBaseRevisionUpdate field to given value.

### HasAutoBaseRevisionUpdate

`func (o *ServiceSettingsInput) HasAutoBaseRevisionUpdate() bool`

HasAutoBaseRevisionUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


