# ServiceManifestEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Secret** | **bool** |  | 
**Env** | Pointer to **string** |  | [optional] 
**Runtime** | **bool** |  | 
**Build** | **bool** |  | 

## Methods

### NewServiceManifestEnvVar

`func NewServiceManifestEnvVar(name string, value string, secret bool, runtime bool, build bool, ) *ServiceManifestEnvVar`

NewServiceManifestEnvVar instantiates a new ServiceManifestEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestEnvVarWithDefaults

`func NewServiceManifestEnvVarWithDefaults() *ServiceManifestEnvVar`

NewServiceManifestEnvVarWithDefaults instantiates a new ServiceManifestEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceManifestEnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceManifestEnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceManifestEnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ServiceManifestEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceManifestEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceManifestEnvVar) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *ServiceManifestEnvVar) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceManifestEnvVar) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceManifestEnvVar) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetEnv

`func (o *ServiceManifestEnvVar) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServiceManifestEnvVar) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ServiceManifestEnvVar) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ServiceManifestEnvVar) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetRuntime

`func (o *ServiceManifestEnvVar) GetRuntime() bool`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ServiceManifestEnvVar) GetRuntimeOk() (*bool, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ServiceManifestEnvVar) SetRuntime(v bool)`

SetRuntime sets Runtime field to given value.


### GetBuild

`func (o *ServiceManifestEnvVar) GetBuild() bool`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ServiceManifestEnvVar) GetBuildOk() (*bool, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ServiceManifestEnvVar) SetBuild(v bool)`

SetBuild sets Build field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


