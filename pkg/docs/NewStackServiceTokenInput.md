# NewStackServiceTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Secret** | **bool** |  | 
**Regex** | Pointer to **NullableString** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewStackServiceTokenInput

`func NewNewStackServiceTokenInput(name string, secret bool, ) *NewStackServiceTokenInput`

NewNewStackServiceTokenInput instantiates a new NewStackServiceTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackServiceTokenInputWithDefaults

`func NewNewStackServiceTokenInputWithDefaults() *NewStackServiceTokenInput`

NewNewStackServiceTokenInputWithDefaults instantiates a new NewStackServiceTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewStackServiceTokenInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackServiceTokenInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackServiceTokenInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewStackServiceTokenInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewStackServiceTokenInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewStackServiceTokenInput) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NewStackServiceTokenInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *NewStackServiceTokenInput) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *NewStackServiceTokenInput) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *NewStackServiceTokenInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NewStackServiceTokenInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NewStackServiceTokenInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetRegex

`func (o *NewStackServiceTokenInput) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *NewStackServiceTokenInput) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *NewStackServiceTokenInput) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *NewStackServiceTokenInput) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *NewStackServiceTokenInput) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *NewStackServiceTokenInput) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetEnvType

`func (o *NewStackServiceTokenInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *NewStackServiceTokenInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *NewStackServiceTokenInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *NewStackServiceTokenInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *NewStackServiceTokenInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *NewStackServiceTokenInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


