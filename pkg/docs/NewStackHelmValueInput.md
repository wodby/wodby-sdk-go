# NewStackHelmValueInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Secret** | **bool** |  | 
**EnvType** | Pointer to **NullableString** | Limit the value to app environments of this type. | [optional] 

## Methods

### NewNewStackHelmValueInput

`func NewNewStackHelmValueInput(name string, value string, secret bool, ) *NewStackHelmValueInput`

NewNewStackHelmValueInput instantiates a new NewStackHelmValueInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackHelmValueInputWithDefaults

`func NewNewStackHelmValueInputWithDefaults() *NewStackHelmValueInput`

NewNewStackHelmValueInputWithDefaults instantiates a new NewStackHelmValueInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewStackHelmValueInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackHelmValueInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackHelmValueInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewStackHelmValueInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewStackHelmValueInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewStackHelmValueInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *NewStackHelmValueInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NewStackHelmValueInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NewStackHelmValueInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetEnvType

`func (o *NewStackHelmValueInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *NewStackHelmValueInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *NewStackHelmValueInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *NewStackHelmValueInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *NewStackHelmValueInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *NewStackHelmValueInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


