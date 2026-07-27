# ClusterMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ClusterId** | **int32** |  | 
**NodesTotal** | **int32** |  | 
**NodesReady** | **int32** |  | 
**Cpu** | [**KubeCPUMetrics**](KubeCPUMetrics.md) |  | 
**Memory** | [**KubeMemoryMetrics**](KubeMemoryMetrics.md) |  | 
**KubeCPUCap** | [**CapacityMetricsFloat**](CapacityMetricsFloat.md) |  | 
**KubeMemoryCap** | [**CapacityMetrics**](CapacityMetrics.md) |  | 
**KubePodsCap** | [**CapacityMetrics**](CapacityMetrics.md) |  | 
**HostDisk** | Pointer to [**NodeDiskMetrics**](NodeDiskMetrics.md) |  | [optional] 

## Methods

### NewClusterMetrics

`func NewClusterMetrics(id int32, clusterId int32, nodesTotal int32, nodesReady int32, cpu KubeCPUMetrics, memory KubeMemoryMetrics, kubeCPUCap CapacityMetricsFloat, kubeMemoryCap CapacityMetrics, kubePodsCap CapacityMetrics, ) *ClusterMetrics`

NewClusterMetrics instantiates a new ClusterMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMetricsWithDefaults

`func NewClusterMetricsWithDefaults() *ClusterMetrics`

NewClusterMetricsWithDefaults instantiates a new ClusterMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterMetrics) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterMetrics) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterMetrics) SetId(v int32)`

SetId sets Id field to given value.


### GetClusterId

`func (o *ClusterMetrics) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterMetrics) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterMetrics) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.


### GetNodesTotal

`func (o *ClusterMetrics) GetNodesTotal() int32`

GetNodesTotal returns the NodesTotal field if non-nil, zero value otherwise.

### GetNodesTotalOk

`func (o *ClusterMetrics) GetNodesTotalOk() (*int32, bool)`

GetNodesTotalOk returns a tuple with the NodesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesTotal

`func (o *ClusterMetrics) SetNodesTotal(v int32)`

SetNodesTotal sets NodesTotal field to given value.


### GetNodesReady

`func (o *ClusterMetrics) GetNodesReady() int32`

GetNodesReady returns the NodesReady field if non-nil, zero value otherwise.

### GetNodesReadyOk

`func (o *ClusterMetrics) GetNodesReadyOk() (*int32, bool)`

GetNodesReadyOk returns a tuple with the NodesReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesReady

`func (o *ClusterMetrics) SetNodesReady(v int32)`

SetNodesReady sets NodesReady field to given value.


### GetCpu

`func (o *ClusterMetrics) GetCpu() KubeCPUMetrics`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterMetrics) GetCpuOk() (*KubeCPUMetrics, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterMetrics) SetCpu(v KubeCPUMetrics)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ClusterMetrics) GetMemory() KubeMemoryMetrics`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterMetrics) GetMemoryOk() (*KubeMemoryMetrics, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterMetrics) SetMemory(v KubeMemoryMetrics)`

SetMemory sets Memory field to given value.


### GetKubeCPUCap

`func (o *ClusterMetrics) GetKubeCPUCap() CapacityMetricsFloat`

GetKubeCPUCap returns the KubeCPUCap field if non-nil, zero value otherwise.

### GetKubeCPUCapOk

`func (o *ClusterMetrics) GetKubeCPUCapOk() (*CapacityMetricsFloat, bool)`

GetKubeCPUCapOk returns a tuple with the KubeCPUCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeCPUCap

`func (o *ClusterMetrics) SetKubeCPUCap(v CapacityMetricsFloat)`

SetKubeCPUCap sets KubeCPUCap field to given value.


### GetKubeMemoryCap

`func (o *ClusterMetrics) GetKubeMemoryCap() CapacityMetrics`

GetKubeMemoryCap returns the KubeMemoryCap field if non-nil, zero value otherwise.

### GetKubeMemoryCapOk

`func (o *ClusterMetrics) GetKubeMemoryCapOk() (*CapacityMetrics, bool)`

GetKubeMemoryCapOk returns a tuple with the KubeMemoryCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeMemoryCap

`func (o *ClusterMetrics) SetKubeMemoryCap(v CapacityMetrics)`

SetKubeMemoryCap sets KubeMemoryCap field to given value.


### GetKubePodsCap

`func (o *ClusterMetrics) GetKubePodsCap() CapacityMetrics`

GetKubePodsCap returns the KubePodsCap field if non-nil, zero value otherwise.

### GetKubePodsCapOk

`func (o *ClusterMetrics) GetKubePodsCapOk() (*CapacityMetrics, bool)`

GetKubePodsCapOk returns a tuple with the KubePodsCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubePodsCap

`func (o *ClusterMetrics) SetKubePodsCap(v CapacityMetrics)`

SetKubePodsCap sets KubePodsCap field to given value.


### GetHostDisk

`func (o *ClusterMetrics) GetHostDisk() NodeDiskMetrics`

GetHostDisk returns the HostDisk field if non-nil, zero value otherwise.

### GetHostDiskOk

`func (o *ClusterMetrics) GetHostDiskOk() (*NodeDiskMetrics, bool)`

GetHostDiskOk returns a tuple with the HostDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDisk

`func (o *ClusterMetrics) SetHostDisk(v NodeDiskMetrics)`

SetHostDisk sets HostDisk field to given value.

### HasHostDisk

`func (o *ClusterMetrics) HasHostDisk() bool`

HasHostDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


