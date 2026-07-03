# AppInstanceAutoStackUpgradeSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**UpgradeSettings** | Pointer to [**AppInstanceStackUpgradeSettingsInput**](AppInstanceStackUpgradeSettingsInput.md) |  | [optional] 

## Methods

### NewAppInstanceAutoStackUpgradeSettingsInput

`func NewAppInstanceAutoStackUpgradeSettingsInput(enabled bool, ) *AppInstanceAutoStackUpgradeSettingsInput`

NewAppInstanceAutoStackUpgradeSettingsInput instantiates a new AppInstanceAutoStackUpgradeSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceAutoStackUpgradeSettingsInputWithDefaults

`func NewAppInstanceAutoStackUpgradeSettingsInputWithDefaults() *AppInstanceAutoStackUpgradeSettingsInput`

NewAppInstanceAutoStackUpgradeSettingsInputWithDefaults instantiates a new AppInstanceAutoStackUpgradeSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppInstanceAutoStackUpgradeSettingsInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppInstanceAutoStackUpgradeSettingsInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppInstanceAutoStackUpgradeSettingsInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUpgradeSettings

`func (o *AppInstanceAutoStackUpgradeSettingsInput) GetUpgradeSettings() AppInstanceStackUpgradeSettingsInput`

GetUpgradeSettings returns the UpgradeSettings field if non-nil, zero value otherwise.

### GetUpgradeSettingsOk

`func (o *AppInstanceAutoStackUpgradeSettingsInput) GetUpgradeSettingsOk() (*AppInstanceStackUpgradeSettingsInput, bool)`

GetUpgradeSettingsOk returns a tuple with the UpgradeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSettings

`func (o *AppInstanceAutoStackUpgradeSettingsInput) SetUpgradeSettings(v AppInstanceStackUpgradeSettingsInput)`

SetUpgradeSettings sets UpgradeSettings field to given value.

### HasUpgradeSettings

`func (o *AppInstanceAutoStackUpgradeSettingsInput) HasUpgradeSettings() bool`

HasUpgradeSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


