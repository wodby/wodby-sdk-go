# KubeCPUMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cores** | **int32** |  | 
**Load1** | **float64** |  | 
**Load5** | **float64** |  | 
**Load15** | **float64** |  | 
**Idle** | Pointer to **NullableFloat64** |  | [optional] 
**IoWait** | Pointer to **NullableFloat64** |  | [optional] 
**Steal** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewKubeCPUMetrics

`func NewKubeCPUMetrics(cores int32, load1 float64, load5 float64, load15 float64, ) *KubeCPUMetrics`

NewKubeCPUMetrics instantiates a new KubeCPUMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeCPUMetricsWithDefaults

`func NewKubeCPUMetricsWithDefaults() *KubeCPUMetrics`

NewKubeCPUMetricsWithDefaults instantiates a new KubeCPUMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *KubeCPUMetrics) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *KubeCPUMetrics) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *KubeCPUMetrics) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetLoad1

`func (o *KubeCPUMetrics) GetLoad1() float64`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *KubeCPUMetrics) GetLoad1Ok() (*float64, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *KubeCPUMetrics) SetLoad1(v float64)`

SetLoad1 sets Load1 field to given value.


### GetLoad5

`func (o *KubeCPUMetrics) GetLoad5() float64`

GetLoad5 returns the Load5 field if non-nil, zero value otherwise.

### GetLoad5Ok

`func (o *KubeCPUMetrics) GetLoad5Ok() (*float64, bool)`

GetLoad5Ok returns a tuple with the Load5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5

`func (o *KubeCPUMetrics) SetLoad5(v float64)`

SetLoad5 sets Load5 field to given value.


### GetLoad15

`func (o *KubeCPUMetrics) GetLoad15() float64`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *KubeCPUMetrics) GetLoad15Ok() (*float64, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *KubeCPUMetrics) SetLoad15(v float64)`

SetLoad15 sets Load15 field to given value.


### GetIdle

`func (o *KubeCPUMetrics) GetIdle() float64`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *KubeCPUMetrics) GetIdleOk() (*float64, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *KubeCPUMetrics) SetIdle(v float64)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *KubeCPUMetrics) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### SetIdleNil

`func (o *KubeCPUMetrics) SetIdleNil(b bool)`

 SetIdleNil sets the value for Idle to be an explicit nil

### UnsetIdle
`func (o *KubeCPUMetrics) UnsetIdle()`

UnsetIdle ensures that no value is present for Idle, not even an explicit nil
### GetIoWait

`func (o *KubeCPUMetrics) GetIoWait() float64`

GetIoWait returns the IoWait field if non-nil, zero value otherwise.

### GetIoWaitOk

`func (o *KubeCPUMetrics) GetIoWaitOk() (*float64, bool)`

GetIoWaitOk returns a tuple with the IoWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoWait

`func (o *KubeCPUMetrics) SetIoWait(v float64)`

SetIoWait sets IoWait field to given value.

### HasIoWait

`func (o *KubeCPUMetrics) HasIoWait() bool`

HasIoWait returns a boolean if a field has been set.

### SetIoWaitNil

`func (o *KubeCPUMetrics) SetIoWaitNil(b bool)`

 SetIoWaitNil sets the value for IoWait to be an explicit nil

### UnsetIoWait
`func (o *KubeCPUMetrics) UnsetIoWait()`

UnsetIoWait ensures that no value is present for IoWait, not even an explicit nil
### GetSteal

`func (o *KubeCPUMetrics) GetSteal() float64`

GetSteal returns the Steal field if non-nil, zero value otherwise.

### GetStealOk

`func (o *KubeCPUMetrics) GetStealOk() (*float64, bool)`

GetStealOk returns a tuple with the Steal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteal

`func (o *KubeCPUMetrics) SetSteal(v float64)`

SetSteal sets Steal field to given value.

### HasSteal

`func (o *KubeCPUMetrics) HasSteal() bool`

HasSteal returns a boolean if a field has been set.

### SetStealNil

`func (o *KubeCPUMetrics) SetStealNil(b bool)`

 SetStealNil sets the value for Steal to be an explicit nil

### UnsetSteal
`func (o *KubeCPUMetrics) UnsetSteal()`

UnsetSteal ensures that no value is present for Steal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


