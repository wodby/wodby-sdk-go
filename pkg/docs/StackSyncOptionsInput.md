# StackSyncOptionsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteStackHelmValues** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackEnvVars** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackTokens** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackAnnotations** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackServices** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackServicesConfiguration** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewStackSyncOptionsInput

`func NewStackSyncOptionsInput() *StackSyncOptionsInput`

NewStackSyncOptionsInput instantiates a new StackSyncOptionsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackSyncOptionsInputWithDefaults

`func NewStackSyncOptionsInputWithDefaults() *StackSyncOptionsInput`

NewStackSyncOptionsInputWithDefaults instantiates a new StackSyncOptionsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteStackHelmValues

`func (o *StackSyncOptionsInput) GetDeleteStackHelmValues() bool`

GetDeleteStackHelmValues returns the DeleteStackHelmValues field if non-nil, zero value otherwise.

### GetDeleteStackHelmValuesOk

`func (o *StackSyncOptionsInput) GetDeleteStackHelmValuesOk() (*bool, bool)`

GetDeleteStackHelmValuesOk returns a tuple with the DeleteStackHelmValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackHelmValues

`func (o *StackSyncOptionsInput) SetDeleteStackHelmValues(v bool)`

SetDeleteStackHelmValues sets DeleteStackHelmValues field to given value.

### HasDeleteStackHelmValues

`func (o *StackSyncOptionsInput) HasDeleteStackHelmValues() bool`

HasDeleteStackHelmValues returns a boolean if a field has been set.

### SetDeleteStackHelmValuesNil

`func (o *StackSyncOptionsInput) SetDeleteStackHelmValuesNil(b bool)`

 SetDeleteStackHelmValuesNil sets the value for DeleteStackHelmValues to be an explicit nil

### UnsetDeleteStackHelmValues
`func (o *StackSyncOptionsInput) UnsetDeleteStackHelmValues()`

UnsetDeleteStackHelmValues ensures that no value is present for DeleteStackHelmValues, not even an explicit nil
### GetDeleteStackEnvVars

`func (o *StackSyncOptionsInput) GetDeleteStackEnvVars() bool`

GetDeleteStackEnvVars returns the DeleteStackEnvVars field if non-nil, zero value otherwise.

### GetDeleteStackEnvVarsOk

`func (o *StackSyncOptionsInput) GetDeleteStackEnvVarsOk() (*bool, bool)`

GetDeleteStackEnvVarsOk returns a tuple with the DeleteStackEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackEnvVars

`func (o *StackSyncOptionsInput) SetDeleteStackEnvVars(v bool)`

SetDeleteStackEnvVars sets DeleteStackEnvVars field to given value.

### HasDeleteStackEnvVars

`func (o *StackSyncOptionsInput) HasDeleteStackEnvVars() bool`

HasDeleteStackEnvVars returns a boolean if a field has been set.

### SetDeleteStackEnvVarsNil

`func (o *StackSyncOptionsInput) SetDeleteStackEnvVarsNil(b bool)`

 SetDeleteStackEnvVarsNil sets the value for DeleteStackEnvVars to be an explicit nil

### UnsetDeleteStackEnvVars
`func (o *StackSyncOptionsInput) UnsetDeleteStackEnvVars()`

UnsetDeleteStackEnvVars ensures that no value is present for DeleteStackEnvVars, not even an explicit nil
### GetDeleteStackTokens

`func (o *StackSyncOptionsInput) GetDeleteStackTokens() bool`

GetDeleteStackTokens returns the DeleteStackTokens field if non-nil, zero value otherwise.

### GetDeleteStackTokensOk

`func (o *StackSyncOptionsInput) GetDeleteStackTokensOk() (*bool, bool)`

GetDeleteStackTokensOk returns a tuple with the DeleteStackTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackTokens

`func (o *StackSyncOptionsInput) SetDeleteStackTokens(v bool)`

SetDeleteStackTokens sets DeleteStackTokens field to given value.

### HasDeleteStackTokens

`func (o *StackSyncOptionsInput) HasDeleteStackTokens() bool`

HasDeleteStackTokens returns a boolean if a field has been set.

### SetDeleteStackTokensNil

`func (o *StackSyncOptionsInput) SetDeleteStackTokensNil(b bool)`

 SetDeleteStackTokensNil sets the value for DeleteStackTokens to be an explicit nil

