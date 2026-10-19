# CloudMachineType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Mem** | Pointer to **int32** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**SharedCPU** | Pointer to **bool** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**MinStorageSize** | Pointer to **NullableInt32** |  | [optional] 
**MaxStorageSize** | Pointer to **NullableInt32** |  | [optional] 
**MinIOPS** | Pointer to **NullableInt32** |  | [optional] 
**MaxIOPS** | Pointer to **NullableInt32** |  | [optional] 
**DefaultIOPS** | Pointer to **NullableInt32** |  | [optional] 
**StorageOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Gpus** | Pointer to [**[]CloudGPU**](CloudGPU.md) | Provider-reported devices. This does not guarantee driver setup or spare capacity. | [optional] 

## Methods

### NewCloudMachineType

`func NewCloudMachineType() *CloudMachineType`

NewCloudMachineType instantiates a new CloudMachineType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMachineTypeWithDefaults

`func NewCloudMachineTypeWithDefaults() *CloudMachineType`

NewCloudMachineTypeWithDefaults instantiates a new CloudMachineType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudMachineType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMachineType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMachineType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMachineType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpu

`func (o *CloudMachineType) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CloudMachineType) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CloudMachineType) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *CloudMachineType) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMem

`func (o *CloudMachineType) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *CloudMachineType) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *CloudMachineType) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *CloudMachineType) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetDisk

`func (o *CloudMachineType) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CloudMachineType) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CloudMachineType) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *CloudMachineType) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetSharedCPU

`func (o *CloudMachineType) GetSharedCPU() bool`

GetSharedCPU returns the SharedCPU field if non-nil, zero value otherwise.

### GetSharedCPUOk

`func (o *CloudMachineType) GetSharedCPUOk() (*bool, bool)`

GetSharedCPUOk returns a tuple with the SharedCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedCPU

`func (o *CloudMachineType) SetSharedCPU(v bool)`

SetSharedCPU sets SharedCPU field to given value.

### HasSharedCPU

`func (o *CloudMachineType) HasSharedCPU() bool`

HasSharedCPU returns a boolean if a field has been set.

### GetAvailable

`func (o *CloudMachineType) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudMachineType) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudMachineType) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudMachineType) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetFamily

