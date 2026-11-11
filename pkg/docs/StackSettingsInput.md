# StackSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitAutoUpdate** | Pointer to [**GitAutoUpdateSettingsInput**](GitAutoUpdateSettingsInput.md) |  | [optional] 
**AutoServiceRevisionUpdate** | Pointer to [**StackAutoServiceRevisionUpdateSettingsInput**](StackAutoServiceRevisionUpdateSettingsInput.md) |  | [optional] 
**AutoOriginStackUpdate** | Pointer to [**StackAutoOriginUpdateSettingsInput**](StackAutoOriginUpdateSettingsInput.md) |  | [optional] 
**GitServiceUpdates** | Pointer to **NullableString** | What an automatic update from Git does with existing stack services&#39; versions; see StackSettings. | [optional] 

## Methods

### NewStackSettingsInput

`func NewStackSettingsInput() *StackSettingsInput`

NewStackSettingsInput instantiates a new StackSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackSettingsInputWithDefaults

`func NewStackSettingsInputWithDefaults() *StackSettingsInput`

NewStackSettingsInputWithDefaults instantiates a new StackSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitAutoUpdate

`func (o *StackSettingsInput) GetGitAutoUpdate() GitAutoUpdateSettingsInput`

GetGitAutoUpdate returns the GitAutoUpdate field if non-nil, zero value otherwise.

### GetGitAutoUpdateOk

`func (o *StackSettingsInput) GetGitAutoUpdateOk() (*GitAutoUpdateSettingsInput, bool)`

GetGitAutoUpdateOk returns a tuple with the GitAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitAutoUpdate

`func (o *StackSettingsInput) SetGitAutoUpdate(v GitAutoUpdateSettingsInput)`

SetGitAutoUpdate sets GitAutoUpdate field to given value.

### HasGitAutoUpdate

`func (o *StackSettingsInput) HasGitAutoUpdate() bool`

HasGitAutoUpdate returns a boolean if a field has been set.

### GetAutoServiceRevisionUpdate

`func (o *StackSettingsInput) GetAutoServiceRevisionUpdate() StackAutoServiceRevisionUpdateSettingsInput`

GetAutoServiceRevisionUpdate returns the AutoServiceRevisionUpdate field if non-nil, zero value otherwise.

### GetAutoServiceRevisionUpdateOk

`func (o *StackSettingsInput) GetAutoServiceRevisionUpdateOk() (*StackAutoServiceRevisionUpdateSettingsInput, bool)`

GetAutoServiceRevisionUpdateOk returns a tuple with the AutoServiceRevisionUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoServiceRevisionUpdate

`func (o *StackSettingsInput) SetAutoServiceRevisionUpdate(v StackAutoServiceRevisionUpdateSettingsInput)`

SetAutoServiceRevisionUpdate sets AutoServiceRevisionUpdate field to given value.

### HasAutoServiceRevisionUpdate

`func (o *StackSettingsInput) HasAutoServiceRevisionUpdate() bool`

HasAutoServiceRevisionUpdate returns a boolean if a field has been set.

### GetAutoOriginStackUpdate

`func (o *StackSettingsInput) GetAutoOriginStackUpdate() StackAutoOriginUpdateSettingsInput`

GetAutoOriginStackUpdate returns the AutoOriginStackUpdate field if non-nil, zero value otherwise.

### GetAutoOriginStackUpdateOk

`func (o *StackSettingsInput) GetAutoOriginStackUpdateOk() (*StackAutoOriginUpdateSettingsInput, bool)`

GetAutoOriginStackUpdateOk returns a tuple with the AutoOriginStackUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOriginStackUpdate

`func (o *StackSettingsInput) SetAutoOriginStackUpdate(v StackAutoOriginUpdateSettingsInput)`

SetAutoOriginStackUpdate sets AutoOriginStackUpdate field to given value.

### HasAutoOriginStackUpdate

`func (o *StackSettingsInput) HasAutoOriginStackUpdate() bool`

HasAutoOriginStackUpdate returns a boolean if a field has been set.

### GetGitServiceUpdates

`func (o *StackSettingsInput) GetGitServiceUpdates() string`

GetGitServiceUpdates returns the GitServiceUpdates field if non-nil, zero value otherwise.

### GetGitServiceUpdatesOk

`func (o *StackSettingsInput) GetGitServiceUpdatesOk() (*string, bool)`

GetGitServiceUpdatesOk returns a tuple with the GitServiceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitServiceUpdates

`func (o *StackSettingsInput) SetGitServiceUpdates(v string)`

SetGitServiceUpdates sets GitServiceUpdates field to given value.

### HasGitServiceUpdates

`func (o *StackSettingsInput) HasGitServiceUpdates() bool`

HasGitServiceUpdates returns a boolean if a field has been set.

### SetGitServiceUpdatesNil

`func (o *StackSettingsInput) SetGitServiceUpdatesNil(b bool)`

 SetGitServiceUpdatesNil sets the value for GitServiceUpdates to be an explicit nil

### UnsetGitServiceUpdates
`func (o *StackSettingsInput) UnsetGitServiceUpdates()`

UnsetGitServiceUpdates ensures that no value is present for GitServiceUpdates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


