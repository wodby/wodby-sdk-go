# ClusterAutoInfrastructureUpgradeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Infra** | Pointer to [**ClusterAutoInfrastructureComponentSettings**](ClusterAutoInfrastructureComponentSettings.md) |  | [optional] 
**Apps** | Pointer to [**ClusterAutoInfrastructureComponentSettings**](ClusterAutoInfrastructureComponentSettings.md) |  | [optional] 

## Methods

### NewClusterAutoInfrastructureUpgradeSettings

`func NewClusterAutoInfrastructureUpgradeSettings(enabled bool, ) *ClusterAutoInfrastructureUpgradeSettings`

NewClusterAutoInfrastructureUpgradeSettings instantiates a new ClusterAutoInfrastructureUpgradeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoInfrastructureUpgradeSettingsWithDefaults

`func NewClusterAutoInfrastructureUpgradeSettingsWithDefaults() *ClusterAutoInfrastructureUpgradeSettings`

NewClusterAutoInfrastructureUpgradeSettingsWithDefaults instantiates a new ClusterAutoInfrastructureUpgradeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAutoInfrastructureUpgradeSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAutoInfrastructureUpgradeSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAutoInfrastructureUpgradeSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInfra

`func (o *ClusterAutoInfrastructureUpgradeSettings) GetInfra() ClusterAutoInfrastructureComponentSettings`

GetInfra returns the Infra field if non-nil, zero value otherwise.

### GetInfraOk

`func (o *ClusterAutoInfrastructureUpgradeSettings) GetInfraOk() (*ClusterAutoInfrastructureComponentSettings, bool)`

GetInfraOk returns a tuple with the Infra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfra

`func (o *ClusterAutoInfrastructureUpgradeSettings) SetInfra(v ClusterAutoInfrastructureComponentSettings)`

SetInfra sets Infra field to given value.

### HasInfra

`func (o *ClusterAutoInfrastructureUpgradeSettings) HasInfra() bool`

HasInfra returns a boolean if a field has been set.

### GetApps

`func (o *ClusterAutoInfrastructureUpgradeSettings) GetApps() ClusterAutoInfrastructureComponentSettings`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ClusterAutoInfrastructureUpgradeSettings) GetAppsOk() (*ClusterAutoInfrastructureComponentSettings, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ClusterAutoInfrastructureUpgradeSettings) SetApps(v ClusterAutoInfrastructureComponentSettings)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ClusterAutoInfrastructureUpgradeSettings) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


