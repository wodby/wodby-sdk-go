# AppServiceCustomLinkValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Secret** | **bool** |  | 
**HasValue** | **bool** |  | 
**Value** | **NullableString** | The value; null for a secret value. | 

## Methods

### NewAppServiceCustomLinkValue

`func NewAppServiceCustomLinkValue(name string, secret bool, hasValue bool, value NullableString, ) *AppServiceCustomLinkValue`

NewAppServiceCustomLinkValue instantiates a new AppServiceCustomLinkValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceCustomLinkValueWithDefaults

`func NewAppServiceCustomLinkValueWithDefaults() *AppServiceCustomLinkValue`

NewAppServiceCustomLinkValueWithDefaults instantiates a new AppServiceCustomLinkValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppServiceCustomLinkValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceCustomLinkValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceCustomLinkValue) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *AppServiceCustomLinkValue) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AppServiceCustomLinkValue) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AppServiceCustomLinkValue) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetHasValue

`func (o *AppServiceCustomLinkValue) GetHasValue() bool`

GetHasValue returns the HasValue field if non-nil, zero value otherwise.

### GetHasValueOk

`func (o *AppServiceCustomLinkValue) GetHasValueOk() (*bool, bool)`

GetHasValueOk returns a tuple with the HasValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValue

`func (o *AppServiceCustomLinkValue) SetHasValue(v bool)`

SetHasValue sets HasValue field to given value.


### GetValue

`func (o *AppServiceCustomLinkValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceCustomLinkValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceCustomLinkValue) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AppServiceCustomLinkValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AppServiceCustomLinkValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


