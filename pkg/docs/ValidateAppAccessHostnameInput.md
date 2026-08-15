# ValidateAppAccessHostnameInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**[]AppAccessSettingInput**](AppAccessSettingInput.md) |  | [optional] 
**Host** | **string** |  | 

## Methods

### NewValidateAppAccessHostnameInput

`func NewValidateAppAccessHostnameInput(host string, ) *ValidateAppAccessHostnameInput`

NewValidateAppAccessHostnameInput instantiates a new ValidateAppAccessHostnameInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAppAccessHostnameInputWithDefaults

`func NewValidateAppAccessHostnameInputWithDefaults() *ValidateAppAccessHostnameInput`

NewValidateAppAccessHostnameInputWithDefaults instantiates a new ValidateAppAccessHostnameInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *ValidateAppAccessHostnameInput) GetSettings() []AppAccessSettingInput`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ValidateAppAccessHostnameInput) GetSettingsOk() (*[]AppAccessSettingInput, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ValidateAppAccessHostnameInput) SetSettings(v []AppAccessSettingInput)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ValidateAppAccessHostnameInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetHost

`func (o *ValidateAppAccessHostnameInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ValidateAppAccessHostnameInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ValidateAppAccessHostnameInput) SetHost(v string)`

SetHost sets Host field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


