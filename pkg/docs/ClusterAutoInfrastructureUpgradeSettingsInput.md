# ClusterAutoInfrastructureUpgradeSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Infra** | Pointer to [**ClusterAutoInfrastructureComponentSettingsInput**](ClusterAutoInfrastructureComponentSettingsInput.md) |  | [optional] 
**Apps** | Pointer to [**ClusterAutoInfrastructureComponentSettingsInput**](ClusterAutoInfrastructureComponentSettingsInput.md) |  | [optional] 

## Methods

### NewClusterAutoInfrastructureUpgradeSettingsInput

`func NewClusterAutoInfrastructureUpgradeSettingsInput(enabled bool, ) *ClusterAutoInfrastructureUpgradeSettingsInput`

NewClusterAutoInfrastructureUpgradeSettingsInput instantiates a new ClusterAutoInfrastructureUpgradeSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoInfrastructureUpgradeSettingsInputWithDefaults

`func NewClusterAutoInfrastructureUpgradeSettingsInputWithDefaults() *ClusterAutoInfrastructureUpgradeSettingsInput`

NewClusterAutoInfrastructureUpgradeSettingsInputWithDefaults instantiates a new ClusterAutoInfrastructureUpgradeSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInfra

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) GetInfra() ClusterAutoInfrastructureComponentSettingsInput`

GetInfra returns the Infra field if non-nil, zero value otherwise.

### GetInfraOk

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) GetInfraOk() (*ClusterAutoInfrastructureComponentSettingsInput, bool)`

GetInfraOk returns a tuple with the Infra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfra

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) SetInfra(v ClusterAutoInfrastructureComponentSettingsInput)`

SetInfra sets Infra field to given value.

### HasInfra

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) HasInfra() bool`

HasInfra returns a boolean if a field has been set.

### GetApps

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) GetApps() ClusterAutoInfrastructureComponentSettingsInput`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) GetAppsOk() (*ClusterAutoInfrastructureComponentSettingsInput, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) SetApps(v ClusterAutoInfrastructureComponentSettingsInput)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ClusterAutoInfrastructureUpgradeSettingsInput) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


