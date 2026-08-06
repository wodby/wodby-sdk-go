# AppServiceScalabilityUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**AverageCPU** | Pointer to **NullableInt32** |  | [optional] 
**MinReplicas** | Pointer to **NullableInt32** |  | [optional] 
**MaxReplicas** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppServiceScalabilityUpdateInput

`func NewAppServiceScalabilityUpdateInput(enabled bool, ) *AppServiceScalabilityUpdateInput`

NewAppServiceScalabilityUpdateInput instantiates a new AppServiceScalabilityUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceScalabilityUpdateInputWithDefaults

`func NewAppServiceScalabilityUpdateInputWithDefaults() *AppServiceScalabilityUpdateInput`

NewAppServiceScalabilityUpdateInputWithDefaults instantiates a new AppServiceScalabilityUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppServiceScalabilityUpdateInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppServiceScalabilityUpdateInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppServiceScalabilityUpdateInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAverageCPU

`func (o *AppServiceScalabilityUpdateInput) GetAverageCPU() int32`

GetAverageCPU returns the AverageCPU field if non-nil, zero value otherwise.

### GetAverageCPUOk

`func (o *AppServiceScalabilityUpdateInput) GetAverageCPUOk() (*int32, bool)`

GetAverageCPUOk returns a tuple with the AverageCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCPU

`func (o *AppServiceScalabilityUpdateInput) SetAverageCPU(v int32)`

SetAverageCPU sets AverageCPU field to given value.

### HasAverageCPU

`func (o *AppServiceScalabilityUpdateInput) HasAverageCPU() bool`

HasAverageCPU returns a boolean if a field has been set.

### SetAverageCPUNil

`func (o *AppServiceScalabilityUpdateInput) SetAverageCPUNil(b bool)`

 SetAverageCPUNil sets the value for AverageCPU to be an explicit nil

### UnsetAverageCPU
`func (o *AppServiceScalabilityUpdateInput) UnsetAverageCPU()`

UnsetAverageCPU ensures that no value is present for AverageCPU, not even an explicit nil
### GetMinReplicas

`func (o *AppServiceScalabilityUpdateInput) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *AppServiceScalabilityUpdateInput) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *AppServiceScalabilityUpdateInput) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *AppServiceScalabilityUpdateInput) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### SetMinReplicasNil

`func (o *AppServiceScalabilityUpdateInput) SetMinReplicasNil(b bool)`

 SetMinReplicasNil sets the value for MinReplicas to be an explicit nil

### UnsetMinReplicas
`func (o *AppServiceScalabilityUpdateInput) UnsetMinReplicas()`

UnsetMinReplicas ensures that no value is present for MinReplicas, not even an explicit nil
### GetMaxReplicas

`func (o *AppServiceScalabilityUpdateInput) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *AppServiceScalabilityUpdateInput) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *AppServiceScalabilityUpdateInput) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *AppServiceScalabilityUpdateInput) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### SetMaxReplicasNil

`func (o *AppServiceScalabilityUpdateInput) SetMaxReplicasNil(b bool)`

 SetMaxReplicasNil sets the value for MaxReplicas to be an explicit nil

### UnsetMaxReplicas
`func (o *AppServiceScalabilityUpdateInput) UnsetMaxReplicas()`

UnsetMaxReplicas ensures that no value is present for MaxReplicas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


