# ClusterCapabilityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoInstall** | **bool** |  | 
**StackRevisionIds** | **[]int32** | Approved immutable stack revision IDs for discovery and optional installation. | 

## Methods

### NewClusterCapabilityPolicy

`func NewClusterCapabilityPolicy(autoInstall bool, stackRevisionIds []int32, ) *ClusterCapabilityPolicy`

NewClusterCapabilityPolicy instantiates a new ClusterCapabilityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCapabilityPolicyWithDefaults

`func NewClusterCapabilityPolicyWithDefaults() *ClusterCapabilityPolicy`

NewClusterCapabilityPolicyWithDefaults instantiates a new ClusterCapabilityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoInstall

`func (o *ClusterCapabilityPolicy) GetAutoInstall() bool`

GetAutoInstall returns the AutoInstall field if non-nil, zero value otherwise.

### GetAutoInstallOk

`func (o *ClusterCapabilityPolicy) GetAutoInstallOk() (*bool, bool)`

GetAutoInstallOk returns a tuple with the AutoInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInstall

`func (o *ClusterCapabilityPolicy) SetAutoInstall(v bool)`

SetAutoInstall sets AutoInstall field to given value.


### GetStackRevisionIds

`func (o *ClusterCapabilityPolicy) GetStackRevisionIds() []int32`

GetStackRevisionIds returns the StackRevisionIds field if non-nil, zero value otherwise.

### GetStackRevisionIdsOk

`func (o *ClusterCapabilityPolicy) GetStackRevisionIdsOk() (*[]int32, bool)`

GetStackRevisionIdsOk returns a tuple with the StackRevisionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackRevisionIds

`func (o *ClusterCapabilityPolicy) SetStackRevisionIds(v []int32)`

SetStackRevisionIds sets StackRevisionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


