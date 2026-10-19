# NewStackTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** | Leave it out to generate a value from regex. | [optional] 
**Secret** | **bool** |  | 
**Regex** | Pointer to **NullableString** | Pattern used to generate the value for each app environment. | [optional] 
**EnvType** | Pointer to **NullableString** | Limit the value to app environments of this type. | [optional] 

## Methods

### NewNewStackTokenInput

`func NewNewStackTokenInput(name string, secret bool, ) *NewStackTokenInput`

NewNewStackTokenInput instantiates a new NewStackTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackTokenInputWithDefaults

`func NewNewStackTokenInputWithDefaults() *NewStackTokenInput`

NewNewStackTokenInputWithDefaults instantiates a new NewStackTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewStackTokenInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackTokenInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackTokenInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewStackTokenInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewStackTokenInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewStackTokenInput) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NewStackTokenInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *NewStackTokenInput) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *NewStackTokenInput) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *NewStackTokenInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NewStackTokenInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NewStackTokenInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetRegex

`func (o *NewStackTokenInput) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *NewStackTokenInput) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *NewStackTokenInput) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *NewStackTokenInput) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *NewStackTokenInput) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *NewStackTokenInput) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetEnvType

`func (o *NewStackTokenInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *NewStackTokenInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *NewStackTokenInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *NewStackTokenInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *NewStackTokenInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *NewStackTokenInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


