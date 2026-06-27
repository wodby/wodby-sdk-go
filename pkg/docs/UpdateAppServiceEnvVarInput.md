# UpdateAppServiceEnvVarInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**Secret** | **bool** |  | 
**Runtime** | Pointer to **NullableBool** |  | [optional] 
**Build** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateAppServiceEnvVarInput

`func NewUpdateAppServiceEnvVarInput(secret bool, ) *UpdateAppServiceEnvVarInput`

NewUpdateAppServiceEnvVarInput instantiates a new UpdateAppServiceEnvVarInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppServiceEnvVarInputWithDefaults

`func NewUpdateAppServiceEnvVarInputWithDefaults() *UpdateAppServiceEnvVarInput`

NewUpdateAppServiceEnvVarInputWithDefaults instantiates a new UpdateAppServiceEnvVarInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateAppServiceEnvVarInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateAppServiceEnvVarInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateAppServiceEnvVarInput) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateAppServiceEnvVarInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateAppServiceEnvVarInput) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateAppServiceEnvVarInput) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *UpdateAppServiceEnvVarInput) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateAppServiceEnvVarInput) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateAppServiceEnvVarInput) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetRuntime

`func (o *UpdateAppServiceEnvVarInput) GetRuntime() bool`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *UpdateAppServiceEnvVarInput) GetRuntimeOk() (*bool, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *UpdateAppServiceEnvVarInput) SetRuntime(v bool)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *UpdateAppServiceEnvVarInput) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *UpdateAppServiceEnvVarInput) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *UpdateAppServiceEnvVarInput) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetBuild

`func (o *UpdateAppServiceEnvVarInput) GetBuild() bool`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *UpdateAppServiceEnvVarInput) GetBuildOk() (*bool, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *UpdateAppServiceEnvVarInput) SetBuild(v bool)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *UpdateAppServiceEnvVarInput) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *UpdateAppServiceEnvVarInput) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *UpdateAppServiceEnvVarInput) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


