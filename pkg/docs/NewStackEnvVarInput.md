# NewStackEnvVarInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Secret** | **bool** |  | 
**EnvType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewStackEnvVarInput

`func NewNewStackEnvVarInput(name string, value string, secret bool, ) *NewStackEnvVarInput`

NewNewStackEnvVarInput instantiates a new NewStackEnvVarInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackEnvVarInputWithDefaults

`func NewNewStackEnvVarInputWithDefaults() *NewStackEnvVarInput`

NewNewStackEnvVarInputWithDefaults instantiates a new NewStackEnvVarInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewStackEnvVarInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackEnvVarInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackEnvVarInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewStackEnvVarInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewStackEnvVarInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewStackEnvVarInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *NewStackEnvVarInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NewStackEnvVarInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NewStackEnvVarInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetEnvType

`func (o *NewStackEnvVarInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *NewStackEnvVarInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *NewStackEnvVarInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *NewStackEnvVarInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *NewStackEnvVarInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *NewStackEnvVarInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


