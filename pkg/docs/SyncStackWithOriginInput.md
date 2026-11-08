# SyncStackWithOriginInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteStackHelmValues** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackEnvVars** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackTokens** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackAnnotations** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackServices** | Pointer to **NullableBool** |  | [optional] 
**DeleteStackServicesConfiguration** | Pointer to **NullableBool** |  | [optional] 
**UpdateStackServices** | Pointer to **NullableBool** | Let the sync update existing stack services to the service version the origin uses when the origin&#39;s changes need it. Without it, such a sync fails with an error that lists those stack services. | [optional] 

## Methods

### NewSyncStackWithOriginInput

`func NewSyncStackWithOriginInput() *SyncStackWithOriginInput`

NewSyncStackWithOriginInput instantiates a new SyncStackWithOriginInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStackWithOriginInputWithDefaults

`func NewSyncStackWithOriginInputWithDefaults() *SyncStackWithOriginInput`

NewSyncStackWithOriginInputWithDefaults instantiates a new SyncStackWithOriginInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteStackHelmValues

`func (o *SyncStackWithOriginInput) GetDeleteStackHelmValues() bool`

GetDeleteStackHelmValues returns the DeleteStackHelmValues field if non-nil, zero value otherwise.

### GetDeleteStackHelmValuesOk

`func (o *SyncStackWithOriginInput) GetDeleteStackHelmValuesOk() (*bool, bool)`

GetDeleteStackHelmValuesOk returns a tuple with the DeleteStackHelmValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackHelmValues

`func (o *SyncStackWithOriginInput) SetDeleteStackHelmValues(v bool)`

SetDeleteStackHelmValues sets DeleteStackHelmValues field to given value.

### HasDeleteStackHelmValues

`func (o *SyncStackWithOriginInput) HasDeleteStackHelmValues() bool`

HasDeleteStackHelmValues returns a boolean if a field has been set.

### SetDeleteStackHelmValuesNil

`func (o *SyncStackWithOriginInput) SetDeleteStackHelmValuesNil(b bool)`

 SetDeleteStackHelmValuesNil sets the value for DeleteStackHelmValues to be an explicit nil

### UnsetDeleteStackHelmValues
`func (o *SyncStackWithOriginInput) UnsetDeleteStackHelmValues()`

UnsetDeleteStackHelmValues ensures that no value is present for DeleteStackHelmValues, not even an explicit nil
### GetDeleteStackEnvVars

`func (o *SyncStackWithOriginInput) GetDeleteStackEnvVars() bool`

GetDeleteStackEnvVars returns the DeleteStackEnvVars field if non-nil, zero value otherwise.

### GetDeleteStackEnvVarsOk

`func (o *SyncStackWithOriginInput) GetDeleteStackEnvVarsOk() (*bool, bool)`

GetDeleteStackEnvVarsOk returns a tuple with the DeleteStackEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackEnvVars

`func (o *SyncStackWithOriginInput) SetDeleteStackEnvVars(v bool)`

SetDeleteStackEnvVars sets DeleteStackEnvVars field to given value.

### HasDeleteStackEnvVars

`func (o *SyncStackWithOriginInput) HasDeleteStackEnvVars() bool`

HasDeleteStackEnvVars returns a boolean if a field has been set.

### SetDeleteStackEnvVarsNil

`func (o *SyncStackWithOriginInput) SetDeleteStackEnvVarsNil(b bool)`

 SetDeleteStackEnvVarsNil sets the value for DeleteStackEnvVars to be an explicit nil

### UnsetDeleteStackEnvVars
`func (o *SyncStackWithOriginInput) UnsetDeleteStackEnvVars()`

UnsetDeleteStackEnvVars ensures that no value is present for DeleteStackEnvVars, not even an explicit nil
### GetDeleteStackTokens

`func (o *SyncStackWithOriginInput) GetDeleteStackTokens() bool`

GetDeleteStackTokens returns the DeleteStackTokens field if non-nil, zero value otherwise.

### GetDeleteStackTokensOk

`func (o *SyncStackWithOriginInput) GetDeleteStackTokensOk() (*bool, bool)`

GetDeleteStackTokensOk returns a tuple with the DeleteStackTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackTokens

`func (o *SyncStackWithOriginInput) SetDeleteStackTokens(v bool)`

SetDeleteStackTokens sets DeleteStackTokens field to given value.

### HasDeleteStackTokens

`func (o *SyncStackWithOriginInput) HasDeleteStackTokens() bool`

HasDeleteStackTokens returns a boolean if a field has been set.

### SetDeleteStackTokensNil

`func (o *SyncStackWithOriginInput) SetDeleteStackTokensNil(b bool)`

 SetDeleteStackTokensNil sets the value for DeleteStackTokens to be an explicit nil

### UnsetDeleteStackTokens
`func (o *SyncStackWithOriginInput) UnsetDeleteStackTokens()`

UnsetDeleteStackTokens ensures that no value is present for DeleteStackTokens, not even an explicit nil
### GetDeleteStackAnnotations

`func (o *SyncStackWithOriginInput) GetDeleteStackAnnotations() bool`

