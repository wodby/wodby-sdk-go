# ClusterEnvironmentPolicyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvId** | Pointer to **int32** |  | [optional] 
**EnvType** | Pointer to **string** |  | [optional] 
**Scope** | **string** |  | 
**AllowedEnvIds** | Pointer to **[]int32** |  | [optional] 
**AllowedEnvTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClusterEnvironmentPolicyInput

`func NewClusterEnvironmentPolicyInput(scope string, ) *ClusterEnvironmentPolicyInput`

NewClusterEnvironmentPolicyInput instantiates a new ClusterEnvironmentPolicyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEnvironmentPolicyInputWithDefaults

`func NewClusterEnvironmentPolicyInputWithDefaults() *ClusterEnvironmentPolicyInput`

NewClusterEnvironmentPolicyInputWithDefaults instantiates a new ClusterEnvironmentPolicyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvId

`func (o *ClusterEnvironmentPolicyInput) GetEnvId() int32`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *ClusterEnvironmentPolicyInput) GetEnvIdOk() (*int32, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *ClusterEnvironmentPolicyInput) SetEnvId(v int32)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *ClusterEnvironmentPolicyInput) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetEnvType

`func (o *ClusterEnvironmentPolicyInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *ClusterEnvironmentPolicyInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *ClusterEnvironmentPolicyInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *ClusterEnvironmentPolicyInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### GetScope

`func (o *ClusterEnvironmentPolicyInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ClusterEnvironmentPolicyInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ClusterEnvironmentPolicyInput) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAllowedEnvIds

`func (o *ClusterEnvironmentPolicyInput) GetAllowedEnvIds() []int32`

GetAllowedEnvIds returns the AllowedEnvIds field if non-nil, zero value otherwise.

### GetAllowedEnvIdsOk

`func (o *ClusterEnvironmentPolicyInput) GetAllowedEnvIdsOk() (*[]int32, bool)`

GetAllowedEnvIdsOk returns a tuple with the AllowedEnvIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvIds

`func (o *ClusterEnvironmentPolicyInput) SetAllowedEnvIds(v []int32)`

SetAllowedEnvIds sets AllowedEnvIds field to given value.

### HasAllowedEnvIds

`func (o *ClusterEnvironmentPolicyInput) HasAllowedEnvIds() bool`

HasAllowedEnvIds returns a boolean if a field has been set.

### GetAllowedEnvTypes

`func (o *ClusterEnvironmentPolicyInput) GetAllowedEnvTypes() []string`

GetAllowedEnvTypes returns the AllowedEnvTypes field if non-nil, zero value otherwise.

### GetAllowedEnvTypesOk

`func (o *ClusterEnvironmentPolicyInput) GetAllowedEnvTypesOk() (*[]string, bool)`

GetAllowedEnvTypesOk returns a tuple with the AllowedEnvTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvTypes

`func (o *ClusterEnvironmentPolicyInput) SetAllowedEnvTypes(v []string)`

SetAllowedEnvTypes sets AllowedEnvTypes field to given value.

### HasAllowedEnvTypes

`func (o *ClusterEnvironmentPolicyInput) HasAllowedEnvTypes() bool`

HasAllowedEnvTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


