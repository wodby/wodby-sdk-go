# StackAutoServiceRevisionUpdateSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Policy** | Pointer to [**StackAutoUpdatePolicyInput**](StackAutoUpdatePolicyInput.md) |  | [optional] 

## Methods

### NewStackAutoServiceRevisionUpdateSettingsInput

`func NewStackAutoServiceRevisionUpdateSettingsInput(enabled bool, ) *StackAutoServiceRevisionUpdateSettingsInput`

NewStackAutoServiceRevisionUpdateSettingsInput instantiates a new StackAutoServiceRevisionUpdateSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackAutoServiceRevisionUpdateSettingsInputWithDefaults

`func NewStackAutoServiceRevisionUpdateSettingsInputWithDefaults() *StackAutoServiceRevisionUpdateSettingsInput`

NewStackAutoServiceRevisionUpdateSettingsInputWithDefaults instantiates a new StackAutoServiceRevisionUpdateSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StackAutoServiceRevisionUpdateSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StackAutoServiceRevisionUpdateSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StackAutoServiceRevisionUpdateSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPolicy

`func (o *StackAutoServiceRevisionUpdateSettingsInput) GetPolicy() StackAutoUpdatePolicyInput`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StackAutoServiceRevisionUpdateSettingsInput) GetPolicyOk() (*StackAutoUpdatePolicyInput, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StackAutoServiceRevisionUpdateSettingsInput) SetPolicy(v StackAutoUpdatePolicyInput)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StackAutoServiceRevisionUpdateSettingsInput) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


