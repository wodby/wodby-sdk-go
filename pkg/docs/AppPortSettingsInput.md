# AppPortSettingsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | [**[]AppPortSettingInput**](AppPortSettingInput.md) | The complete list of settings. A setting left out is removed. | 

## Methods

### NewAppPortSettingsInput

`func NewAppPortSettingsInput(settings []AppPortSettingInput, ) *AppPortSettingsInput`

NewAppPortSettingsInput instantiates a new AppPortSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPortSettingsInputWithDefaults

`func NewAppPortSettingsInputWithDefaults() *AppPortSettingsInput`

NewAppPortSettingsInputWithDefaults instantiates a new AppPortSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *AppPortSettingsInput) GetSettings() []AppPortSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AppPortSettingsInput) GetSettingsOk() (*[]AppPortSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AppPortSettingsInput) SetSettings(v []AppPortSettingInput)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


