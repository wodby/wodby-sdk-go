# ClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Ready** | **bool** |  | 
**Schedulable** | **bool** | Whether the node accepts new pods. A node being drained, recycled or upgraded doesn&#39;t. | 
**Version** | **string** | The node&#39;s Kubernetes version, as its kubelet reports it. | 
**MachineType** | Pointer to **string** | The provider&#39;s machine type of the node, such as m6i.large, when the node reports it. | [optional] 
**Zone** | Pointer to **string** | The availability zone of the node, when the node reports it. | [optional] 
**GpuNodePool** | **bool** | Whether the node belongs to the cluster&#39;s GPU node pool. | 
**InternalIp** | Pointer to **string** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**Capacity** | [**NodeCapacity**](NodeCapacity.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewClusterNode

`func NewClusterNode(name string, ready bool, schedulable bool, version string, gpuNodePool bool, capacity NodeCapacity, createdAt time.Time, ) *ClusterNode`

NewClusterNode instantiates a new ClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeWithDefaults

`func NewClusterNodeWithDefaults() *ClusterNode`

NewClusterNodeWithDefaults instantiates a new ClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNode) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *ClusterNode) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ClusterNode) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ClusterNode) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetSchedulable

`func (o *ClusterNode) GetSchedulable() bool`

GetSchedulable returns the Schedulable field if non-nil, zero value otherwise.

### GetSchedulableOk

`func (o *ClusterNode) GetSchedulableOk() (*bool, bool)`

GetSchedulableOk returns a tuple with the Schedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulable

`func (o *ClusterNode) SetSchedulable(v bool)`

SetSchedulable sets Schedulable field to given value.


### GetVersion

`func (o *ClusterNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterNode) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetMachineType

`func (o *ClusterNode) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *ClusterNode) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *ClusterNode) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *ClusterNode) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetZone

`func (o *ClusterNode) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ClusterNode) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ClusterNode) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ClusterNode) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetGpuNodePool

`func (o *ClusterNode) GetGpuNodePool() bool`

GetGpuNodePool returns the GpuNodePool field if non-nil, zero value otherwise.

### GetGpuNodePoolOk

`func (o *ClusterNode) GetGpuNodePoolOk() (*bool, bool)`

GetGpuNodePoolOk returns a tuple with the GpuNodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuNodePool

`func (o *ClusterNode) SetGpuNodePool(v bool)`

SetGpuNodePool sets GpuNodePool field to given value.


### GetInternalIp

`func (o *ClusterNode) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ClusterNode) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ClusterNode) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ClusterNode) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetExternalIp

`func (o *ClusterNode) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ClusterNode) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ClusterNode) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ClusterNode) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetCapacity

`func (o *ClusterNode) GetCapacity() NodeCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ClusterNode) GetCapacityOk() (*NodeCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ClusterNode) SetCapacity(v NodeCapacity)`

SetCapacity sets Capacity field to given value.


### GetCreatedAt

`func (o *ClusterNode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterNode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterNode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


