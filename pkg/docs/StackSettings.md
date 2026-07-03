# StackSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitAutoUpdate** | Pointer to [**GitAutoUpdateSettings**](GitAutoUpdateSettings.md) |  | [optional] 
**AutoServiceRevisionUpdate** | Pointer to [**StackAutoServiceRevisionUpdateSettings**](StackAutoServiceRevisionUpdateSettings.md) |  | [optional] 
**AutoOriginStackUpdate** | Pointer to [**StackAutoOriginUpdateSettings**](StackAutoOriginUpdateSettings.md) |  | [optional] 

## Methods

### NewStackSettings

`func NewStackSettings() *StackSettings`

NewStackSettings instantiates a new StackSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackSettingsWithDefaults

`func NewStackSettingsWithDefaults() *StackSettings`

NewStackSettingsWithDefaults instantiates a new StackSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitAutoUpdate

`func (o *StackSettings) GetGitAutoUpdate() GitAutoUpdateSettings`

GetGitAutoUpdate returns the GitAutoUpdate field if non-nil, zero value otherwise.

### GetGitAutoUpdateOk

`func (o *StackSettings) GetGitAutoUpdateOk() (*GitAutoUpdateSettings, bool)`

GetGitAutoUpdateOk returns a tuple with the GitAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitAutoUpdate

`func (o *StackSettings) SetGitAutoUpdate(v GitAutoUpdateSettings)`

SetGitAutoUpdate sets GitAutoUpdate field to given value.

### HasGitAutoUpdate

`func (o *StackSettings) HasGitAutoUpdate() bool`

HasGitAutoUpdate returns a boolean if a field has been set.

### GetAutoServiceRevisionUpdate

`func (o *StackSettings) GetAutoServiceRevisionUpdate() StackAutoServiceRevisionUpdateSettings`

GetAutoServiceRevisionUpdate returns the AutoServiceRevisionUpdate field if non-nil, zero value otherwise.

### GetAutoServiceRevisionUpdateOk

`func (o *StackSettings) GetAutoServiceRevisionUpdateOk() (*StackAutoServiceRevisionUpdateSettings, bool)`

GetAutoServiceRevisionUpdateOk returns a tuple with the AutoServiceRevisionUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoServiceRevisionUpdate

`func (o *StackSettings) SetAutoServiceRevisionUpdate(v StackAutoServiceRevisionUpdateSettings)`

SetAutoServiceRevisionUpdate sets AutoServiceRevisionUpdate field to given value.

### HasAutoServiceRevisionUpdate

`func (o *StackSettings) HasAutoServiceRevisionUpdate() bool`

HasAutoServiceRevisionUpdate returns a boolean if a field has been set.

### GetAutoOriginStackUpdate

`func (o *StackSettings) GetAutoOriginStackUpdate() StackAutoOriginUpdateSettings`

GetAutoOriginStackUpdate returns the AutoOriginStackUpdate field if non-nil, zero value otherwise.

### GetAutoOriginStackUpdateOk

`func (o *StackSettings) GetAutoOriginStackUpdateOk() (*StackAutoOriginUpdateSettings, bool)`

GetAutoOriginStackUpdateOk returns a tuple with the AutoOriginStackUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOriginStackUpdate

`func (o *StackSettings) SetAutoOriginStackUpdate(v StackAutoOriginUpdateSettings)`

SetAutoOriginStackUpdate sets AutoOriginStackUpdate field to given value.

### HasAutoOriginStackUpdate

`func (o *StackSettings) HasAutoOriginStackUpdate() bool`

HasAutoOriginStackUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


