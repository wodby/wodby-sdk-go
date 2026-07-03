# StackAutoOriginUpdateSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**VersionPolicy** | Pointer to [**StackAutoOriginUpdateVersionPolicyInput**](StackAutoOriginUpdateVersionPolicyInput.md) |  | [optional] 
**SyncOptions** | Pointer to [**StackSyncOptionsInput**](StackSyncOptionsInput.md) |  | [optional] 

## Methods

### NewStackAutoOriginUpdateSettingsInput

`func NewStackAutoOriginUpdateSettingsInput(enabled bool, ) *StackAutoOriginUpdateSettingsInput`

NewStackAutoOriginUpdateSettingsInput instantiates a new StackAutoOriginUpdateSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackAutoOriginUpdateSettingsInputWithDefaults

`func NewStackAutoOriginUpdateSettingsInputWithDefaults() *StackAutoOriginUpdateSettingsInput`

NewStackAutoOriginUpdateSettingsInputWithDefaults instantiates a new StackAutoOriginUpdateSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StackAutoOriginUpdateSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StackAutoOriginUpdateSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StackAutoOriginUpdateSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVersionPolicy

`func (o *StackAutoOriginUpdateSettingsInput) GetVersionPolicy() StackAutoOriginUpdateVersionPolicyInput`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *StackAutoOriginUpdateSettingsInput) GetVersionPolicyOk() (*StackAutoOriginUpdateVersionPolicyInput, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *StackAutoOriginUpdateSettingsInput) SetVersionPolicy(v StackAutoOriginUpdateVersionPolicyInput)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *StackAutoOriginUpdateSettingsInput) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.

### GetSyncOptions

`func (o *StackAutoOriginUpdateSettingsInput) GetSyncOptions() StackSyncOptionsInput`

GetSyncOptions returns the SyncOptions field if non-nil, zero value otherwise.

### GetSyncOptionsOk

`func (o *StackAutoOriginUpdateSettingsInput) GetSyncOptionsOk() (*StackSyncOptionsInput, bool)`

GetSyncOptionsOk returns a tuple with the SyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOptions

`func (o *StackAutoOriginUpdateSettingsInput) SetSyncOptions(v StackSyncOptionsInput)`

SetSyncOptions sets SyncOptions field to given value.

### HasSyncOptions

`func (o *StackAutoOriginUpdateSettingsInput) HasSyncOptions() bool`

HasSyncOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


