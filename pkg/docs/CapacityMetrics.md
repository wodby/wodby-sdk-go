# CapacityMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Allocated** | **int32** |  | 

## Methods

### NewCapacityMetrics

`func NewCapacityMetrics(total int32, allocated int32, ) *CapacityMetrics`

NewCapacityMetrics instantiates a new CapacityMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityMetricsWithDefaults

`func NewCapacityMetricsWithDefaults() *CapacityMetrics`

NewCapacityMetricsWithDefaults instantiates a new CapacityMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CapacityMetrics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CapacityMetrics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CapacityMetrics) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetAllocated

`func (o *CapacityMetrics) GetAllocated() int32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *CapacityMetrics) GetAllocatedOk() (*int32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *CapacityMetrics) SetAllocated(v int32)`

SetAllocated sets Allocated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


