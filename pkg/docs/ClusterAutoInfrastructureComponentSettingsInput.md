# ClusterAutoInfrastructureComponentSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**VersionPolicy** | Pointer to [**ClusterAutoUpgradeVersionPolicyInput**](ClusterAutoUpgradeVersionPolicyInput.md) |  | [optional] 

## Methods

### NewClusterAutoInfrastructureComponentSettingsInput

`func NewClusterAutoInfrastructureComponentSettingsInput() *ClusterAutoInfrastructureComponentSettingsInput`

NewClusterAutoInfrastructureComponentSettingsInput instantiates a new ClusterAutoInfrastructureComponentSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoInfrastructureComponentSettingsInputWithDefaults

`func NewClusterAutoInfrastructureComponentSettingsInputWithDefaults() *ClusterAutoInfrastructureComponentSettingsInput`

NewClusterAutoInfrastructureComponentSettingsInputWithDefaults instantiates a new ClusterAutoInfrastructureComponentSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAutoInfrastructureComponentSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAutoInfrastructureComponentSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAutoInfrastructureComponentSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterAutoInfrastructureComponentSettingsInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ClusterAutoInfrastructureComponentSettingsInput) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ClusterAutoInfrastructureComponentSettingsInput) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetVersionPolicy

`func (o *ClusterAutoInfrastructureComponentSettingsInput) GetVersionPolicy() ClusterAutoUpgradeVersionPolicyInput`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *ClusterAutoInfrastructureComponentSettingsInput) GetVersionPolicyOk() (*ClusterAutoUpgradeVersionPolicyInput, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *ClusterAutoInfrastructureComponentSettingsInput) SetVersionPolicy(v ClusterAutoUpgradeVersionPolicyInput)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *ClusterAutoInfrastructureComponentSettingsInput) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


