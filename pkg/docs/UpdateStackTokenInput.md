# UpdateStackTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**Secret** | **bool** |  | 
**Regex** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateStackTokenInput

`func NewUpdateStackTokenInput(secret bool, ) *UpdateStackTokenInput`

NewUpdateStackTokenInput instantiates a new UpdateStackTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStackTokenInputWithDefaults

`func NewUpdateStackTokenInputWithDefaults() *UpdateStackTokenInput`

NewUpdateStackTokenInputWithDefaults instantiates a new UpdateStackTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateStackTokenInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateStackTokenInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateStackTokenInput) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateStackTokenInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateStackTokenInput) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateStackTokenInput) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *UpdateStackTokenInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateStackTokenInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateStackTokenInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetRegex

`func (o *UpdateStackTokenInput) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *UpdateStackTokenInput) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *UpdateStackTokenInput) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *UpdateStackTokenInput) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *UpdateStackTokenInput) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *UpdateStackTokenInput) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


