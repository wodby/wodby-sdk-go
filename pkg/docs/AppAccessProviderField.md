# AppAccessProviderField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | **string** |  | 
**Required** | **bool** |  | 
**Radio** | **bool** |  | 
**HostnameSuffix** | **bool** |  | 
**AffectsHostname** | **bool** |  | 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Options** | [**[]AppAccessProviderOption**](AppAccessProviderOption.md) |  | 

## Methods

### NewAppAccessProviderField

`func NewAppAccessProviderField(name string, label string, description string, required bool, radio bool, hostnameSuffix bool, affectsHostname bool, options []AppAccessProviderOption, ) *AppAccessProviderField`

NewAppAccessProviderField instantiates a new AppAccessProviderField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessProviderFieldWithDefaults

`func NewAppAccessProviderFieldWithDefaults() *AppAccessProviderField`

NewAppAccessProviderFieldWithDefaults instantiates a new AppAccessProviderField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppAccessProviderField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppAccessProviderField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppAccessProviderField) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *AppAccessProviderField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppAccessProviderField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppAccessProviderField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *AppAccessProviderField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppAccessProviderField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppAccessProviderField) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRequired

`func (o *AppAccessProviderField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AppAccessProviderField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AppAccessProviderField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetRadio

`func (o *AppAccessProviderField) GetRadio() bool`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *AppAccessProviderField) GetRadioOk() (*bool, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *AppAccessProviderField) SetRadio(v bool)`

SetRadio sets Radio field to given value.


### GetHostnameSuffix

`func (o *AppAccessProviderField) GetHostnameSuffix() bool`

GetHostnameSuffix returns the HostnameSuffix field if non-nil, zero value otherwise.

### GetHostnameSuffixOk

`func (o *AppAccessProviderField) GetHostnameSuffixOk() (*bool, bool)`

GetHostnameSuffixOk returns a tuple with the HostnameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameSuffix

`func (o *AppAccessProviderField) SetHostnameSuffix(v bool)`

SetHostnameSuffix sets HostnameSuffix field to given value.


### GetAffectsHostname

`func (o *AppAccessProviderField) GetAffectsHostname() bool`

GetAffectsHostname returns the AffectsHostname field if non-nil, zero value otherwise.

### GetAffectsHostnameOk

`func (o *AppAccessProviderField) GetAffectsHostnameOk() (*bool, bool)`

GetAffectsHostnameOk returns a tuple with the AffectsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectsHostname

`func (o *AppAccessProviderField) SetAffectsHostname(v bool)`

SetAffectsHostname sets AffectsHostname field to given value.


### GetDefaultValue

`func (o *AppAccessProviderField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AppAccessProviderField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AppAccessProviderField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AppAccessProviderField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *AppAccessProviderField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *AppAccessProviderField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptions

`func (o *AppAccessProviderField) GetOptions() []AppAccessProviderOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppAccessProviderField) GetOptionsOk() (*[]AppAccessProviderOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppAccessProviderField) SetOptions(v []AppAccessProviderOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


