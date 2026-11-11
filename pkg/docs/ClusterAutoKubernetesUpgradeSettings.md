# ClusterAutoKubernetesUpgradeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**AllowPatch** | **bool** |  | 
**AllowMinor** | **bool** |  | 
**MinorOnlyAfterEndOfLife** | **bool** | Holds an allowed minor upgrade until the current minor reaches upstream end of life. | 
**TimeWindow** | Pointer to [**AutomationTimeWindow**](AutomationTimeWindow.md) | When an upgrade may start. Absent means any time. | [optional] 
**IncludeComponentUpdates** | **bool** | Also applies the cluster&#39;s component updates when there is no version upgrade to start, under the same policy and window. On by default. | 

## Methods

### NewClusterAutoKubernetesUpgradeSettings

`func NewClusterAutoKubernetesUpgradeSettings(enabled bool, allowPatch bool, allowMinor bool, minorOnlyAfterEndOfLife bool, includeComponentUpdates bool, ) *ClusterAutoKubernetesUpgradeSettings`

NewClusterAutoKubernetesUpgradeSettings instantiates a new ClusterAutoKubernetesUpgradeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoKubernetesUpgradeSettingsWithDefaults

`func NewClusterAutoKubernetesUpgradeSettingsWithDefaults() *ClusterAutoKubernetesUpgradeSettings`

NewClusterAutoKubernetesUpgradeSettingsWithDefaults instantiates a new ClusterAutoKubernetesUpgradeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAutoKubernetesUpgradeSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAutoKubernetesUpgradeSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAutoKubernetesUpgradeSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowPatch

`func (o *ClusterAutoKubernetesUpgradeSettings) GetAllowPatch() bool`

GetAllowPatch returns the AllowPatch field if non-nil, zero value otherwise.

### GetAllowPatchOk

`func (o *ClusterAutoKubernetesUpgradeSettings) GetAllowPatchOk() (*bool, bool)`

GetAllowPatchOk returns a tuple with the AllowPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPatch

`func (o *ClusterAutoKubernetesUpgradeSettings) SetAllowPatch(v bool)`

SetAllowPatch sets AllowPatch field to given value.


### GetAllowMinor

`func (o *ClusterAutoKubernetesUpgradeSettings) GetAllowMinor() bool`

GetAllowMinor returns the AllowMinor field if non-nil, zero value otherwise.

### GetAllowMinorOk

`func (o *ClusterAutoKubernetesUpgradeSettings) GetAllowMinorOk() (*bool, bool)`

GetAllowMinorOk returns a tuple with the AllowMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMinor

`func (o *ClusterAutoKubernetesUpgradeSettings) SetAllowMinor(v bool)`

SetAllowMinor sets AllowMinor field to given value.


### GetMinorOnlyAfterEndOfLife

`func (o *ClusterAutoKubernetesUpgradeSettings) GetMinorOnlyAfterEndOfLife() bool`

GetMinorOnlyAfterEndOfLife returns the MinorOnlyAfterEndOfLife field if non-nil, zero value otherwise.

### GetMinorOnlyAfterEndOfLifeOk

`func (o *ClusterAutoKubernetesUpgradeSettings) GetMinorOnlyAfterEndOfLifeOk() (*bool, bool)`

GetMinorOnlyAfterEndOfLifeOk returns a tuple with the MinorOnlyAfterEndOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorOnlyAfterEndOfLife

`func (o *ClusterAutoKubernetesUpgradeSettings) SetMinorOnlyAfterEndOfLife(v bool)`

SetMinorOnlyAfterEndOfLife sets MinorOnlyAfterEndOfLife field to given value.


### GetTimeWindow

`func (o *ClusterAutoKubernetesUpgradeSettings) GetTimeWindow() AutomationTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ClusterAutoKubernetesUpgradeSettings) GetTimeWindowOk() (*AutomationTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ClusterAutoKubernetesUpgradeSettings) SetTimeWindow(v AutomationTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *ClusterAutoKubernetesUpgradeSettings) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetIncludeComponentUpdates

`func (o *ClusterAutoKubernetesUpgradeSettings) GetIncludeComponentUpdates() bool`

GetIncludeComponentUpdates returns the IncludeComponentUpdates field if non-nil, zero value otherwise.

### GetIncludeComponentUpdatesOk

`func (o *ClusterAutoKubernetesUpgradeSettings) GetIncludeComponentUpdatesOk() (*bool, bool)`

GetIncludeComponentUpdatesOk returns a tuple with the IncludeComponentUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeComponentUpdates

`func (o *ClusterAutoKubernetesUpgradeSettings) SetIncludeComponentUpdates(v bool)`

SetIncludeComponentUpdates sets IncludeComponentUpdates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


