# ServiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitAutoUpdate** | Pointer to [**GitAutoUpdateSettings**](GitAutoUpdateSettings.md) |  | [optional] 

## Methods

### NewServiceSettings

`func NewServiceSettings() *ServiceSettings`

NewServiceSettings instantiates a new ServiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSettingsWithDefaults

`func NewServiceSettingsWithDefaults() *ServiceSettings`

NewServiceSettingsWithDefaults instantiates a new ServiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitAutoUpdate

`func (o *ServiceSettings) GetGitAutoUpdate() GitAutoUpdateSettings`

GetGitAutoUpdate returns the GitAutoUpdate field if non-nil, zero value otherwise.

### GetGitAutoUpdateOk

`func (o *ServiceSettings) GetGitAutoUpdateOk() (*GitAutoUpdateSettings, bool)`

GetGitAutoUpdateOk returns a tuple with the GitAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitAutoUpdate

`func (o *ServiceSettings) SetGitAutoUpdate(v GitAutoUpdateSettings)`

SetGitAutoUpdate sets GitAutoUpdate field to given value.

### HasGitAutoUpdate

`func (o *ServiceSettings) HasGitAutoUpdate() bool`

HasGitAutoUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