`func (o *CloudMachineType) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *CloudMachineType) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *CloudMachineType) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *CloudMachineType) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetMinStorageSize

`func (o *CloudMachineType) GetMinStorageSize() int32`

GetMinStorageSize returns the MinStorageSize field if non-nil, zero value otherwise.

### GetMinStorageSizeOk

`func (o *CloudMachineType) GetMinStorageSizeOk() (*int32, bool)`

GetMinStorageSizeOk returns a tuple with the MinStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorageSize

`func (o *CloudMachineType) SetMinStorageSize(v int32)`

SetMinStorageSize sets MinStorageSize field to given value.

### HasMinStorageSize

`func (o *CloudMachineType) HasMinStorageSize() bool`

HasMinStorageSize returns a boolean if a field has been set.

### SetMinStorageSizeNil

`func (o *CloudMachineType) SetMinStorageSizeNil(b bool)`

 SetMinStorageSizeNil sets the value for MinStorageSize to be an explicit nil

### UnsetMinStorageSize
`func (o *CloudMachineType) UnsetMinStorageSize()`

UnsetMinStorageSize ensures that no value is present for MinStorageSize, not even an explicit nil
### GetMaxStorageSize

`func (o *CloudMachineType) GetMaxStorageSize() int32`

GetMaxStorageSize returns the MaxStorageSize field if non-nil, zero value otherwise.

### GetMaxStorageSizeOk

`func (o *CloudMachineType) GetMaxStorageSizeOk() (*int32, bool)`

GetMaxStorageSizeOk returns a tuple with the MaxStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorageSize

`func (o *CloudMachineType) SetMaxStorageSize(v int32)`

SetMaxStorageSize sets MaxStorageSize field to given value.

### HasMaxStorageSize

`func (o *CloudMachineType) HasMaxStorageSize() bool`

HasMaxStorageSize returns a boolean if a field has been set.

### SetMaxStorageSizeNil

`func (o *CloudMachineType) SetMaxStorageSizeNil(b bool)`

 SetMaxStorageSizeNil sets the value for MaxStorageSize to be an explicit nil

### UnsetMaxStorageSize
`func (o *CloudMachineType) UnsetMaxStorageSize()`

UnsetMaxStorageSize ensures that no value is present for MaxStorageSize, not even an explicit nil
### GetMinIOPS

`func (o *CloudMachineType) GetMinIOPS() int32`

GetMinIOPS returns the MinIOPS field if non-nil, zero value otherwise.

### GetMinIOPSOk

`func (o *CloudMachineType) GetMinIOPSOk() (*int32, bool)`

GetMinIOPSOk returns a tuple with the MinIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIOPS

`func (o *CloudMachineType) SetMinIOPS(v int32)`

SetMinIOPS sets MinIOPS field to given value.

### HasMinIOPS

`func (o *CloudMachineType) HasMinIOPS() bool`

HasMinIOPS returns a boolean if a field has been set.

### SetMinIOPSNil

`func (o *CloudMachineType) SetMinIOPSNil(b bool)`

 SetMinIOPSNil sets the value for MinIOPS to be an explicit nil

### UnsetMinIOPS
`func (o *CloudMachineType) UnsetMinIOPS()`

UnsetMinIOPS ensures that no value is present for MinIOPS, not even an explicit nil
### GetMaxIOPS

`func (o *CloudMachineType) GetMaxIOPS() int32`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *CloudMachineType) GetMaxIOPSOk() (*int32, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *CloudMachineType) SetMaxIOPS(v int32)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *CloudMachineType) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *CloudMachineType) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *CloudMachineType) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetDefaultIOPS

`func (o *CloudMachineType) GetDefaultIOPS() int32`

GetDefaultIOPS returns the DefaultIOPS field if non-nil, zero value otherwise.

### GetDefaultIOPSOk

`func (o *CloudMachineType) GetDefaultIOPSOk() (*int32, bool)`

GetDefaultIOPSOk returns a tuple with the DefaultIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIOPS

`func (o *CloudMachineType) SetDefaultIOPS(v int32)`

SetDefaultIOPS sets DefaultIOPS field to given value.

### HasDefaultIOPS

`func (o *CloudMachineType) HasDefaultIOPS() bool`

HasDefaultIOPS returns a boolean if a field has been set.

### SetDefaultIOPSNil

`func (o *CloudMachineType) SetDefaultIOPSNil(b bool)`

 SetDefaultIOPSNil sets the value for DefaultIOPS to be an explicit nil

### UnsetDefaultIOPS
`func (o *CloudMachineType) UnsetDefaultIOPS()`

UnsetDefaultIOPS ensures that no value is present for DefaultIOPS, not even an explicit nil
### GetStorageOptions

`func (o *CloudMachineType) GetStorageOptions() []map[string]interface{}`

GetStorageOptions returns the StorageOptions field if non-nil, zero value otherwise.

### GetStorageOptionsOk

`func (o *CloudMachineType) GetStorageOptionsOk() (*[]map[string]interface{}, bool)`

GetStorageOptionsOk returns a tuple with the StorageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOptions

`func (o *CloudMachineType) SetStorageOptions(v []map[string]interface{})`

SetStorageOptions sets StorageOptions field to given value.

### HasStorageOptions

`func (o *CloudMachineType) HasStorageOptions() bool`

HasStorageOptions returns a boolean if a field has been set.

### GetGpus

`func (o *CloudMachineType) GetGpus() []CloudGPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *CloudMachineType) GetGpusOk() (*[]CloudGPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *CloudMachineType) SetGpus(v []CloudGPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *CloudMachineType) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### SetGpusNil

`func (o *CloudMachineType) SetGpusNil(b bool)`

 SetGpusNil sets the value for Gpus to be an explicit nil

### UnsetGpus
`func (o *CloudMachineType) UnsetGpus()`

UnsetGpus ensures that no value is present for Gpus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


