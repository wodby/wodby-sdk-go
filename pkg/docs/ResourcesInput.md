# ResourcesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workload** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**RequestCPU** | Pointer to **NullableInt32** |  | [optional] 
**RequestMem** | Pointer to **NullableInt32** |  | [optional] 
**LimitCPU** | Pointer to **NullableInt32** |  | [optional] 
**LimitMem** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewResourcesInput

`func NewResourcesInput() *ResourcesInput`

NewResourcesInput instantiates a new ResourcesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesInputWithDefaults

`func NewResourcesInputWithDefaults() *ResourcesInput`

NewResourcesInputWithDefaults instantiates a new ResourcesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkload

`func (o *ResourcesInput) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *ResourcesInput) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *ResourcesInput) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *ResourcesInput) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### SetWorkloadNil

`func (o *ResourcesInput) SetWorkloadNil(b bool)`

 SetWorkloadNil sets the value for Workload to be an explicit nil

### UnsetWorkload
`func (o *ResourcesInput) UnsetWorkload()`

UnsetWorkload ensures that no value is present for Workload, not even an explicit nil
### GetContainer

`func (o *ResourcesInput) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ResourcesInput) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ResourcesInput) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ResourcesInput) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ResourcesInput) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ResourcesInput) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetRequestCPU

`func (o *ResourcesInput) GetRequestCPU() int32`

GetRequestCPU returns the RequestCPU field if non-nil, zero value otherwise.

### GetRequestCPUOk

`func (o *ResourcesInput) GetRequestCPUOk() (*int32, bool)`

GetRequestCPUOk returns a tuple with the RequestCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCPU

`func (o *ResourcesInput) SetRequestCPU(v int32)`

SetRequestCPU sets RequestCPU field to given value.

### HasRequestCPU

`func (o *ResourcesInput) HasRequestCPU() bool`

HasRequestCPU returns a boolean if a field has been set.

### SetRequestCPUNil

`func (o *ResourcesInput) SetRequestCPUNil(b bool)`

 SetRequestCPUNil sets the value for RequestCPU to be an explicit nil

### UnsetRequestCPU
`func (o *ResourcesInput) UnsetRequestCPU()`

UnsetRequestCPU ensures that no value is present for RequestCPU, not even an explicit nil
### GetRequestMem

`func (o *ResourcesInput) GetRequestMem() int32`

GetRequestMem returns the RequestMem field if non-nil, zero value otherwise.

### GetRequestMemOk

`func (o *ResourcesInput) GetRequestMemOk() (*int32, bool)`

GetRequestMemOk returns a tuple with the RequestMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMem

`func (o *ResourcesInput) SetRequestMem(v int32)`

SetRequestMem sets RequestMem field to given value.

### HasRequestMem

`func (o *ResourcesInput) HasRequestMem() bool`

HasRequestMem returns a boolean if a field has been set.

### SetRequestMemNil

`func (o *ResourcesInput) SetRequestMemNil(b bool)`

 SetRequestMemNil sets the value for RequestMem to be an explicit nil

### UnsetRequestMem
`func (o *ResourcesInput) UnsetRequestMem()`

UnsetRequestMem ensures that no value is present for RequestMem, not even an explicit nil
### GetLimitCPU

`func (o *ResourcesInput) GetLimitCPU() int32`

GetLimitCPU returns the LimitCPU field if non-nil, zero value otherwise.

### GetLimitCPUOk

`func (o *ResourcesInput) GetLimitCPUOk() (*int32, bool)`

GetLimitCPUOk returns a tuple with the LimitCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCPU

`func (o *ResourcesInput) SetLimitCPU(v int32)`

SetLimitCPU sets LimitCPU field to given value.

### HasLimitCPU

`func (o *ResourcesInput) HasLimitCPU() bool`

HasLimitCPU returns a boolean if a field has been set.

### SetLimitCPUNil

`func (o *ResourcesInput) SetLimitCPUNil(b bool)`

 SetLimitCPUNil sets the value for LimitCPU to be an explicit nil

### UnsetLimitCPU
`func (o *ResourcesInput) UnsetLimitCPU()`

UnsetLimitCPU ensures that no value is present for LimitCPU, not even an explicit nil
### GetLimitMem

`func (o *ResourcesInput) GetLimitMem() int32`

GetLimitMem returns the LimitMem field if non-nil, zero value otherwise.

### GetLimitMemOk

`func (o *ResourcesInput) GetLimitMemOk() (*int32, bool)`

GetLimitMemOk returns a tuple with the LimitMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMem

`func (o *ResourcesInput) SetLimitMem(v int32)`

SetLimitMem sets LimitMem field to given value.

### HasLimitMem

`func (o *ResourcesInput) HasLimitMem() bool`

HasLimitMem returns a boolean if a field has been set.

### SetLimitMemNil

`func (o *ResourcesInput) SetLimitMemNil(b bool)`

 SetLimitMemNil sets the value for LimitMem to be an explicit nil

### UnsetLimitMem
`func (o *ResourcesInput) UnsetLimitMem()`

UnsetLimitMem ensures that no value is present for LimitMem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


