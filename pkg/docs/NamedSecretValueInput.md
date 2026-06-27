# NamedSecretValueInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Secret** | **bool** |  | 

## Methods

### NewNamedSecretValueInput

`func NewNamedSecretValueInput(name string, value string, secret bool, ) *NamedSecretValueInput`

NewNamedSecretValueInput instantiates a new NamedSecretValueInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedSecretValueInputWithDefaults

`func NewNamedSecretValueInputWithDefaults() *NamedSecretValueInput`

NewNamedSecretValueInputWithDefaults instantiates a new NamedSecretValueInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NamedSecretValueInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamedSecretValueInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamedSecretValueInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NamedSecretValueInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NamedSecretValueInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NamedSecretValueInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *NamedSecretValueInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NamedSecretValueInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NamedSecretValueInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


