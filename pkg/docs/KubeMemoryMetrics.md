# KubeMemoryMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Available** | **int32** |  | 

## Methods

### NewKubeMemoryMetrics

`func NewKubeMemoryMetrics(total int32, available int32, ) *KubeMemoryMetrics`

NewKubeMemoryMetrics instantiates a new KubeMemoryMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeMemoryMetricsWithDefaults

`func NewKubeMemoryMetricsWithDefaults() *KubeMemoryMetrics`

NewKubeMemoryMetricsWithDefaults instantiates a new KubeMemoryMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *KubeMemoryMetrics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KubeMemoryMetrics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KubeMemoryMetrics) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetAvailable

`func (o *KubeMemoryMetrics) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *KubeMemoryMetrics) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *KubeMemoryMetrics) SetAvailable(v int32)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


