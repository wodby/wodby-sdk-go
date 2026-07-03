# StackAutoServiceRevisionUpdateSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Policy** | Pointer to [**StackAutoUpdatePolicy**](StackAutoUpdatePolicy.md) |  | [optional] 

## Methods

### NewStackAutoServiceRevisionUpdateSettings

`func NewStackAutoServiceRevisionUpdateSettings(enabled bool, ) *StackAutoServiceRevisionUpdateSettings`

NewStackAutoServiceRevisionUpdateSettings instantiates a new StackAutoServiceRevisionUpdateSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackAutoServiceRevisionUpdateSettingsWithDefaults

`func NewStackAutoServiceRevisionUpdateSettingsWithDefaults() *StackAutoServiceRevisionUpdateSettings`

NewStackAutoServiceRevisionUpdateSettingsWithDefaults instantiates a new StackAutoServiceRevisionUpdateSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StackAutoServiceRevisionUpdateSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StackAutoServiceRevisionUpdateSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StackAutoServiceRevisionUpdateSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPolicy

`func (o *StackAutoServiceRevisionUpdateSettings) GetPolicy() StackAutoUpdatePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StackAutoServiceRevisionUpdateSettings) GetPolicyOk() (*StackAutoUpdatePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StackAutoServiceRevisionUpdateSettings) SetPolicy(v StackAutoUpdatePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StackAutoServiceRevisionUpdateSettings) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


