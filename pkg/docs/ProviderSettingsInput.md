# ProviderSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitAutoUpdate** | [**GitAutoUpdateSettingsInput**](GitAutoUpdateSettingsInput.md) |  | 

## Methods

### NewProviderSettingsInput

`func NewProviderSettingsInput(gitAutoUpdate GitAutoUpdateSettingsInput, ) *ProviderSettingsInput`

NewProviderSettingsInput instantiates a new ProviderSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingsInputWithDefaults

`func NewProviderSettingsInputWithDefaults() *ProviderSettingsInput`

NewProviderSettingsInputWithDefaults instantiates a new ProviderSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitAutoUpdate

`func (o *ProviderSettingsInput) GetGitAutoUpdate() GitAutoUpdateSettingsInput`

GetGitAutoUpdate returns the GitAutoUpdate field if non-nil, zero value otherwise.

### GetGitAutoUpdateOk

`func (o *ProviderSettingsInput) GetGitAutoUpdateOk() (*GitAutoUpdateSettingsInput, bool)`

GetGitAutoUpdateOk returns a tuple with the GitAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitAutoUpdate

`func (o *ProviderSettingsInput) SetGitAutoUpdate(v GitAutoUpdateSettingsInput)`

SetGitAutoUpdate sets GitAutoUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