GetDeleteStackAnnotations returns the DeleteStackAnnotations field if non-nil, zero value otherwise.

### GetDeleteStackAnnotationsOk

`func (o *SyncStackWithOriginInput) GetDeleteStackAnnotationsOk() (*bool, bool)`

GetDeleteStackAnnotationsOk returns a tuple with the DeleteStackAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackAnnotations

`func (o *SyncStackWithOriginInput) SetDeleteStackAnnotations(v bool)`

SetDeleteStackAnnotations sets DeleteStackAnnotations field to given value.

### HasDeleteStackAnnotations

`func (o *SyncStackWithOriginInput) HasDeleteStackAnnotations() bool`

HasDeleteStackAnnotations returns a boolean if a field has been set.

### SetDeleteStackAnnotationsNil

`func (o *SyncStackWithOriginInput) SetDeleteStackAnnotationsNil(b bool)`

 SetDeleteStackAnnotationsNil sets the value for DeleteStackAnnotations to be an explicit nil

### UnsetDeleteStackAnnotations
`func (o *SyncStackWithOriginInput) UnsetDeleteStackAnnotations()`

UnsetDeleteStackAnnotations ensures that no value is present for DeleteStackAnnotations, not even an explicit nil
### GetDeleteStackServices

`func (o *SyncStackWithOriginInput) GetDeleteStackServices() bool`

GetDeleteStackServices returns the DeleteStackServices field if non-nil, zero value otherwise.

### GetDeleteStackServicesOk

`func (o *SyncStackWithOriginInput) GetDeleteStackServicesOk() (*bool, bool)`

GetDeleteStackServicesOk returns a tuple with the DeleteStackServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackServices

`func (o *SyncStackWithOriginInput) SetDeleteStackServices(v bool)`

SetDeleteStackServices sets DeleteStackServices field to given value.

### HasDeleteStackServices

`func (o *SyncStackWithOriginInput) HasDeleteStackServices() bool`

HasDeleteStackServices returns a boolean if a field has been set.

### SetDeleteStackServicesNil

`func (o *SyncStackWithOriginInput) SetDeleteStackServicesNil(b bool)`

 SetDeleteStackServicesNil sets the value for DeleteStackServices to be an explicit nil

### UnsetDeleteStackServices
`func (o *SyncStackWithOriginInput) UnsetDeleteStackServices()`

UnsetDeleteStackServices ensures that no value is present for DeleteStackServices, not even an explicit nil
### GetDeleteStackServicesConfiguration

`func (o *SyncStackWithOriginInput) GetDeleteStackServicesConfiguration() bool`

GetDeleteStackServicesConfiguration returns the DeleteStackServicesConfiguration field if non-nil, zero value otherwise.

### GetDeleteStackServicesConfigurationOk

`func (o *SyncStackWithOriginInput) GetDeleteStackServicesConfigurationOk() (*bool, bool)`

GetDeleteStackServicesConfigurationOk returns a tuple with the DeleteStackServicesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStackServicesConfiguration

`func (o *SyncStackWithOriginInput) SetDeleteStackServicesConfiguration(v bool)`

SetDeleteStackServicesConfiguration sets DeleteStackServicesConfiguration field to given value.

### HasDeleteStackServicesConfiguration

`func (o *SyncStackWithOriginInput) HasDeleteStackServicesConfiguration() bool`

HasDeleteStackServicesConfiguration returns a boolean if a field has been set.

### SetDeleteStackServicesConfigurationNil

`func (o *SyncStackWithOriginInput) SetDeleteStackServicesConfigurationNil(b bool)`

 SetDeleteStackServicesConfigurationNil sets the value for DeleteStackServicesConfiguration to be an explicit nil

### UnsetDeleteStackServicesConfiguration
`func (o *SyncStackWithOriginInput) UnsetDeleteStackServicesConfiguration()`

UnsetDeleteStackServicesConfiguration ensures that no value is present for DeleteStackServicesConfiguration, not even an explicit nil
### GetUpdateStackServices

`func (o *SyncStackWithOriginInput) GetUpdateStackServices() bool`

GetUpdateStackServices returns the UpdateStackServices field if non-nil, zero value otherwise.

### GetUpdateStackServicesOk

`func (o *SyncStackWithOriginInput) GetUpdateStackServicesOk() (*bool, bool)`

GetUpdateStackServicesOk returns a tuple with the UpdateStackServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStackServices

`func (o *SyncStackWithOriginInput) SetUpdateStackServices(v bool)`

SetUpdateStackServices sets UpdateStackServices field to given value.

### HasUpdateStackServices

`func (o *SyncStackWithOriginInput) HasUpdateStackServices() bool`

HasUpdateStackServices returns a boolean if a field has been set.

### SetUpdateStackServicesNil

`func (o *SyncStackWithOriginInput) SetUpdateStackServicesNil(b bool)`

 SetUpdateStackServicesNil sets the value for UpdateStackServices to be an explicit nil

### UnsetUpdateStackServices
`func (o *SyncStackWithOriginInput) UnsetUpdateStackServices()`

UnsetUpdateStackServices ensures that no value is present for UpdateStackServices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


