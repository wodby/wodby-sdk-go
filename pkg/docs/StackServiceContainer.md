# StackServiceContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**StackServiceId** | **int32** |  | 
**Workload** | **string** |  | 
**Name** | **string** |  | 
**RequestCPU** | Pointer to **NullableInt32** |  | [optional] 
**RequestMem** | Pointer to **NullableInt32** |  | [optional] 
**LimitCPU** | Pointer to **NullableInt32** |  | [optional] 
**LimitMem** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewStackServiceContainer

`func NewStackServiceContainer(stackServiceId int32, workload string, name string, ) *StackServiceContainer`

NewStackServiceContainer instantiates a new StackServiceContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceContainerWithDefaults

`func NewStackServiceContainerWithDefaults() *StackServiceContainer`

NewStackServiceContainerWithDefaults instantiates a new StackServiceContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackServiceContainer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackServiceContainer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackServiceContainer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StackServiceContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StackServiceContainer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StackServiceContainer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStackServiceId

`func (o *StackServiceContainer) GetStackServiceId() int32`

GetStackServiceId returns the StackServiceId field if non-nil, zero value otherwise.

### GetStackServiceIdOk

`func (o *StackServiceContainer) GetStackServiceIdOk() (*int32, bool)`

GetStackServiceIdOk returns a tuple with the StackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackServiceId

`func (o *StackServiceContainer) SetStackServiceId(v int32)`

SetStackServiceId sets StackServiceId field to given value.


### GetWorkload

`func (o *StackServiceContainer) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *StackServiceContainer) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *StackServiceContainer) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetName

`func (o *StackServiceContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackServiceContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackServiceContainer) SetName(v string)`

SetName sets Name field to given value.


### GetRequestCPU

`func (o *StackServiceContainer) GetRequestCPU() int32`

GetRequestCPU returns the RequestCPU field if non-nil, zero value otherwise.

### GetRequestCPUOk

`func (o *StackServiceContainer) GetRequestCPUOk() (*int32, bool)`

GetRequestCPUOk returns a tuple with the RequestCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCPU

`func (o *StackServiceContainer) SetRequestCPU(v int32)`

SetRequestCPU sets RequestCPU field to given value.

### HasRequestCPU

`func (o *StackServiceContainer) HasRequestCPU() bool`

HasRequestCPU returns a boolean if a field has been set.

### SetRequestCPUNil

`func (o *StackServiceContainer) SetRequestCPUNil(b bool)`

 SetRequestCPUNil sets the value for RequestCPU to be an explicit nil

### UnsetRequestCPU
`func (o *StackServiceContainer) UnsetRequestCPU()`

UnsetRequestCPU ensures that no value is present for RequestCPU, not even an explicit nil
### GetRequestMem

`func (o *StackServiceContainer) GetRequestMem() int32`

GetRequestMem returns the RequestMem field if non-nil, zero value otherwise.

### GetRequestMemOk

`func (o *StackServiceContainer) GetRequestMemOk() (*int32, bool)`

GetRequestMemOk returns a tuple with the RequestMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMem

`func (o *StackServiceContainer) SetRequestMem(v int32)`

SetRequestMem sets RequestMem field to given value.

### HasRequestMem

`func (o *StackServiceContainer) HasRequestMem() bool`

HasRequestMem returns a boolean if a field has been set.

### SetRequestMemNil

`func (o *StackServiceContainer) SetRequestMemNil(b bool)`

 SetRequestMemNil sets the value for RequestMem to be an explicit nil

### UnsetRequestMem
`func (o *StackServiceContainer) UnsetRequestMem()`

UnsetRequestMem ensures that no value is present for RequestMem, not even an explicit nil
### GetLimitCPU

`func (o *StackServiceContainer) GetLimitCPU() int32`

GetLimitCPU returns the LimitCPU field if non-nil, zero value otherwise.

### GetLimitCPUOk

`func (o *StackServiceContainer) GetLimitCPUOk() (*int32, bool)`

GetLimitCPUOk returns a tuple with the LimitCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCPU

`func (o *StackServiceContainer) SetLimitCPU(v int32)`

SetLimitCPU sets LimitCPU field to given value.

### HasLimitCPU

`func (o *StackServiceContainer) HasLimitCPU() bool`

HasLimitCPU returns a boolean if a field has been set.

### SetLimitCPUNil

`func (o *StackServiceContainer) SetLimitCPUNil(b bool)`

 SetLimitCPUNil sets the value for LimitCPU to be an explicit nil

### UnsetLimitCPU
`func (o *StackServiceContainer) UnsetLimitCPU()`

UnsetLimitCPU ensures that no value is present for LimitCPU, not even an explicit nil
### GetLimitMem

`func (o *StackServiceContainer) GetLimitMem() int32`

GetLimitMem returns the LimitMem field if non-nil, zero value otherwise.

### GetLimitMemOk

`func (o *StackServiceContainer) GetLimitMemOk() (*int32, bool)`

GetLimitMemOk returns a tuple with the LimitMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMem

`func (o *StackServiceContainer) SetLimitMem(v int32)`

SetLimitMem sets LimitMem field to given value.

### HasLimitMem

`func (o *StackServiceContainer) HasLimitMem() bool`

HasLimitMem returns a boolean if a field has been set.

### SetLimitMemNil

`func (o *StackServiceContainer) SetLimitMemNil(b bool)`

 SetLimitMemNil sets the value for LimitMem to be an explicit nil

### UnsetLimitMem
`func (o *StackServiceContainer) UnsetLimitMem()`

UnsetLimitMem ensures that no value is present for LimitMem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


