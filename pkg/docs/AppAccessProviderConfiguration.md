# AppAccessProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**EndpointHostMode** | **string** |  | 
**Fields** | [**[]AppAccessProviderField**](AppAccessProviderField.md) |  | 

## Methods

### NewAppAccessProviderConfiguration

`func NewAppAccessProviderConfiguration(mode string, endpointHostMode string, fields []AppAccessProviderField, ) *AppAccessProviderConfiguration`

NewAppAccessProviderConfiguration instantiates a new AppAccessProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessProviderConfigurationWithDefaults

`func NewAppAccessProviderConfigurationWithDefaults() *AppAccessProviderConfiguration`

NewAppAccessProviderConfigurationWithDefaults instantiates a new AppAccessProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AppAccessProviderConfiguration) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AppAccessProviderConfiguration) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AppAccessProviderConfiguration) SetMode(v string)`

SetMode sets Mode field to given value.


### GetEndpointHostMode

`func (o *AppAccessProviderConfiguration) GetEndpointHostMode() string`

GetEndpointHostMode returns the EndpointHostMode field if non-nil, zero value otherwise.

### GetEndpointHostModeOk

`func (o *AppAccessProviderConfiguration) GetEndpointHostModeOk() (*string, bool)`

GetEndpointHostModeOk returns a tuple with the EndpointHostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointHostMode

`func (o *AppAccessProviderConfiguration) SetEndpointHostMode(v string)`

SetEndpointHostMode sets EndpointHostMode field to given value.


### GetFields

`func (o *AppAccessProviderConfiguration) GetFields() []AppAccessProviderField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AppAccessProviderConfiguration) GetFieldsOk() (*[]AppAccessProviderField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AppAccessProviderConfiguration) SetFields(v []AppAccessProviderField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