### UnsetDeleteStackTokens
`func (o *StackSyncOptionsInput) UnsetDeleteStackTokens()`

UnsetDeleteStackTokens ensures that no value is present for DeleteStackTokens, not even an explicit nil
### GetDeleteStackAnnotations

`func (o *StackSyncOptionsInput) GetDeleteStackAnnotations() bool`

GetDeleteStackAnnotations returns the DeleteStackAnnotations field if non-nil, zero value otherwise.

### GetDeleteStackAnnotationsOk

`func (o *StackSyncOptionsInput) GetDeleteStackAnnotationsOk() (*bool, bool)`

GetDeleteStackAnnotationsOk returns a tuple with the DeleteStackAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackAnnotations

`func (o *StackSyncOptionsInput) SetDeleteStackAnnotations(v bool)`

SetDeleteStackAnnotations sets DeleteStackAnnotations field to given value.

### HasDeleteStackAnnotations

`func (o *StackSyncOptionsInput) HasDeleteStackAnnotations() bool`

HasDeleteStackAnnotations returns a boolean if a field has been set.

### SetDeleteStackAnnotationsNil

`func (o *StackSyncOptionsInput) SetDeleteStackAnnotationsNil(b bool)`

 SetDeleteStackAnnotationsNil sets the value for DeleteStackAnnotations to be an explicit nil

### UnsetDeleteStackAnnotations
`func (o *StackSyncOptionsInput) UnsetDeleteStackAnnotations()`

UnsetDeleteStackAnnotations ensures that no value is present for DeleteStackAnnotations, not even an explicit nil
### GetDeleteStackServices

`func (o *StackSyncOptionsInput) GetDeleteStackServices() bool`

GetDeleteStackServices returns the DeleteStackServices field if non-nil, zero value otherwise.

### GetDeleteStackServicesOk

`func (o *StackSyncOptionsInput) GetDeleteStackServicesOk() (*bool, bool)`

GetDeleteStackServicesOk returns a tuple with the DeleteStackServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackServices

`func (o *StackSyncOptionsInput) SetDeleteStackServices(v bool)`

SetDeleteStackServices sets DeleteStackServices field to given value.

### HasDeleteStackServices

`func (o *StackSyncOptionsInput) HasDeleteStackServices() bool`

HasDeleteStackServices returns a boolean if a field has been set.

### SetDeleteStackServicesNil

`func (o *StackSyncOptionsInput) SetDeleteStackServicesNil(b bool)`

 SetDeleteStackServicesNil sets the value for DeleteStackServices to be an explicit nil

### UnsetDeleteStackServices
`func (o *StackSyncOptionsInput) UnsetDeleteStackServices()`

UnsetDeleteStackServices ensures that no value is present for DeleteStackServices, not even an explicit nil
### GetDeleteStackServicesConfiguration

`func (o *StackSyncOptionsInput) GetDeleteStackServicesConfiguration() bool`

GetDeleteStackServicesConfiguration returns the DeleteStackServicesConfiguration field if non-nil, zero value otherwise.

### GetDeleteStackServicesConfigurationOk

`func (o *StackSyncOptionsInput) GetDeleteStackServicesConfigurationOk() (*bool, bool)`

GetDeleteStackServicesConfigurationOk returns a tuple with the DeleteStackServicesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackServicesConfiguration

`func (o *StackSyncOptionsInput) SetDeleteStackServicesConfiguration(v bool)`

SetDeleteStackServicesConfiguration sets DeleteStackServicesConfiguration field to given value.

### HasDeleteStackServicesConfiguration

`func (o *StackSyncOptionsInput) HasDeleteStackServicesConfiguration() bool`

HasDeleteStackServicesConfiguration returns a boolean if a field has been set.

### SetDeleteStackServicesConfigurationNil

`func (o *StackSyncOptionsInput) SetDeleteStackServicesConfigurationNil(b bool)`

 SetDeleteStackServicesConfigurationNil sets the value for DeleteStackServicesConfiguration to be an explicit nil

### UnsetDeleteStackServicesConfiguration
`func (o *StackSyncOptionsInput) UnsetDeleteStackServicesConfiguration()`

UnsetDeleteStackServicesConfiguration ensures that no value is present for DeleteStackServicesConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


