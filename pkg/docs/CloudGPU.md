# CloudGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | **string** |  | 
**Model** | **string** |  | 
**Count** | **int32** |  | 
**MemoryMiB** | Pointer to **NullableInt32** | Memory per device in MiB; absent when the provider does not report it. | [optional] 

## Methods

### NewCloudGPU

`func NewCloudGPU(manufacturer string, model string, count int32, ) *CloudGPU`

NewCloudGPU instantiates a new CloudGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGPUWithDefaults

`func NewCloudGPUWithDefaults() *CloudGPU`

NewCloudGPUWithDefaults instantiates a new CloudGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *CloudGPU) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CloudGPU) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CloudGPU) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *CloudGPU) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudGPU) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudGPU) SetModel(v string)`

SetModel sets Model field to given value.


### GetCount

`func (o *CloudGPU) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudGPU) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudGPU) SetCount(v int32)`

SetCount sets Count field to given value.


### GetMemoryMiB

`func (o *CloudGPU) GetMemoryMiB() int32`

GetMemoryMiB returns the MemoryMiB field if non-nil, zero value otherwise.

### GetMemoryMiBOk

`func (o *CloudGPU) GetMemoryMiBOk() (*int32, bool)`

GetMemoryMiBOk returns a tuple with the MemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMiB

`func (o *CloudGPU) SetMemoryMiB(v int32)`

SetMemoryMiB sets MemoryMiB field to given value.

### HasMemoryMiB

`func (o *CloudGPU) HasMemoryMiB() bool`

HasMemoryMiB returns a boolean if a field has been set.

### SetMemoryMiBNil

`func (o *CloudGPU) SetMemoryMiBNil(b bool)`

 SetMemoryMiBNil sets the value for MemoryMiB to be an explicit nil

### UnsetMemoryMiB
`func (o *CloudGPU) UnsetMemoryMiB()`

UnsetMemoryMiB ensures that no value is present for MemoryMiB, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


