# ClusterAutoKubernetesUpgradeSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**AllowPatch** | Pointer to **NullableBool** |  | [optional] 
**AllowMinor** | Pointer to **NullableBool** |  | [optional] 
**MinorOnlyAfterEndOfLife** | Pointer to **NullableBool** |  | [optional] 
**IncludeComponentUpdates** | Pointer to **NullableBool** |  | [optional] 
**TimeWindow** | Pointer to [**AutomationTimeWindowInput**](AutomationTimeWindowInput.md) | A window with enabled false lets upgrades start at any time. | [optional] 

## Methods

### NewClusterAutoKubernetesUpgradeSettingsInput

`func NewClusterAutoKubernetesUpgradeSettingsInput() *ClusterAutoKubernetesUpgradeSettingsInput`

NewClusterAutoKubernetesUpgradeSettingsInput instantiates a new ClusterAutoKubernetesUpgradeSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoKubernetesUpgradeSettingsInputWithDefaults

`func NewClusterAutoKubernetesUpgradeSettingsInputWithDefaults() *ClusterAutoKubernetesUpgradeSettingsInput`

NewClusterAutoKubernetesUpgradeSettingsInputWithDefaults instantiates a new ClusterAutoKubernetesUpgradeSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ClusterAutoKubernetesUpgradeSettingsInput) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetAllowPatch

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetAllowPatch() bool`

GetAllowPatch returns the AllowPatch field if non-nil, zero value otherwise.

### GetAllowPatchOk

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetAllowPatchOk() (*bool, bool)`

GetAllowPatchOk returns a tuple with the AllowPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPatch

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetAllowPatch(v bool)`

SetAllowPatch sets AllowPatch field to given value.

### HasAllowPatch

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) HasAllowPatch() bool`

HasAllowPatch returns a boolean if a field has been set.

### SetAllowPatchNil

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetAllowPatchNil(b bool)`

 SetAllowPatchNil sets the value for AllowPatch to be an explicit nil

### UnsetAllowPatch
`func (o *ClusterAutoKubernetesUpgradeSettingsInput) UnsetAllowPatch()`

UnsetAllowPatch ensures that no value is present for AllowPatch, not even an explicit nil
### GetAllowMinor

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetAllowMinor() bool`

GetAllowMinor returns the AllowMinor field if non-nil, zero value otherwise.

### GetAllowMinorOk

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetAllowMinorOk() (*bool, bool)`

GetAllowMinorOk returns a tuple with the AllowMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMinor

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetAllowMinor(v bool)`

SetAllowMinor sets AllowMinor field to given value.

### HasAllowMinor

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) HasAllowMinor() bool`

HasAllowMinor returns a boolean if a field has been set.

### SetAllowMinorNil

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetAllowMinorNil(b bool)`

 SetAllowMinorNil sets the value for AllowMinor to be an explicit nil

### UnsetAllowMinor
`func (o *ClusterAutoKubernetesUpgradeSettingsInput) UnsetAllowMinor()`

UnsetAllowMinor ensures that no value is present for AllowMinor, not even an explicit nil
### GetMinorOnlyAfterEndOfLife

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetMinorOnlyAfterEndOfLife() bool`

GetMinorOnlyAfterEndOfLife returns the MinorOnlyAfterEndOfLife field if non-nil, zero value otherwise.

### GetMinorOnlyAfterEndOfLifeOk

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetMinorOnlyAfterEndOfLifeOk() (*bool, bool)`

GetMinorOnlyAfterEndOfLifeOk returns a tuple with the MinorOnlyAfterEndOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorOnlyAfterEndOfLife

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetMinorOnlyAfterEndOfLife(v bool)`

SetMinorOnlyAfterEndOfLife sets MinorOnlyAfterEndOfLife field to given value.

### HasMinorOnlyAfterEndOfLife

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) HasMinorOnlyAfterEndOfLife() bool`

HasMinorOnlyAfterEndOfLife returns a boolean if a field has been set.

### SetMinorOnlyAfterEndOfLifeNil

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetMinorOnlyAfterEndOfLifeNil(b bool)`

 SetMinorOnlyAfterEndOfLifeNil sets the value for MinorOnlyAfterEndOfLife to be an explicit nil

### UnsetMinorOnlyAfterEndOfLife
`func (o *ClusterAutoKubernetesUpgradeSettingsInput) UnsetMinorOnlyAfterEndOfLife()`

UnsetMinorOnlyAfterEndOfLife ensures that no value is present for MinorOnlyAfterEndOfLife, not even an explicit nil
### GetIncludeComponentUpdates

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetIncludeComponentUpdates() bool`

GetIncludeComponentUpdates returns the IncludeComponentUpdates field if non-nil, zero value otherwise.

### GetIncludeComponentUpdatesOk

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetIncludeComponentUpdatesOk() (*bool, bool)`

GetIncludeComponentUpdatesOk returns a tuple with the IncludeComponentUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeComponentUpdates

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetIncludeComponentUpdates(v bool)`

SetIncludeComponentUpdates sets IncludeComponentUpdates field to given value.

### HasIncludeComponentUpdates

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) HasIncludeComponentUpdates() bool`

HasIncludeComponentUpdates returns a boolean if a field has been set.

### SetIncludeComponentUpdatesNil

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetIncludeComponentUpdatesNil(b bool)`

 SetIncludeComponentUpdatesNil sets the value for IncludeComponentUpdates to be an explicit nil

### UnsetIncludeComponentUpdates
`func (o *ClusterAutoKubernetesUpgradeSettingsInput) UnsetIncludeComponentUpdates()`

UnsetIncludeComponentUpdates ensures that no value is present for IncludeComponentUpdates, not even an explicit nil
### GetTimeWindow

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetTimeWindow() AutomationTimeWindowInput`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) GetTimeWindowOk() (*AutomationTimeWindowInput, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) SetTimeWindow(v AutomationTimeWindowInput)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *ClusterAutoKubernetesUpgradeSettingsInput) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


