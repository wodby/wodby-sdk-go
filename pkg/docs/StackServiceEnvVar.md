# StackServiceEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StackServiceId** | **int32** |  | 
**Workload** | **string** |  | 
**Container** | **string** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**ValueSecretId** | Pointer to **NullableInt32** |  | [optional] 
**EnvType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStackServiceEnvVar

`func NewStackServiceEnvVar(id int32, stackServiceId int32, workload string, container string, name string, createdAt time.Time, ) *StackServiceEnvVar`

NewStackServiceEnvVar instantiates a new StackServiceEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceEnvVarWithDefaults

`func NewStackServiceEnvVarWithDefaults() *StackServiceEnvVar`

NewStackServiceEnvVarWithDefaults instantiates a new StackServiceEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceEnvVar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceEnvVar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceEnvVar) SetId(v int32)`

SetId sets Id field to given value.


### GetStackServiceId

`func (o *StackServiceEnvVar) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceEnvVar) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceEnvVar) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetWorkload

`func (o *StackServiceEnvVar) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *StackServiceEnvVar) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *StackServiceEnvVar) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetContainer

`func (o *StackServiceEnvVar) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *StackServiceEnvVar) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *StackServiceEnvVar) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetName

`func (o *StackServiceEnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceEnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceEnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StackServiceEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StackServiceEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StackServiceEnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StackServiceEnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StackServiceEnvVar) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StackServiceEnvVar) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueSecretId

`func (o *StackServiceEnvVar) GetValueSecretId() int32`

GetValueSecretId returns the ValueSecretId field if non-nil, zero value otherwise.

### GetValueSecretIdOk

`func (o *StackServiceEnvVar) GetValueSecretIdOk() (*int32, bool)`

GetValueSecretIdOk returns a tuple with the ValueSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSecretId

`func (o *StackServiceEnvVar) SetValueSecretId(v int32)`

SetValueSecretId sets ValueSecretId field to given value.

### HasValueSecretId

`func (o *StackServiceEnvVar) HasValueSecretId() bool`

HasValueSecretId returns a boolean if a field has been set.

### SetValueSecretIdNil

`func (o *StackServiceEnvVar) SetValueSecretIdNil(b bool)`

 SetValueSecretIdNil sets the value for ValueSecretId to be an explicit nil

### UnsetValueSecretId
`func (o *StackServiceEnvVar) UnsetValueSecretId()`

UnsetValueSecretId ensures that no value is present for ValueSecretId, not even an explicit nil
### GetEnvType

`func (o *StackServiceEnvVar) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *StackServiceEnvVar) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *StackServiceEnvVar) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *StackServiceEnvVar) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *StackServiceEnvVar) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *StackServiceEnvVar) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetCreatedAt

`func (o *StackServiceEnvVar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackServiceEnvVar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackServiceEnvVar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


