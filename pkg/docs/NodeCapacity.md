# NodeCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **float32** | CPU cores. | 
**Memory** | **int64** | Memory in bytes. | 
**Disk** | Pointer to **int64** | The size of the node&#39;s root disk in bytes, when the node reports it. | [optional] 
**Gpus** | **int32** | GPUs the node offers to pods. 0 without GPUs, and until the node&#39;s GPU driver reports them. | 

## Methods

### NewNodeCapacity

`func NewNodeCapacity(cpu float32, memory int64, gpus int32, ) *NodeCapacity`

NewNodeCapacity instantiates a new NodeCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeCapacityWithDefaults

`func NewNodeCapacityWithDefaults() *NodeCapacity`

NewNodeCapacityWithDefaults instantiates a new NodeCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NodeCapacity) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeCapacity) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeCapacity) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *NodeCapacity) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeCapacity) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeCapacity) SetMemory(v int64)`

SetMemory sets Memory field to given value.


### GetDisk

`func (o *NodeCapacity) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeCapacity) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeCapacity) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NodeCapacity) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGpus

`func (o *NodeCapacity) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *NodeCapacity) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *NodeCapacity) SetGpus(v int32)`

SetGpus sets Gpus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


