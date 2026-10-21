# ClusterNodePoolInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineType** | **string** |  | 
**MinNodeCount** | **int32** | Zero lets the pool scale down to no nodes where the provider supports it. | 
**MaxNodeCount** | **int32** |  | 
**NodeDiskSize** | Pointer to **NullableInt32** | Node disk size in GB. Required when the provider&#39;s node disk size is configurable. | [optional] 

## Methods

### NewClusterNodePoolInput

`func NewClusterNodePoolInput(machineType string, minNodeCount int32, maxNodeCount int32, ) *ClusterNodePoolInput`

NewClusterNodePoolInput instantiates a new ClusterNodePoolInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodePoolInputWithDefaults

`func NewClusterNodePoolInputWithDefaults() *ClusterNodePoolInput`

NewClusterNodePoolInputWithDefaults instantiates a new ClusterNodePoolInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineType

`func (o *ClusterNodePoolInput) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *ClusterNodePoolInput) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *ClusterNodePoolInput) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.


### GetMinNodeCount

`func (o *ClusterNodePoolInput) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *ClusterNodePoolInput) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *ClusterNodePoolInput) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.


### GetMaxNodeCount

`func (o *ClusterNodePoolInput) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *ClusterNodePoolInput) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *ClusterNodePoolInput) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.


### GetNodeDiskSize

`func (o *ClusterNodePoolInput) GetNodeDiskSize() int32`

GetNodeDiskSize returns the NodeDiskSize field if non-nil, zero value otherwise.

### GetNodeDiskSizeOk

`func (o *ClusterNodePoolInput) GetNodeDiskSizeOk() (*int32, bool)`

GetNodeDiskSizeOk returns a tuple with the NodeDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDiskSize

`func (o *ClusterNodePoolInput) SetNodeDiskSize(v int32)`

SetNodeDiskSize sets NodeDiskSize field to given value.

### HasNodeDiskSize

`func (o *ClusterNodePoolInput) HasNodeDiskSize() bool`

HasNodeDiskSize returns a boolean if a field has been set.

### SetNodeDiskSizeNil

`func (o *ClusterNodePoolInput) SetNodeDiskSizeNil(b bool)`

 SetNodeDiskSizeNil sets the value for NodeDiskSize to be an explicit nil

### UnsetNodeDiskSize
`func (o *ClusterNodePoolInput) UnsetNodeDiskSize()`

UnsetNodeDiskSize ensures that no value is present for NodeDiskSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


