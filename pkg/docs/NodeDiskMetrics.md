# NodeDiskMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **string** |  | 
**FsType** | **string** |  | 
**Size** | **int32** |  | 
**Free** | **int32** |  | 
**INodes** | **int32** |  | 
**INodesFree** | **int32** |  | 

## Methods

### NewNodeDiskMetrics

`func NewNodeDiskMetrics(device string, fsType string, size int32, free int32, iNodes int32, iNodesFree int32, ) *NodeDiskMetrics`

NewNodeDiskMetrics instantiates a new NodeDiskMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDiskMetricsWithDefaults

`func NewNodeDiskMetricsWithDefaults() *NodeDiskMetrics`

NewNodeDiskMetricsWithDefaults instantiates a new NodeDiskMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *NodeDiskMetrics) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NodeDiskMetrics) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NodeDiskMetrics) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetFsType

`func (o *NodeDiskMetrics) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *NodeDiskMetrics) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *NodeDiskMetrics) SetFsType(v string)`

SetFsType sets FsType field to given value.


### GetSize

`func (o *NodeDiskMetrics) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NodeDiskMetrics) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NodeDiskMetrics) SetSize(v int32)`

SetSize sets Size field to given value.


### GetFree

`func (o *NodeDiskMetrics) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *NodeDiskMetrics) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *NodeDiskMetrics) SetFree(v int32)`

SetFree sets Free field to given value.


### GetINodes

`func (o *NodeDiskMetrics) GetINodes() int32`

GetINodes returns the INodes field if non-nil, zero value otherwise.

### GetINodesOk

`func (o *NodeDiskMetrics) GetINodesOk() (*int32, bool)`

GetINodesOk returns a tuple with the INodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINodes

`func (o *NodeDiskMetrics) SetINodes(v int32)`

SetINodes sets INodes field to given value.


### GetINodesFree

`func (o *NodeDiskMetrics) GetINodesFree() int32`

GetINodesFree returns the INodesFree field if non-nil, zero value otherwise.

### GetINodesFreeOk

`func (o *NodeDiskMetrics) GetINodesFreeOk() (*int32, bool)`

GetINodesFreeOk returns a tuple with the INodesFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINodesFree

`func (o *NodeDiskMetrics) SetINodesFree(v int32)`

SetINodesFree sets INodesFree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


