# GPUResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | Qualified device-plugin resource name, such as nvidia.com/gpu or amd.com/gpu. Native and reserved Kubernetes resource names are rejected. | 
**Count** | **int32** | Device count per container replica. Kubernetes requests and limits are equal; zero explicitly disables GPUs. | 

## Methods

### NewGPUResources

`func NewGPUResources(resourceName string, count int32, ) *GPUResources`

NewGPUResources instantiates a new GPUResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPUResourcesWithDefaults

`func NewGPUResourcesWithDefaults() *GPUResources`

NewGPUResourcesWithDefaults instantiates a new GPUResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *GPUResources) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *GPUResources) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *GPUResources) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetCount

`func (o *GPUResources) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GPUResources) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GPUResources) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


