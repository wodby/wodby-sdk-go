# IntegrationEnvironmentPolicyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryEnvId** | Pointer to **NullableInt32** |  | [optional] 
**Scope** | **string** |  | 
**AllowedEnvIds** | **[]int32** |  | 

## Methods

### NewIntegrationEnvironmentPolicyInput

`func NewIntegrationEnvironmentPolicyInput(scope string, allowedEnvIds []int32, ) *IntegrationEnvironmentPolicyInput`

NewIntegrationEnvironmentPolicyInput instantiates a new IntegrationEnvironmentPolicyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEnvironmentPolicyInputWithDefaults

`func NewIntegrationEnvironmentPolicyInputWithDefaults() *IntegrationEnvironmentPolicyInput`

NewIntegrationEnvironmentPolicyInputWithDefaults instantiates a new IntegrationEnvironmentPolicyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryEnvId

`func (o *IntegrationEnvironmentPolicyInput) GetPrimaryEnvId() int32`

GetPrimaryEnvId returns the PrimaryEnvId field if non-nil, zero value otherwise.

### GetPrimaryEnvIdOk

`func (o *IntegrationEnvironmentPolicyInput) GetPrimaryEnvIdOk() (*int32, bool)`

GetPrimaryEnvIdOk returns a tuple with the PrimaryEnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnvId

`func (o *IntegrationEnvironmentPolicyInput) SetPrimaryEnvId(v int32)`

SetPrimaryEnvId sets PrimaryEnvId field to given value.

### HasPrimaryEnvId

`func (o *IntegrationEnvironmentPolicyInput) HasPrimaryEnvId() bool`

HasPrimaryEnvId returns a boolean if a field has been set.

### SetPrimaryEnvIdNil

`func (o *IntegrationEnvironmentPolicyInput) SetPrimaryEnvIdNil(b bool)`

 SetPrimaryEnvIdNil sets the value for PrimaryEnvId to be an explicit nil

### UnsetPrimaryEnvId
`func (o *IntegrationEnvironmentPolicyInput) UnsetPrimaryEnvId()`

UnsetPrimaryEnvId ensures that no value is present for PrimaryEnvId, not even an explicit nil
### GetScope

`func (o *IntegrationEnvironmentPolicyInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IntegrationEnvironmentPolicyInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IntegrationEnvironmentPolicyInput) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAllowedEnvIds

`func (o *IntegrationEnvironmentPolicyInput) GetAllowedEnvIds() []int32`

GetAllowedEnvIds returns the AllowedEnvIds field if non-nil, zero value otherwise.

### GetAllowedEnvIdsOk

`func (o *IntegrationEnvironmentPolicyInput) GetAllowedEnvIdsOk() (*[]int32, bool)`

GetAllowedEnvIdsOk returns a tuple with the AllowedEnvIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvIds

`func (o *IntegrationEnvironmentPolicyInput) SetAllowedEnvIds(v []int32)`

SetAllowedEnvIds sets AllowedEnvIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


