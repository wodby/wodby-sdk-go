# NewStackServiceEnvVarInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workload** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Value** | **string** |  | 
**Secret** | **bool** |  | 
**EnvType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewStackServiceEnvVarInput

`func NewNewStackServiceEnvVarInput(name string, value string, secret bool, ) *NewStackServiceEnvVarInput`

NewNewStackServiceEnvVarInput instantiates a new NewStackServiceEnvVarInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStackServiceEnvVarInputWithDefaults

`func NewNewStackServiceEnvVarInputWithDefaults() *NewStackServiceEnvVarInput`

NewNewStackServiceEnvVarInputWithDefaults instantiates a new NewStackServiceEnvVarInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkload

`func (o *NewStackServiceEnvVarInput) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *NewStackServiceEnvVarInput) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *NewStackServiceEnvVarInput) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *NewStackServiceEnvVarInput) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *NewStackServiceEnvVarInput) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *NewStackServiceEnvVarInput) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetContainer

`func (o *NewStackServiceEnvVarInput) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *NewStackServiceEnvVarInput) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *NewStackServiceEnvVarInput) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *NewStackServiceEnvVarInput) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *NewStackServiceEnvVarInput) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *NewStackServiceEnvVarInput) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetName

`func (o *NewStackServiceEnvVarInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStackServiceEnvVarInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStackServiceEnvVarInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewStackServiceEnvVarInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewStackServiceEnvVarInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewStackServiceEnvVarInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *NewStackServiceEnvVarInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NewStackServiceEnvVarInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NewStackServiceEnvVarInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetEnvType

`func (o *NewStackServiceEnvVarInput) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *NewStackServiceEnvVarInput) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *NewStackServiceEnvVarInput) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *NewStackServiceEnvVarInput) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *NewStackServiceEnvVarInput) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *NewStackServiceEnvVarInput) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


