# StackAutoOriginUpdateSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**VersionPolicy** | Pointer to [**StackAutoOriginUpdateVersionPolicy**](StackAutoOriginUpdateVersionPolicy.md) |  | [optional] 
**SyncOptions** | Pointer to [**StackSyncOptions**](StackSyncOptions.md) |  | [optional] 

## Methods

### NewStackAutoOriginUpdateSettings

`func NewStackAutoOriginUpdateSettings(enabled bool, ) *StackAutoOriginUpdateSettings`

NewStackAutoOriginUpdateSettings instantiates a new StackAutoOriginUpdateSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackAutoOriginUpdateSettingsWithDefaults

`func NewStackAutoOriginUpdateSettingsWithDefaults() *StackAutoOriginUpdateSettings`

NewStackAutoOriginUpdateSettingsWithDefaults instantiates a new StackAutoOriginUpdateSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StackAutoOriginUpdateSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StackAutoOriginUpdateSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StackAutoOriginUpdateSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVersionPolicy

`func (o *StackAutoOriginUpdateSettings) GetVersionPolicy() StackAutoOriginUpdateVersionPolicy`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *StackAutoOriginUpdateSettings) GetVersionPolicyOk() (*StackAutoOriginUpdateVersionPolicy, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *StackAutoOriginUpdateSettings) SetVersionPolicy(v StackAutoOriginUpdateVersionPolicy)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *StackAutoOriginUpdateSettings) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.

### GetSyncOptions

`func (o *StackAutoOriginUpdateSettings) GetSyncOptions() StackSyncOptions`

GetSyncOptions returns the SyncOptions field if non-nil, zero value otherwise.

### GetSyncOptionsOk

`func (o *StackAutoOriginUpdateSettings) GetSyncOptionsOk() (*StackSyncOptions, bool)`

GetSyncOptionsOk returns a tuple with the SyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOptions

`func (o *StackAutoOriginUpdateSettings) SetSyncOptions(v StackSyncOptions)`

SetSyncOptions sets SyncOptions field to given value.

### HasSyncOptions

`func (o *StackAutoOriginUpdateSettings) HasSyncOptions() bool`

HasSyncOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


