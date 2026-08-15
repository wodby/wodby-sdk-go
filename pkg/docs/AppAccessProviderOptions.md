# AppAccessProviderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Modes** | **[]string** |  | 
**Scopes** | **[]string** |  | 
**EndpointHostMode** | **string** |  | 
**Fields** | [**[]AppAccessProviderField**](AppAccessProviderField.md) |  | 
**Configurations** | [**[]AppAccessProviderConfiguration**](AppAccessProviderConfiguration.md) |  | 

## Methods

### NewAppAccessProviderOptions

`func NewAppAccessProviderOptions(provider string, modes []string, scopes []string, endpointHostMode string, fields []AppAccessProviderField, configurations []AppAccessProviderConfiguration, ) *AppAccessProviderOptions`

NewAppAccessProviderOptions instantiates a new AppAccessProviderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessProviderOptionsWithDefaults

`func NewAppAccessProviderOptionsWithDefaults() *AppAccessProviderOptions`

NewAppAccessProviderOptionsWithDefaults instantiates a new AppAccessProviderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AppAccessProviderOptions) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AppAccessProviderOptions) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AppAccessProviderOptions) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModes

`func (o *AppAccessProviderOptions) GetModes() []string`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *AppAccessProviderOptions) GetModesOk() (*[]string, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *AppAccessProviderOptions) SetModes(v []string)`

SetModes sets Modes field to given value.


### GetScopes

`func (o *AppAccessProviderOptions) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppAccessProviderOptions) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppAccessProviderOptions) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetEndpointHostMode

`func (o *AppAccessProviderOptions) GetEndpointHostMode() string`

GetEndpointHostMode returns the EndpointHostMode field if non-nil, zero value otherwise.

### GetEndpointHostModeOk

`func (o *AppAccessProviderOptions) GetEndpointHostModeOk() (*string, bool)`

GetEndpointHostModeOk returns a tuple with the EndpointHostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointHostMode

`func (o *AppAccessProviderOptions) SetEndpointHostMode(v string)`

SetEndpointHostMode sets EndpointHostMode field to given value.


### GetFields

`func (o *AppAccessProviderOptions) GetFields() []AppAccessProviderField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AppAccessProviderOptions) GetFieldsOk() (*[]AppAccessProviderField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AppAccessProviderOptions) SetFields(v []AppAccessProviderField)`

SetFields sets Fields field to given value.


### GetConfigurations

`func (o *AppAccessProviderOptions) GetConfigurations() []AppAccessProviderConfiguration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *AppAccessProviderOptions) GetConfigurationsOk() (*[]AppAccessProviderConfiguration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *AppAccessProviderOptions) SetConfigurations(v []AppAccessProviderConfiguration)`

SetConfigurations sets Configurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


