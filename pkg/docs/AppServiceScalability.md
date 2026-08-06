# AppServiceScalability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageCPU** | **int32** |  | 
**MinReplicas** | **int32** |  | 
**MaxReplicas** | **int32** |  | 

## Methods

### NewAppServiceScalability

`func NewAppServiceScalability(averageCPU int32, minReplicas int32, maxReplicas int32, ) *AppServiceScalability`

NewAppServiceScalability instantiates a new AppServiceScalability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceScalabilityWithDefaults

`func NewAppServiceScalabilityWithDefaults() *AppServiceScalability`

NewAppServiceScalabilityWithDefaults instantiates a new AppServiceScalability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageCPU

`func (o *AppServiceScalability) GetAverageCPU() int32`

GetAverageCPU returns the AverageCPU field if non-nil, zero value otherwise.

### GetAverageCPUOk

`func (o *AppServiceScalability) GetAverageCPUOk() (*int32, bool)`

GetAverageCPUOk returns a tuple with the AverageCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCPU

`func (o *AppServiceScalability) SetAverageCPU(v int32)`

SetAverageCPU sets AverageCPU field to given value.


### GetMinReplicas

`func (o *AppServiceScalability) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *AppServiceScalability) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *AppServiceScalability) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.


### GetMaxReplicas

`func (o *AppServiceScalability) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *AppServiceScalability) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *AppServiceScalability) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


