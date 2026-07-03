# AppInstanceAutoStackUpgradeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**UpgradeSettings** | Pointer to [**AppInstanceStackUpgradeSettings**](AppInstanceStackUpgradeSettings.md) |  | [optional] 

## Methods

### NewAppInstanceAutoStackUpgradeSettings

`func NewAppInstanceAutoStackUpgradeSettings(enabled bool, ) *AppInstanceAutoStackUpgradeSettings`

NewAppInstanceAutoStackUpgradeSettings instantiates a new AppInstanceAutoStackUpgradeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceAutoStackUpgradeSettingsWithDefaults

`func NewAppInstanceAutoStackUpgradeSettingsWithDefaults() *AppInstanceAutoStackUpgradeSettings`

NewAppInstanceAutoStackUpgradeSettingsWithDefaults instantiates a new AppInstanceAutoStackUpgradeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppInstanceAutoStackUpgradeSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppInstanceAutoStackUpgradeSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppInstanceAutoStackUpgradeSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUpgradeSettings

`func (o *AppInstanceAutoStackUpgradeSettings) GetUpgradeSettings() AppInstanceStackUpgradeSettings`

GetUpgradeSettings returns the UpgradeSettings field if non-nil, zero value otherwise.

### GetUpgradeSettingsOk

`func (o *AppInstanceAutoStackUpgradeSettings) GetUpgradeSettingsOk() (*AppInstanceStackUpgradeSettings, bool)`

GetUpgradeSettingsOk returns a tuple with the UpgradeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSettings

`func (o *AppInstanceAutoStackUpgradeSettings) SetUpgradeSettings(v AppInstanceStackUpgradeSettings)`

SetUpgradeSettings sets UpgradeSettings field to given value.

### HasUpgradeSettings

`func (o *AppInstanceAutoStackUpgradeSettings) HasUpgradeSettings() bool`

HasUpgradeSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


