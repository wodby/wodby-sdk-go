# AppEnvironmentAutoStackUpgradeSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**UpgradeSettings** | Pointer to [**AppEnvironmentStackUpgradeSettingsInput**](AppEnvironmentStackUpgradeSettingsInput.md) |  | [optional] 
**TimeWindow** | Pointer to [**AutomationTimeWindowInput**](AutomationTimeWindowInput.md) |  | [optional] 

## Methods

### NewAppEnvironmentAutoStackUpgradeSettingsInput

`func NewAppEnvironmentAutoStackUpgradeSettingsInput(enabled bool, ) *AppEnvironmentAutoStackUpgradeSettingsInput`

NewAppEnvironmentAutoStackUpgradeSettingsInput instantiates a new AppEnvironmentAutoStackUpgradeSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentAutoStackUpgradeSettingsInputWithDefaults

`func NewAppEnvironmentAutoStackUpgradeSettingsInputWithDefaults() *AppEnvironmentAutoStackUpgradeSettingsInput`

NewAppEnvironmentAutoStackUpgradeSettingsInputWithDefaults instantiates a new AppEnvironmentAutoStackUpgradeSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUpgradeSettings

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) GetUpgradeSettings() AppEnvironmentStackUpgradeSettingsInput`

GetUpgradeSettings returns the UpgradeSettings field if non-nil, zero value otherwise.

### GetUpgradeSettingsOk

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) GetUpgradeSettingsOk() (*AppEnvironmentStackUpgradeSettingsInput, bool)`

GetUpgradeSettingsOk returns a tuple with the UpgradeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSettings

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) SetUpgradeSettings(v AppEnvironmentStackUpgradeSettingsInput)`

SetUpgradeSettings sets UpgradeSettings field to given value.

### HasUpgradeSettings

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) HasUpgradeSettings() bool`

HasUpgradeSettings returns a boolean if a field has been set.

### GetTimeWindow

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) GetTimeWindow() AutomationTimeWindowInput`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) GetTimeWindowOk() (*AutomationTimeWindowInput, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) SetTimeWindow(v AutomationTimeWindowInput)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *AppEnvironmentAutoStackUpgradeSettingsInput) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


