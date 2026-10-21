# UpdateClusterGPUNodePoolInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuNodePool** | [**NullableClusterNodePoolInput**](ClusterNodePoolInput.md) |  | 

## Methods

### NewUpdateClusterGPUNodePoolInput

`func NewUpdateClusterGPUNodePoolInput(gpuNodePool NullableClusterNodePoolInput, ) *UpdateClusterGPUNodePoolInput`

NewUpdateClusterGPUNodePoolInput instantiates a new UpdateClusterGPUNodePoolInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterGPUNodePoolInputWithDefaults

`func NewUpdateClusterGPUNodePoolInputWithDefaults() *UpdateClusterGPUNodePoolInput`

NewUpdateClusterGPUNodePoolInputWithDefaults instantiates a new UpdateClusterGPUNodePoolInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuNodePool

`func (o *UpdateClusterGPUNodePoolInput) GetGpuNodePool() ClusterNodePoolInput`

GetGpuNodePool returns the GpuNodePool field if non-nil, zero value otherwise.

### GetGpuNodePoolOk

`func (o *UpdateClusterGPUNodePoolInput) GetGpuNodePoolOk() (*ClusterNodePoolInput, bool)`

GetGpuNodePoolOk returns a tuple with the GpuNodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuNodePool

`func (o *UpdateClusterGPUNodePoolInput) SetGpuNodePool(v ClusterNodePoolInput)`

SetGpuNodePool sets GpuNodePool field to given value.


### SetGpuNodePoolNil

`func (o *UpdateClusterGPUNodePoolInput) SetGpuNodePoolNil(b bool)`

 SetGpuNodePoolNil sets the value for GpuNodePool to be an explicit nil

### UnsetGpuNodePool
`func (o *UpdateClusterGPUNodePoolInput) UnsetGpuNodePool()`

UnsetGpuNodePool ensures that no value is present for GpuNodePool, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


