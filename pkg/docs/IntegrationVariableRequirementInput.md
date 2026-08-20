# IntegrationVariableRequirementInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Secret** | Pointer to **bool** |  | [optional] [default to false]
**Optional** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIntegrationVariableRequirementInput

`func NewIntegrationVariableRequirementInput(name string, ) *IntegrationVariableRequirementInput`

NewIntegrationVariableRequirementInput instantiates a new IntegrationVariableRequirementInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationVariableRequirementInputWithDefaults

`func NewIntegrationVariableRequirementInputWithDefaults() *IntegrationVariableRequirementInput`

NewIntegrationVariableRequirementInputWithDefaults instantiates a new IntegrationVariableRequirementInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegrationVariableRequirementInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationVariableRequirementInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationVariableRequirementInput) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *IntegrationVariableRequirementInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IntegrationVariableRequirementInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IntegrationVariableRequirementInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IntegrationVariableRequirementInput) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetOptional

`func (o *IntegrationVariableRequirementInput) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IntegrationVariableRequirementInput) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IntegrationVariableRequirementInput) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IntegrationVariableRequirementInput) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


