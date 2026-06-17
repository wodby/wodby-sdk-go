# NewManagedClusterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Demo** | **bool** |  | 
**SingleNode** | Pointer to **NullableBool** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**MachineType** | Pointer to **NullableString** |  | [optional] 
**MinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNewManagedClusterInput

`func NewNewManagedClusterInput(demo bool, ) *NewManagedClusterInput`

NewNewManagedClusterInput instantiates a new NewManagedClusterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewManagedClusterInputWithDefaults

`func NewNewManagedClusterInputWithDefaults() *NewManagedClusterInput`

NewNewManagedClusterInputWithDefaults instantiates a new NewManagedClusterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDemo

`func (o *NewManagedClusterInput) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *NewManagedClusterInput) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *NewManagedClusterInput) SetDemo(v bool)`

SetDemo sets Demo field to given value.


### GetSingleNode

`func (o *NewManagedClusterInput) GetSingleNode() bool`

GetSingleNode returns the SingleNode field if non-nil, zero value otherwise.

### GetSingleNodeOk

`func (o *NewManagedClusterInput) GetSingleNodeOk() (*bool, bool)`

GetSingleNodeOk returns a tuple with the SingleNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNode

`func (o *NewManagedClusterInput) SetSingleNode(v bool)`

SetSingleNode sets SingleNode field to given value.

### HasSingleNode

`func (o *NewManagedClusterInput) HasSingleNode() bool`

HasSingleNode returns a boolean if a field has been set.

### SetSingleNodeNil

`func (o *NewManagedClusterInput) SetSingleNodeNil(b bool)`

 SetSingleNodeNil sets the value for SingleNode to be an explicit nil

### UnsetSingleNode
`func (o *NewManagedClusterInput) UnsetSingleNode()`

UnsetSingleNode ensures that no value is present for SingleNode, not even an explicit nil
### GetRegion

`func (o *NewManagedClusterInput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewManagedClusterInput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewManagedClusterInput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NewManagedClusterInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *NewManagedClusterInput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *NewManagedClusterInput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetMachineType

`func (o *NewManagedClusterInput) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *NewManagedClusterInput) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *NewManagedClusterInput) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *NewManagedClusterInput) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *NewManagedClusterInput) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *NewManagedClusterInput) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetMinNodeCount

`func (o *NewManagedClusterInput) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *NewManagedClusterInput) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *NewManagedClusterInput) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *NewManagedClusterInput) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### SetMinNodeCountNil

`func (o *NewManagedClusterInput) SetMinNodeCountNil(b bool)`

 SetMinNodeCountNil sets the value for MinNodeCount to be an explicit nil

### UnsetMinNodeCount
`func (o *NewManagedClusterInput) UnsetMinNodeCount()`

UnsetMinNodeCount ensures that no value is present for MinNodeCount, not even an explicit nil
### GetMaxNodeCount

`func (o *NewManagedClusterInput) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *NewManagedClusterInput) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *NewManagedClusterInput) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *NewManagedClusterInput) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *NewManagedClusterInput) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *NewManagedClusterInput) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


