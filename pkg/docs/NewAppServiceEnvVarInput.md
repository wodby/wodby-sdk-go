# NewAppServiceEnvVarInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workload** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Value** | **string** |  | 
**Secret** | **bool** |  | 
**Runtime** | Pointer to **NullableBool** |  | [optional] 
**Build** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewNewAppServiceEnvVarInput

`func NewNewAppServiceEnvVarInput(name string, value string, secret bool, ) *NewAppServiceEnvVarInput`

NewNewAppServiceEnvVarInput instantiates a new NewAppServiceEnvVarInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppServiceEnvVarInputWithDefaults

`func NewNewAppServiceEnvVarInputWithDefaults() *NewAppServiceEnvVarInput`

NewNewAppServiceEnvVarInputWithDefaults instantiates a new NewAppServiceEnvVarInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkload

`func (o *NewAppServiceEnvVarInput) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *NewAppServiceEnvVarInput) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *NewAppServiceEnvVarInput) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *NewAppServiceEnvVarInput) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *NewAppServiceEnvVarInput) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *NewAppServiceEnvVarInput) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetContainer

`func (o *NewAppServiceEnvVarInput) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *NewAppServiceEnvVarInput) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *NewAppServiceEnvVarInput) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *NewAppServiceEnvVarInput) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *NewAppServiceEnvVarInput) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *NewAppServiceEnvVarInput) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetName

`func (o *NewAppServiceEnvVarInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAppServiceEnvVarInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewAppServiceEnvVarInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NewAppServiceEnvVarInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewAppServiceEnvVarInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewAppServiceEnvVarInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *NewAppServiceEnvVarInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NewAppServiceEnvVarInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NewAppServiceEnvVarInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetRuntime

`func (o *NewAppServiceEnvVarInput) GetRuntime() bool`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *NewAppServiceEnvVarInput) GetRuntimeOk() (*bool, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *NewAppServiceEnvVarInput) SetRuntime(v bool)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *NewAppServiceEnvVarInput) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *NewAppServiceEnvVarInput) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *NewAppServiceEnvVarInput) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetBuild

`func (o *NewAppServiceEnvVarInput) GetBuild() bool`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *NewAppServiceEnvVarInput) GetBuildOk() (*bool, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *NewAppServiceEnvVarInput) SetBuild(v bool)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *NewAppServiceEnvVarInput) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *NewAppServiceEnvVarInput) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *NewAppServiceEnvVarInput) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


