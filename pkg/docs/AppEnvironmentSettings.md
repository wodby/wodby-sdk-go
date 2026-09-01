# AppEnvironmentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoStackUpgrade** | Pointer to [**AppEnvironmentAutoStackUpgradeSettings**](AppEnvironmentAutoStackUpgradeSettings.md) |  | [optional] 

## Methods

### NewAppEnvironmentSettings

`func NewAppEnvironmentSettings() *AppEnvironmentSettings`

NewAppEnvironmentSettings instantiates a new AppEnvironmentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEnvironmentSettingsWithDefaults

`func NewAppEnvironmentSettingsWithDefaults() *AppEnvironmentSettings`

NewAppEnvironmentSettingsWithDefaults instantiates a new AppEnvironmentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoStackUpgrade

`func (o *AppEnvironmentSettings) GetAutoStackUpgrade() AppEnvironmentAutoStackUpgradeSettings`

GetAutoStackUpgrade returns the AutoStackUpgrade field if non-nil, zero value otherwise.

### GetAutoStackUpgradeOk

`func (o *AppEnvironmentSettings) GetAutoStackUpgradeOk() (*AppEnvironmentAutoStackUpgradeSettings, bool)`

GetAutoStackUpgradeOk returns a tuple with the AutoStackUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStackUpgrade

`func (o *AppEnvironmentSettings) SetAutoStackUpgrade(v AppEnvironmentAutoStackUpgradeSettings)`

SetAutoStackUpgrade sets AutoStackUpgrade field to given value.

### HasAutoStackUpgrade

`func (o *AppEnvironmentSettings) HasAutoStackUpgrade() bool`

HasAutoStackUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


