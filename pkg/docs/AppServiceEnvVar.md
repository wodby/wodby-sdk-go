# AppServiceEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AppServiceId** | **int32** |  | 
**Workload** | **string** |  | 
**Container** | **string** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**Runtime** | **bool** |  | 
**Build** | **bool** |  | 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**NullableAppServiceEnvVarSource**](AppServiceEnvVarSource.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAppServiceEnvVar

`func NewAppServiceEnvVar(id int32, appServiceId int32, workload string, container string, name string, value string, runtime bool, build bool, ) *AppServiceEnvVar`

NewAppServiceEnvVar instantiates a new AppServiceEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceEnvVarWithDefaults

`func NewAppServiceEnvVarWithDefaults() *AppServiceEnvVar`

NewAppServiceEnvVarWithDefaults instantiates a new AppServiceEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppServiceEnvVar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceEnvVar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceEnvVar) SetId(v int32)`

SetId sets Id field to given value.


### GetAppServiceId

`func (o *AppServiceEnvVar) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *AppServiceEnvVar) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *AppServiceEnvVar) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetWorkload

`func (o *AppServiceEnvVar) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *AppServiceEnvVar) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *AppServiceEnvVar) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetContainer

`func (o *AppServiceEnvVar) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AppServiceEnvVar) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AppServiceEnvVar) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetName

`func (o *AppServiceEnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceEnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceEnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppServiceEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceEnvVar) SetValue(v string)`

SetValue sets Value field to given value.


### GetValueSecretId

`func (o *AppServiceEnvVar) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *AppServiceEnvVar) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *AppServiceEnvVar) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *AppServiceEnvVar) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *AppServiceEnvVar) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *AppServiceEnvVar) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetRuntime

`func (o *AppServiceEnvVar) GetRuntime() bool`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *AppServiceEnvVar) GetRuntimeOk() (*bool, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *AppServiceEnvVar) SetRuntime(v bool)`

SetRuntime sets Runtime field to given value.


### GetBuild

`func (o *AppServiceEnvVar) GetBuild() bool`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *AppServiceEnvVar) GetBuildOk() (*bool, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *AppServiceEnvVar) SetBuild(v bool)`

SetBuild sets Build field to given value.


### GetEnvType

`func (o *AppServiceEnvVar) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *AppServiceEnvVar) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *AppServiceEnvVar) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *AppServiceEnvVar) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *AppServiceEnvVar) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *AppServiceEnvVar) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetSource

`func (o *AppServiceEnvVar) GetSource() AppServiceEnvVarSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppServiceEnvVar) GetSourceOk() (*AppServiceEnvVarSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppServiceEnvVar) SetSource(v AppServiceEnvVarSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AppServiceEnvVar) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AppServiceEnvVar) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AppServiceEnvVar) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCreatedAt

`func (o *AppServiceEnvVar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppServiceEnvVar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppServiceEnvVar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppServiceEnvVar) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AppServiceEnvVar) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AppServiceEnvVar) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


