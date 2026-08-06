# ClusterAutoInfrastructureComponentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**VersionPolicy** | Pointer to [**ClusterAutoUpgradeVersionPolicy**](ClusterAutoUpgradeVersionPolicy.md) |  | [optional] 
**TimeWindow** | Pointer to [**AutomationTimeWindow**](AutomationTimeWindow.md) |  | [optional] 

## Methods

### NewClusterAutoInfrastructureComponentSettings

`func NewClusterAutoInfrastructureComponentSettings(enabled bool, ) *ClusterAutoInfrastructureComponentSettings`

NewClusterAutoInfrastructureComponentSettings instantiates a new ClusterAutoInfrastructureComponentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoInfrastructureComponentSettingsWithDefaults

`func NewClusterAutoInfrastructureComponentSettingsWithDefaults() *ClusterAutoInfrastructureComponentSettings`

NewClusterAutoInfrastructureComponentSettingsWithDefaults instantiates a new ClusterAutoInfrastructureComponentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAutoInfrastructureComponentSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAutoInfrastructureComponentSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAutoInfrastructureComponentSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVersionPolicy

`func (o *ClusterAutoInfrastructureComponentSettings) GetVersionPolicy() ClusterAutoUpgradeVersionPolicy`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *ClusterAutoInfrastructureComponentSettings) GetVersionPolicyOk() (*ClusterAutoUpgradeVersionPolicy, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *ClusterAutoInfrastructureComponentSettings) SetVersionPolicy(v ClusterAutoUpgradeVersionPolicy)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *ClusterAutoInfrastructureComponentSettings) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.

### GetTimeWindow

`func (o *ClusterAutoInfrastructureComponentSettings) GetTimeWindow() AutomationTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ClusterAutoInfrastructureComponentSettings) GetTimeWindowOk() (*AutomationTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ClusterAutoInfrastructureComponentSettings) SetTimeWindow(v AutomationTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *ClusterAutoInfrastructureComponentSettings) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


